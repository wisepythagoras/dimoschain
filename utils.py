import os

from definitions import *


def getopts(argv):
    """ Read the arguments from the command line. """

    opts = {}

    # Read through all the arguments.
    while argv:
        # This is an argument.
        if argv[0][0] == '-':
            # Add the argument to the options dict.
            if len(argv) > 1 and argv[1][1] != "-":
                opts[argv[0][1:]] = argv[1]
            else:
                opts[argv[0][1:]] = True

        # Shift the next argument.
        argv = argv[1:]

    # Return a structured dict of the command line arguments.
    return opts


def file_exists(filename):
    """ Checks if a file exists. """

    return os.path.isfile(filename)


def read_blockchain_file(filename):
    """ Reads a file from the blockchain's directory. """

    if not blockchain_dir_exists():
        return None

    if file_exists(filename):
        with open(filename) as f:
            return f.read()

    return None


def get_base_dir():
    """ Returns the base directory. """

    return "{home}/{base}".format(home=os.path.expanduser("~"), base=BASE_DIR)


def ensure_base_dir_exists():
    """ Check if the wallet's dir exists or not. """

    if not os.path.isdir(get_base_dir()):
        os.makedirs(get_base_dir())


def blockchain_dir_exists():
    """ Checks if the blockchain's dir is there or not. """

    return os.path.isdir(BLOCKCHAIN_DIR)


def get_genesis_hash():
    """ Gets the genesis block's hash. """

    return read_blockchain_file(GENESIS_HASH_FILE)
        

def get_current_hash():
    """ Returns the hash of the current block in the blockchain. """

    return read_blockchain_file(CURRENT_HASH_FILE)


def get_current_height():
    """ Gets the current height of the blockchain. """

    height = read_blockchain_file(BLOCKCHAIN_HEIGHT_FILE)

    if height is not None:
        height = int(height)

    return height


def get_urand():
    """ Gets bytes from /dev/urandom """

    return os.urandom(33);

