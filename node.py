import os.path

import coincurve
import random
import requests
import msgpack

from encryption import AESCipher
from blockchain import Blockchain
import utils
from definitions import BASE_DIR
import codecs


class Node(object):

    """
        This object represents a node in the network. A node
        is basically a peer that can connect to the master nodes.
    """

    __private_key__ = None
    __public_key__ = None

    def __init__(self, private_key=None):
        """ Create a new node object """

        if private_key is not None:
            self.set_private_key(private_key)

    def set_private_key(self, private_key):
        """ Set an existing private key. """

        self.__private_key__ = coincurve.PrivateKey.from_hex(private_key)
        self.set_public_key()

    def create_keys(self):
        """ Create a new set of keys. """

        # Create a new private key.
        self.__private_key__ = coincurve.PrivateKey()
        self.set_public_key()

    def set_public_key(self):
        """ Set the public key. """

        self.__public_key__ = self.__private_key__.public_key

    def get_raw_public_key(self):
        """ Get the raw public key """

        return self.__public_key__.format(compressed=True)

    def get_public_key(self):
        """ Get the public key of the node """

        # Get the key.
        key = self.get_raw_public_key()

        return key.hex()

    def get_private_key(self):
        """ Get the private key of the node """

        return self.__private_key__.to_hex()

    def sign(self, message, encoded=False):
        """ Sign the message with the private key """

        # Get the actual signature.
        sig = self.__private_key__.sign(message.encode('utf-8'))

        # Return the raw signature if encoded is False.
        if not encoded:
            return sig

        # Choose what type of signature to return.
        return sig.hex()

    def verify(self, signature, message, public_key=None):
        """ Verify that the message with the public key """

        # This is the option of when a third-party public is imported.
        if public_key is not None:
            return coincurve.PublicKey(public_key) \
                .verify(signature, message.encode('utf-8'))

        # Otherwise verify with our own signature.
        return self.__public_key__.verify(signature, message)

    def save_keys(self, name, password=None):
        """ Save the key to the wallet directory. """

        # Ensure that the base directory exists.
        utils.ensure_base_dir_exists()

        # Construct the path name for the new wallet.
        path_name = "{base}/{name}".format(
            base=utils.get_base_dir(), name=name)

        # Don't overwrite existing wallets because that's dangerous.
        if os.path.isfile("{}.wallet".format(path_name)):
            print("{}.wallet".format(path_name))
            return False

        mode = "wb"

        # Open the wallet file so that we can save the payload into it.
        with open("{}.wallet".format(path_name), mode) as f:
            payload = {
                "private": self.get_private_key(),
                "public": self.get_public_key()
            }

            data = msgpack.packb(payload)

            if password is not None:
                # If a user set a password, then create an AES cipher and
                # encrypt the wallet.
                aes = AESCipher(password)

                # Encrypt the payload before saving it.
                data = aes.encrypt(data)

            # Write the file to the filesystem.
            f.write(data)

        return True

    def read_keys(self, name=None, path=None, password=None):
        """ Read a wallet file. """

        # There must be at least one argument.
        if name is None and path is None:
            return False

        # If a name was passed then create the path string.
        if name is not None:
            path = "{base}/{name}.wallet".format(
                base=utils.get_base_dir(), name=name)

        # The file doesn't exist.
        if path is not None and not os.path.isfile(path):
            return False

        f = codecs.open(path, "rb", errors="ignore")

        # Read the data from the wallet file.
        data = f.read()

        if password is not None:
            # Create a new cipher with the password that was supplied.
            aes = AESCipher(password)

            # Decrypt the data before unpacking it.
            data = aes.decrypt(data)

        try:
            # Read the packed keys.
            keys = msgpack.unpackb(data, raw=False)

            # And set them locally.
            self.set_private_key(keys["private"])

            return True
        except:
            print("Unable to decode/decrypt wallet.")
            return False

    def submit_transaction(self, recipient, amount):
        """ Submit a transaction to the master nodes. """

        pass

    def get_balance(self):
        """ Get the balance of the current account. """

        pass

