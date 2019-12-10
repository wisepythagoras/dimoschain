from node import Node
import utils
from definitions import BASE_DIR


class Peer(Node):

    """
        The peer is the network client and simple user of the
        network who can send simple transactions to other peers.
    """

    def __init__(self, private_key=None):
        """ The constructor of the Peer class. """

        # Initialize the parent class.
        Node.__init__(self, private_key)
