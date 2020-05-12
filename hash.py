import sys
import binascii
import math
import hashlib

if sys.version_info < (3, 6):
    import sha3

from utils import PVERS
from prand import PRand

LEN = 64


def sha3_512(string, raw_digest=False):
    """ Performs the SHA3-512 on a string. """

    # Create the SHA3 hash.
    hash = hashlib.sha3_512()

    # Add the string to the digest.
    hash.update(string)

    # And finally return the hex or raw digest.
    return hash.hexdigest() if not raw_digest else hash.digest()


class Pyrin(object):

    """
        An experimental, computationally heavy hashing algorithm

        WARNING:
            This shouldn't be used until it's proven to provide
            adequate resistance against a birthday attack.
    """

    def shuffle_and_obfuscate(self, arr, rand):
        """ Shuffle an array and obfuscate it """

        for i in range(0, LEN):
            j = rand.rand() % LEN
            randNum = rand.rand() % 256

            arr[j] ^= (randNum & arr[i])
            arr[i] ^= (randNum & arr[j])

        return arr

    def to_hex(self, input):
        """ Converts the input to hexadecimal. """

        return binascii.hexlify(input)

    def hash(self, input):
        """ Hash a string """

        rand = PRand(input)
        part_a = []
        result = []

        for i in range(0, LEN):
            part_a.append(rand.rand() % 255)

        if len(input) <= LEN:
            # Fill up the string to make it into an even 64 bytes.
            for i in range(0, LEN):
                if i < len(input):
                    result.append(part_a[i] ^ ord(input[i]))
                else:
                    result.append(part_a[i] ^ (rand.rand() % 255))

                result[i] = (
                    (result[i] & rand.rand() % 255) ^ rand.rand() % 255)
        else:
            # Get the amount of parts.
            parts = int(math.ceil(len(input) / LEN))
            result = [0] * LEN

            for i in range(0, parts + 1):
                # Get the block and save it in the block variable.
                block = input[i * LEN:(i + 1) * LEN if (i + 1) * LEN < len(input) else len(input)]

                # Get the block's hash.
                block = self.hash(block)

                # XOR the block's hash with the result.
                for j in range(0, LEN):
                    result[j] = part_a[j] ^ ord(block[j])

                result[i] = (
                    (result[i] & rand.rand() % 255) ^ rand.rand() % 255)

        # Run the final round
        for i in range(0, rand.rand() % 90 + 10):
            result = self.shuffle_and_obfuscate(result, rand)

        strResult = ""

        for i in result:
            strResult = strResult + chr(i)

        # Convert to hex
        return strResult
