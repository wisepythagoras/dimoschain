class PRand(object):

    """ This is a seeded PRNG. """

    seed = None
    all_chars = "0123456789abcdef"

    def __init__(self, seed):
        """ Construct the PRNG. """

        self.seed = seed

    def rand(self):
        """ Generate a pseudo random number. """

        # Generate the number.
        num = self.m_djb_hash()

        # Reset the seed.
        self.seed = format(num, "x")

        # And return the number.
        return num

    def m_djb_hash(self):
        """ Computes the DJB int. """

        # If there's no seed, return nothing.
        if self.seed is None:
            return 0

        # The initial hash.
        hash = 5381

        for char in self.seed:
            # Get the character code.
            c = ord(char)

            # Shift the hash by 5 positions left and subtract the current
            # hash from the result of the shift, and adds the character to
            # the new hash.
            hash = ((hash << 5) - hash) + c

            # Avoid having negative values by converting it into a 31-bit
            # integer.
            hash = hash & 0x7fffffff

        return hash
