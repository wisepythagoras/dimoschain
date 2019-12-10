import os.path
import sys

PVERS = sys.version_info[0]
NAME = "Dimos"

# Define the working directories. Aka, where the blockchain and
# the wallet will be stored.
BASE_DIR = ".dimos"
BLOCKCHAIN_DIR = "{home}/{base}/blockchain".format(home=os.path.expanduser("~"),
                                                   base=BASE_DIR)

# This is the "aparchi." In Greek it means the very beginning.
APARCHI = "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"

# Filenames of useful files.
GENESIS_HASH_FILE = "{}/genesis_block".format(BLOCKCHAIN_DIR)
CURRENT_HASH_FILE = "{dir}/current_hash".format(dir=BLOCKCHAIN_DIR)
BLOCKCHAIN_HEIGHT_FILE = "{dir}/blockchain_height".format(dir=BLOCKCHAIN_DIR)

# The magic bytes is just the string that will always preceed all traffic that
# will be sent to any of the master nodes. If traffic doesn't start with this
# particular sequence, then the traffic will be disregarded.
MAGIC_BYTES = "\x69\x0f\x69"
