from node import Node
import utils
from definitions import BASE_DIR


class MasterNode(Node):

    """
        The master node accepts connections, distributes the blockchain, updates
        to the blockchain (ie. new blocks) and accept new transactions from
        peers.
    """

    def __init__(self, private_key=None):
        """ The constructor of the MasterNode class. """

        # Initialize the parent class.
        Node.__init__(self, private_key)
