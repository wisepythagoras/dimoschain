#!/usr/bin/env python3

import sys
import json
import time

sys.path.insert(0, sys.path[0] + r'/../')

import random
import asyncio
import socketio
from hash import sha3_512
from utils import getopts
from peer import Peer

node = None
config = None

loop = asyncio.get_event_loop()
sio = socketio.AsyncClient()

@sio.on('connect')
async def on_connect():
    print('connected to server')
    await sio.emit('hello')

@sio.on('hello')
async def on_hello(args):
    print('hello', args)

async def start_client():
    mn = config['master_nodes'][0]

    print('http://{}:{}'.format(mn['host'], mn['port']))

    await sio.connect('http://{}:{}'.format(mn['host'], mn['port']))
    await sio.wait()

if __name__ == '__main__':
    # Get the command line arguments.
    opts = getopts(sys.argv)

    if "h" in opts:
        print("Wallet usage:")
        print("-o <path> Open an existing wallet")
        print("-c <path> The path to the network configuration")
        print("-p <password> The password to decrypt the wallet with")
        print("-h Show this help message")
        sys.exit(0)

    # Require the path to the configuration file.
    if "c" not in opts:
        print("A network configuration is needed with -c")
        sys.exit(1)

    # Make sure we've got all the necessary options.
    if "o" not in opts:
        print("Peer wallet path is needed")
        sys.exit(1)

    # Get the configuration.
    with open(opts["c"], "r") as jsonfile:
        config = json.load(jsonfile)

    # Create the peer object.
    node = Peer()
    has_password = True if "p" in opts else False

    if not node.read_keys(path=opts["o"], password=(opts["p"] if has_password else None)):
        print("No such wallet or unable to open")
        sys.exit(1)

    print("Public: " + node.get_public_key())
    print("Private: " + node.get_private_key())

    loop.run_until_complete(start_client())

