from Crypto.Hash import SHA512
from Crypto.Hash import HMAC
from struct import pack


class PRBG(object):
    """ The pseudo-random byte/number generator class. """

    def __init__(self, seed):
        """ Default generator/constructor of our custom PRNG """

        # The current position/index.
        self.index = 0

        # The initial seed. This could be anything, but preferably random.
        self.seed = HMAC.new(seed.encode(), digestmod=SHA512).digest()

        # To be used to store the raw hash buffer parts.
        self.buffer = b''

    def __call__(self, n):
        """ Runs when the PRNG is called. """

        while len(self.buffer) < n:
            # Pack the index back into the seed and create a hash from it, and do this
            # until the minimum length criteria is met.
            self.buffer += HMAC.new(self.seed + pack('<I', self.index), digestmod=SHA512).digest()
            self.index += 1

        result, self.buffer = self.buffer[:n], self.buffer[n:]

        # Return the entropy.
        return result

