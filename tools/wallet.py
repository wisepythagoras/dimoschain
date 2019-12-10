import sys
import os.path

sys.path.append(sys.path[0] + r'/../')

from blockchain import Blockchain
from block import Block
from node import Node
from master_node import MasterNode
from utils import getopts
from hash import sha3_512

# Get the command line arguments.
opts = getopts(sys.argv)

if "h" in opts:
    print("Wallet usage:")
    print("-c <name> Create a new wallet")
    print("-o <path> Open an existing wallet")
    print("-m <ip address> Create a master node")
    print("-p <password> The password to encrypt/decrypt the wallet with")
    print("-h Show this help message")
    sys.exit(0)

# Initialize the blockchain.
blockchain = Blockchain()

# Validate the current state of the blockchain.
if not blockchain.load_current_state():
    print("No blockchain to read from")
    sys.exit(1)

has_password = False
is_master = False
node = None

# Create a new node.
if "m" in opts:
    is_master = True
    node = MasterNode()
else:
    node = Node()

if "p" in opts:
    has_password = True

if "c" in opts and type(opts["c"]) is str:
    # Create a new set of keys.
    node.create_keys()

    # Define the name.
    name = opts["c"] + ('_m' if is_master else '')

    # Save the keys.
    if not node.save_keys(name, password=(opts["p"] if has_password else None)):
        print("Unable to save wallet")
        sys.exit(1)

    print("Wallet '{}' was created in ~/.dimos".format(name))
    print("Public: " + node.get_public_key())
    print("Private: " + node.get_private_key())

    if is_master:
        # Create the hash of the IP address.
        address_hash = sha3_512(opts["m"])

        # Get the signature of the hash.
        sig = node.sign(address_hash)

        # Verify the signature with the public key.
        verified = node.verify(sig, address_hash, node.get_raw_public_key())

        print("Sig: " + sig.encode('hex'))
        print("Verified: " + ("Yes" if verified else "No"))

elif "o" in opts and type(opts["o"]) is str:
    if not node.read_keys(path=opts["o"], password=(opts["p"] if has_password else None)):
        print("No such wallet or unable to open")
        sys.exit(1)
    
    print("Public: " + node.get_public_key())
    print("Private: " + node.get_private_key())

elif "o" in opts or "c" in opts:
    print("Unrecognized values")
