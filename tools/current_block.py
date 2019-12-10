import sys
import os.path

sys.path.append(sys.path[0] + r'/../')

from blockchain import Blockchain
from block import Block

# Initialize the blockchain.
blockchain = Blockchain()

if not blockchain.load_current_state():
    print("No blockchain to read from")
    sys.exit(1)

# Get the hash of the current block.
print(blockchain.get_current_block_hash())
