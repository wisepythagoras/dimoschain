import sys
import os.path

sys.path.append(sys.path[0] + r'/../')

from definitions import APARCHI
from blockchain import Blockchain
from block import Block
from transaction import Transaction

# Initialize the blockchain.
blockchain = Blockchain()

# Create a new transaction.
transaction = Transaction()
transaction.set("a", float(0.000000001), "b")

# Read the current block file.
if blockchain.load_current_state():
    # Get the current block.
    hash = blockchain.current_hash

    # Create a new block.
    new_block = Block(blockchain.db)

    # Set a symbolic transaction for the genesis block.
    new_block.set_transactions([transaction])

    # Set the time on the blockchain.
    new_block.set_time()

    # Set the id on the block.
    new_block.set_idx()

    # Set the previous hash for the genesis block to all zeros.
    new_block.prev_hash = hash

    # Mine the new block.
    new_block.calculate_hash(set_hash=True)

    # And add the block.
    blockchain.add_block(new_block)

    print("The hash of the new block is:")
    print(new_block.hash)

