from node import Node
import utils
from definitions import BASE_DIR


class MasterNode(Node):

    """
        The master node accepts connections, distributes the blockchain, updates
        to the blockchain (ie. new blocks) and accept new transactions from
        peers.
    """

    is_voting_member = False
    state = None

    def __init__(self, private_key=None):
        """ The constructor of the MasterNode class. """

        # Initialize the parent class.
        Node.__init__(self, private_key)

    def set_is_voting_member(self, is_voting_member):
        """ Set the local is_voting_member member variable. """

        self.is_voting_member = is_voting_member

    def can_vote(self):
        """ Get True or False if the MN can vote or not. """

        return self.is_voting_member

    def add_randomness(rand):
        """ Add randomness to the master node's state. """

        self.state = [chr(ord(r1) ^ ord(r2)) for r1, r2 in zip(rand, self.state)]

