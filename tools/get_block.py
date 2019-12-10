import sys
import os.path

sys.path.append(sys.path[0] + r'/../')

from blockchain import Blockchain
from block import Block
from utils import getopts

# Get the command line arguments.
opts = getopts(sys.argv)

# Check if there was a hash passed.
if "h" not in opts.keys():
    print("No hash passed")
    sys.exit(1)

# Initialize the blockchain.
blockchain = Blockchain()

if not blockchain.load_current_state():
    print("No blockchain to read from")
    sys.exit(1)

# Create a new block object.
block = Block(blockchain.db)

if len(opts["h"]) != 128:
    print("Malformed hash")
    sys.exit(1)

if not block.load(opts["h"]):
    print("No such block found")
    sys.exit(1)

# Get the hash of the current block.
print(block.printable())
