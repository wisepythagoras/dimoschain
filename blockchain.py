import json
import os.path

import msgpack
import plyvel

from definitions import *
from block import Block
import utils


class Blockchain(object):

    """ This is the blockchain. """

    current_hash = None
    height = 0

    def __init__(self):
        """ Create the blockchain object. """

        self.db = plyvel.DB(BLOCKCHAIN_DIR, create_if_missing=True)
        self.load_current_state()

    def load_current_state(self):
        """ Load the current state of the blockchain. """

        if self.current_hash is not None:
            return True

        # Get the current hash.
        self.current_hash = self.get_current_block_hash()

        if self.current_hash is None:
            return False

        # Try to load the height, otherwise try to calculate it. The result
        # will be 'None' if there was none.
        height = utils.get_current_height()

        if height is not None:
            self.height = height
        else:
            self.calculate_height()

        return True

    def get_height(self):
        """ Returns the height of the blockchain. """

        return self.height

    def get_block(self, hash):
        """
            Retrieves a block.

            Parameters
            ----------
            hash : str
        """

        # Create the block object.
        block = Block(self.db)

        # Get the block.
        if not block.load(hash):
            return None

        return block

    def get_genesis_hash(self):
        """ Does what it says it does. """

        return utils.get_genesis_hash()

    def add_block(self, block):
        """
            Add a block to the blockchain.

            Parameters
            ----------
            block : Block
        """

        # Get the hash of the block
        hash = block.hash

        # If the block already exists, exit.
        if block.load():
            return False

        # Validate the hash of the block.
        if block.calculate_hash() != hash:
            return False

        # There should be an id on the block.
        if block.idx is None:
            return False

        # Get the raw block.
        raw_block = block.get()

        # Set the block into the blockchain DB. The payload, which is the raw
        # block, will be in a binary format created by msgpack.
        self.db.put(b'{}'.format(hash), msgpack.packb(raw_block))

        # Try to see if the put was successfull.
        if block.load():
            with open(CURRENT_HASH_FILE, "w") as f:
                f.write(block.hash)

            # Increment the blockchain's height.
            self.height = self.height + 1

            # Update the height of the blockchain.
            self._write_height()

            return True

        return False

    def blockchain_dir_exists(self):
        """ Checks if the blockchain's dir is there or not. """

        return utils.blockchain_dir_exists()

    def get_current_block_hash(self):
        """ Read the current block hash. """

        # Check if the blockchain is initialized in the first place.
        if not self.blockchain_dir_exists():
            return None

        # This either means that the blockchain has not yet been born,
        # or that some file was deleted. Either way, no good.
        if not os.path.isfile(CURRENT_HASH_FILE):
            return None

        # Read the current block file.
        return utils.get_current_hash()

    def traverse(self):
        """ Traverse the blockchain. """

        for key, value in self.db:
            print("{key}: {value}".format(key=key, value=value))

    def _write_height(self):
        """ Private: Writes the height to the blockchain's directory. """

        # Write the height into the blockchain's directory.
        with open(BLOCKCHAIN_HEIGHT_FILE, "w") as f:
            f.write(str(self.height))

    def calculate_height(self):
        """ Calculate the height of the blockchain. """

        count = 0

        # Iterate through the database to count the height.
        with self.db.iterator() as it:
            for _, _ in it:
                count = count + 1

        self.height = count

        # Write the height into the blockchain's directory.
        self._write_height()

        return count

    def is_valid(self):
        """ Returns if the blockchain is valid. """

        # Make sure things are in place.
        if not self.blockchain_dir_exists():
            return False

        # Get the current hash and set it as the one that will be loaded
        # as the starting point.
        current_hash = self.current_hash

        while True:
            # Create a new block.
            current_block = Block(self.db)

            # Load the block.
            if not current_block.load(current_hash):
                print("Could not get block:\n{}".format(current_hash))

            if current_hash != current_block.calculate_hash():
                print("Failed at the following block:")
                print(json.dumps(current_block.get()))
                return False

            print("[OK] {}".format(current_hash))

            # Set the previous hash.
            current_hash = current_block.prev_hash

            del current_block

            if current_hash == APARCHI:
                break

        return True
