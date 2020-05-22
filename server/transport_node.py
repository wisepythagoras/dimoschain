#!/usr/bin/env python3

"""
    Description
    ===========

    The server, as you can see, is based on a WebSockets logic. The reason I picked WebSockets
    lies in the fact that it's a widely used technology that supports all languages and
    platforms. This will allow anyone to create a real-time web based user interface that will
    be able to completely bypass the use of a single server and allow the user/peer to connect
    directly to any transport node.

    Protocol Description
    ====================

    Much like any type of server, the transport node server will speak a specific language, and
    that's the Dimos Networking Communication Protocol (DNP). DNP will be a layer on top of an
    already secured TCP connection with the use of SSL/TLS certificates.

    1. Message Format

    The format of both requests and responses will be the following:

    +-----+--------------------+------+
    | CMD | H(SIG(CMD + ARGS)) | ARGS |
    +-----+--------------------+------+

    To explain what the heck the table above means, The message starts with the type of request
    (otherwise known as command: CMD) the peer wants carried out. The second part of the message
    is the hash of the signature, which has a predictable length.
    Then, at the end we leave the arguments, which could have an arbitrary size, but no larger
    than 100 characters. The 100 character limit allows anyone to be able to fit the public key
    of a peer and a 33 digit amount to send, which is never going to be the case.

    2. Return Codes

    2.1 Success Codes

    2.1.1 `HELLO` Welcome message
    2.1.2 `SUCCESS` Success message

    2.2 Error Codes

    2.2.1 `ERR_SIG` Error in verifying the signature
    2.2.2 `ERR_TSK` Error in task completion
    2.2.3 `ERR_FRM` Error in request format
    2.2.4 `ERR_CMD` Unknown command

    3. Available Commands

    In any case, if the command sent was not formatted correctly, thus resulting in the inability
    to parse it, the transport node will return an `ERR_FRM`. If the message is parsed, yet the
    command is not valid, the response will be `ERR_CMD`.

    3.1 `HELLO`

    The `HELLO` command is the first that should be sent to a transport node when a peer wishes to
    connect to it. The only argument should be the public key of the connecting peer.

    The server responds with an `ERR_SIG` or with `HELLO`. If the server responds with `HELLO`,
    it also sends a set of challenges that the peer should complete and send to the server. These
    challenges are bytes that will aide in the generation or pseudo random bytes with the use of
    the peer's public key as a seed appended by these byte sequences that were sent from the transport
    node.

    Knowning the algorithm, the transport node will be able to verify these values against the peer's
    public key and respond with `ERR_RNG` if the peer failed to produce valid byte sequences or
    with `SUCCESS`, after which the peer is authenticated and able to carry out other requests.

    3.2 `SEND`

    The `SEND` command will be used by the peers to transfer funds from their ownership to other
    peers. The arguments will contain the receiver's public key followed by the amount that's
    being transfered.

    Since this request is added into a queue and handled by the delegates, the only responses are
    `ERR_SIG` if the transport node was unable to verify the signature, or `SUCCESS`.

    TBD

"""

# https://python-socketio.readthedocs.io/en/latest/client.html
# https://github.com/miguelgrinberg/python-socketio/blob/transport/examples/server/aiohttp/latency.py
# https://github.com/miguelgrinberg/python-socketio/blob/transport/examples/client/asyncio/latency_client.py

import sys
import json
import time

sys.path.insert(0, sys.path[0] + r'/../')

import msgpack
from aiohttp import web
import socketio
from utils import getopts
from definitions import (
    ERROR_NO_PUBKEY,
    ERROR_NO_SIGNATURE,
    WELCOME_MESSAGE,
)

from hash import sha3_512
from transport_node import TransportNode


# Create the transport node object.
mn = TransportNode()

# Get the command line arguments.
opts = getopts(sys.argv)
config = None
mn_config = None

if "h" in opts:
    print("Wallet usage:")
    print("-o <path> Open an existing wallet")
    print("-c <path> The path to the network configuration")
    print("-n <string> The name of the transport node")
    print("-p <password> The password to decrypt the wallet with")
    print("-h Show this help message")
    sys.exit(0)

# Require the name.
if "n" not in opts:
    print("A name is needed with -n")
    sys.exit(1)

# Require the path to the configuration file.
if "c" not in opts:
    print("A network configuration is needed with -c")
    sys.exit(1)

# Make sure we've got all the necessary options.
if "o" not in opts:
    print("Transport node wallet path is needed")
    sys.exit(1)

# Read the keys of the node.
if not mn.read_keys(path=opts["o"], password=(opts["p"] if "p" in opts else None)):
    print("No such wallet or unable to open")
    sys.exit(1)

# Get the configuration.
with open(opts["c"], "r") as jsonfile:
    config = json.load(jsonfile)

# A set that will contain all the users.
USERS = set()

if __name__ == '__main__':
    # Get the configuration for the server.
    for server in config["transport_nodes"]:
        if server["name"] == opts["n"]:
            mn_config = server

    if mn_config is None:
        print("No transport node with name '{}' exists".format(opts["n"]))
        sys.exit(1)

    # Create the hash of the IP address.
    address_hash = sha3_512(mn_config["host"])

    # Get the signature of the hash.
    sig = mn.sign(address_hash, encoded=True)

    # Check that the signatures match.
    if sig != mn_config["sig"]:
        print("Signature doesn't match")
        sys.exit(1)

    print("Sig: " + sig)

    sio = socketio.AsyncServer(async_mode='aiohttp')
    app = web.Application()
    sio.attach(app)

    async def index(request):
        return web.Response(text="This is a Dimosthenes transport node", content_type='text/html')

    async def respond(room: str, event: str, payload: dict):
        """ Send a response to a room. """

        await sio.emit(event, msgpack.packb(payload), room=room)

    def error_payload(error: str):
        """ Create an error payload. """

        return {
            'e': error
        }

    @sio.on('connect')
    async def connect(sid: str, environ):
        print('connect ', sid)

    @sio.on('tn_auth')
    async def transport_node_auth(sid, payload):
        """
            This is the endpoint that's hit in order for a transport node
            to be authenticated.
        """

        # The payload may come as a string, in which case we need to decode it.
        if isinstance(payload, str):
            payload = msgpack.unpackb(payload, raw=False)

        # Validate the inputs from the other transport nodes.
        if 'pk' not in payload:
            await respond(sid, 'tn_auth_res', error_payload(ERROR_NO_PUBKEY))
            sio.disconnect(sid)
            return
        elif 's' not in payload:
            await respond(sid, 'tn_auth_res', error_payload(ERROR_NO_SIGNATURE))
            sio.disconnect(sid)
            return

        response = {
            'm': WELCOME_MESSAGE
        }

        respond(sid, 'tn_auth_res', response)

    @sio.on('hello')
    async def ping(sid):
        print('[CN]:', sid)
        await sio.emit('hello', 'there', room=sid)

    @sio.on('disconnect')
    def test_disconnect(sid):
        print('[DC]', sid)

    app.router.add_get('/', index)

    web.run_app(app, port=mn_config["port"], host=mn_config["host"])

