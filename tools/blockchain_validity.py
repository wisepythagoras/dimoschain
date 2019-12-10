import sys

sys.path.append(sys.path[0] + r'/../')

from definitions import APARCHI
from blockchain import Blockchain
from block import Block

# Initialize the blockchain.
blockchain = Blockchain()

# Load the current state of the blockchain.
if not blockchain.load_current_state():
    print("Unable to load the current state of the blockchain")
    sys.exit(1)

print("Current hash:\n{}".format(blockchain.current_hash))

if blockchain.is_valid():
    print("The blockchain is valid")

    # Get the height and print it.
    print("Height: {}".format(blockchain.get_height()))
else:
    print("The blockchain is invalid")
