import sys

sys.path.append(sys.path[0] + r'/../')

from definitions import APARCHI, BLOCKCHAIN_DIR
from blockchain import Blockchain
from block import Block
from transaction import Transaction
import utils

import plyvel

# Initialize the blockchain.
blockchain = Blockchain()

# Check if the blockchain has already been initialized.
if blockchain.load_current_state():
    print("The blockchain has already been initialized")
    print("The genesis hash is the following:")
    print(utils.get_genesis_hash())
    sys.exit(1)

# Random genesis text.
text = "Blabidy blah; tweedly doo; shut the heck up, or I'll bash you"

# Create a new transaction.
transaction = Transaction()
transaction.set(None, text, None)

# Create a new block.
genesis_block = Block(blockchain.db)

# Set a symbolic transaction for the genesis block.
genesis_block.set_transactions([transaction])

# Set the time on the blockchain.
genesis_block.set_time()

# Set the id, which should be 0.
genesis_block.idx = 0

# Set the previous hash for the genesis block to all zeros.
genesis_block.prev_hash = APARCHI

# Mine the new block.
genesis_block.calculate_hash(set_hash=True)

# And add the block.
if not blockchain.add_block(genesis_block):
    print("Unable to add the genesis block")
    sys.exit(1)

# Write the genesis block hash into the blockchain dir.
with open("{dir}/genesis_block".format(dir=BLOCKCHAIN_DIR), "w") as f:
    f.write(genesis_block.hash)

print("The genesis hash is:")
print(genesis_block.hash)
