import sys
import binascii
import math
import hashlib


def sha3_512(string, raw_digest=False):
    """ Performs the SHA3-512 on a string. """

    # Create the SHA3 hash.
    hash = hashlib.sha3_512()

    if isinstance(string, str):
        string = string.encode()

    # Add the string to the digest.
    hash.update(string)

    # And finally return the hex or raw digest.
    return hash.hexdigest() if not raw_digest else hash.digest()
