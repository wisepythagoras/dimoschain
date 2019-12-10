# Configuration

In this folder there will be configuration for different types of networks. The `testnet_local.json` configuration file contains the appropriate settings for running a full network on your local network. The addresses that are in there, under `master_nodes` are fixed, but can be changed to whatever you wish.

## Fields

Below we'll go through the documentation of each field in the configuration file.

### `master_nodes`: `object[]`

The master nodes are integral in making the network usable at all. Peer and delegate nodes will connect to these for various purposes, which include:

1. Submitting transactions
2. Submitting a request to become a delegate
3. Receive blocks by their hash

#### - `master_nodes.name`: `string`

The name of the master node.

#### - `master_nodes.host`: `string`

The IP address or the domain of the master node.

#### - `master_nodes.port`: `int`

The port that the master node is running on.

#### - `master_nodes.sig`: `string`

The signature of the IP of the master node. This helps to not be able to forge the location of a master node that would put the network in danger.

#### - `master_nodes.pub_key`: `string`

The public key of the master node. This is necessary in order for network participants (all nodes on the network) to be able to verify actions of the master node.

### `geo_based`: `bool`

This will be true if you want to make sure all peers and delegates connect to the master node that is nearest to them.

### `load_based`: `bool`

If set to `true`, the master nodes will have rate limiting based on the total amount of peers on the network, which would ensure that processing load is evenly distributed across the network. This setting overpowers, in a way, the `geo_based` setting.

### `proto`: `string`

This setting tells the blockchain what protocol to use. If you're in a test net and local environment, then a setting without SSL/TLS is typically fine. If you want to run this in production, then it's important to use SSL/TLS.
