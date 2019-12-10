from peer import Peer
import utils
from definitions import BASE_DIR


class Delegate(Peer):

    """
        The delegate is the miner of the network and serves for one
        full term.
    """

    def __init__(self, private_key=None):
        """ The constructor of the Delegate class. """

        # Initialize the parent class.
        Peer.__init__(self, private_key)
