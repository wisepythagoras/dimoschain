import json


class Transaction(object):

    """ This is the transaction object that lives in a block. """

    # The originating public address.
    origin = None

    # The destination public address.
    dest = None

    # The amount to transfer.
    amount = None

    # The id of the transaction.
    id = None

    # The digital signature of the transaction, by the originator.
    sig = None

    # The public address of the delegate who verified the transaction.
    verifier = None

    # The digital signature of the delegate on the transaction.
    verifier_sig = None

    def __init__(self):
        """ Create a transaction. """

    def set(self, origin, amount, dest):
        """
            Set the transaction.

            Parameters
            ----------
            origin : str
            amount : float
            dest : str
        """

        self.origin = origin
        self.dest = dest
        self.amount = amount

    def sign(self, signature):
        """
            Set the signature.

            Parameters
            ----------
            signature : str
        """

        self.sig = signature

    def get_signable(self):
        """ Get the transaction as is for signing. """

        return json.dumps([self.origin, self.amount, self.dest, self.id])

    def get_delegate_signable(self):
        """ Get the transaction as is for signing by a delegate. """

        return json.dumps([self.origin, self.amount, self.dest, self.id, self.sig,
                           self.verifier])

    def submit(self):
        """ Submit a new transaction as a peer for review by the delegation. """

        # TODO: use the private key to sign the transaction.

        pass

    def get(self):
        """ Return the JSON representation of the transaction. """

        return [self.origin, self.amount, self.dest, self.id, self.sig,
                self.verifier, self.verifier_sig]

