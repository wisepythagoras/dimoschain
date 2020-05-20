import json
import time

import msgpack
import merkletools
from hash import sha3_512
from utils import get_current_height
from transaction import Transaction


class Block(object):

    """ This class represents a block in the blockchain. """

    idx = None
    transactions = []
    timestamp = 0
    hash = None
    prev_hash = None
    db = None
    merkle_root = None

    def __init__(self, db):
        """ A block is part of a blockchain and contains transactions. """

        self.db = db

    def load(self, hash=None):
        """ Load the block from the database. """

        hash = self.hash if hash is None else hash

        if hash is None:
            return False

        block_cell = self.db.get(str.encode(hash))

        if block_cell is None:
            return False

        # Load the JSON from the block.
        block_cell = msgpack.unpackb(block_cell)

        # Get the transactions.
        transactions = block_cell["t"]

        # Make sure the array is empty.
        # TODO: There is a memory leak here. For some reason, transactions
        # from a previously loaded block leak into a new one when using the
        # validate function in the blockchain class.
        self.transactions = []

        for t in transactions:
            # Create the transaction object.
            transaction = Transaction()

            # Set the actual transaction.
            transaction.set(t[0], t[1], t[2])

            # Now set the signature and the id.
            transaction.id = t[3]
            transaction.sig = t[4]

            # Append the processed transaction.
            self.transactions.append(transaction)

        # Set the hashes.
        self.hash = block_cell["h"]
        self.prev_hash = block_cell["p"]
        self.merkle_root = block_cell["m"]

        # Set the id.
        self.idx = block_cell["i"]

        # And finally the timestamp.
        self.timestamp = block_cell["u"]

        del block_cell

        return True

    def set_time(self, when=None):
        """
            Set the time on the block.

            Parameters
            ----------
            when: time
        """

        if when is None:
            when = time

        # Set the timestamp on the block.
        self.timestamp = int(when.time() * 1000)

    def set_idx(self):
        """ Sets the block's height based on the current height of the blockchain. """

        self.idx = get_current_height()

    def get(self):
        """ Get the JSON representation of the block. """

        return {
            "i": self.idx,
            "t": self.get_raw_transactions(),
            "h": self.hash,
            "p": self.prev_hash,
            "u": self.timestamp,
            "m": self.merkle_root
        }

    def printable(self):
        """ Pretty prints a block. """

        print("idx         {idx}".format(idx=self.idx))
        print("hash        {hash}".format(hash=self.hash))
        print("prev hash   {prev}".format(prev=self.prev_hash))
        print("merkle root {root}".format(root=self.merkle_root))
        print("timestamp   {time}".format(time=self.timestamp))
        print("transactions")

        # Loop through the transactions to print them.
        for tx in self.transactions:
            print("  {orig}->{dest} : {amount}".format(orig=tx.origin, dest=tx.dest, amount=tx.amount))

    def set_transactions(self, transactions):
        """ Set the transactions on the block. """

        self.transactions = transactions

    def get_raw_transactions(self):
        """ Get the block's transactions as JSON. """

        transactions = []

        for transaction in self.transactions:
            transactions.append(transaction.get())

        return transactions

    def calculate_merkle_root(self, set_hash=False):
        """ Calculates the merkle hash. """

        # Create the merkle tree object.
        mr = merkletools.MerkleTools(hash_type="sha3_512")

        # Loop through the transactions and add them to the tree.
        for tx in self.transactions:
            mr.add_leaf(tx.hash(), True)

        # Create the tree.
        mr.make_tree()

        # Get the merkle root.
        merkle_root = mr.get_merkle_root()

        if set_hash:
            self.merkle_root = merkle_root

        # Return the merkle root.
        return merkle_root

    def calculate_hash(self, set_hash=False):
        """ Calculate this block's hash. """

        # Calculate the merkle root.
        self.calculate_merkle_root(set_hash)

        # Get the block and remove the hash.
        payload = self.get()

        # The hash of the block has not been computed yet, so that should be
        # removed from the payload.
        del payload['h']

        # Similarly, since the transactions will be represented by a merkle
        # tree, they should be removed.
        del payload['t']

        # Use msgpack to create a binary representation of the block, instead.
        payload = msgpack.packb(payload)

        # Calculate the hash.
        calculated_hash = sha3_512(payload)

        if set_hash:
            # Calculate the hash.
            self.hash = calculated_hash

        # Return the new hash.
        return calculated_hash

