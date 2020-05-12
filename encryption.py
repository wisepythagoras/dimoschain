from utils import PVERS
from hash import sha3_512
from Crypto import Random
from Crypto.Cipher import AES


class AESCipher(object):

    """
        A class for AES encryption.
    """

    bs = 32
    key = None

    def __init__(self, key):
        """ Create a new AESCipher object. """

        if PVERS == 3:
            self.key = "{!s: <32}".format(key.encode()).encode("utf-8")
        else:
            self.key = "{: <32}".format(key.encode()).encode("utf-8")

    def encrypt(self, message):
        """ Encrypt a message. """

        # Pad the message so that it's ready for AES.
        message = self._pad(message)

        # Generate a random new initialization vector.
        iv = Random.new().read(AES.block_size)

        # Create a new cipher.
        cipher = AES.new(self.key, AES.MODE_OFB, iv)

        # Return the encrypted string.
        return b"" + iv + cipher.encrypt(message)

    def decrypt(self, ciphertext):
        """ Decrypt some ciphertext. """

        # Extract the initialization vector.
        iv = ciphertext[:AES.block_size]

        # Create the AES cipher.
        cipher = AES.new(self.key, AES.MODE_OFB, iv)

        # Decrypt the string.
        return self._unpad(cipher.decrypt(ciphertext[AES.block_size:])) #.decode('utf-8')

    def _pad(self, string):
        """ Pad a string. """

        return string + ((self.bs - len(string) % self.bs) \
            * chr(self.bs - len(string) % self.bs)).encode()

    def _unpad(self, string):
        """ Unpad a string. """

        return string[:-ord(string[len(string)-1:])]
