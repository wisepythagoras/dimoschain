<p align="center">
  <img src="/images/logo-small.png" alt="Logo of the Dimos project"
    title="Logo of the Dimos project" width="200" height="200">
</p>

# Dimos

Dimos (dee-mos, Δήμος: municipality, citizens) is a cryptocurrency based on a *Proof-of-Delegation* consensus algorithm.

The purpose of this cryptocurrency is to demonstrate that cryptocurrencies don't have to be profitable only for those who have the capital to invest in mining equipment or in units of crypto to stake. In the same way, it eliminates the need to maintain power-hungry miners (including ASIC) and creates a more egalitarian financial platform with fast transaction times.

The whitepaper is on its way.

## Getting Started

Install the dependencies:

``` sh
pip3 install msgpack plyvel merkletools coincurve websockets asyncio
mkdir ~/.dimos # This will become redundant and automated
```

Now you can create the blockchain by initializing it.

``` sh
python tools/genesis.py
```

This should output something like this:

```
The genesis hash is:
dfd13d010fe147734cdc1f63f582fb8b6c5004ba86502dde43d0afed0521d07d6af8cee13504fd4428d275897b32d5c271f75c83bb636d097a6a4b39268a9321
```

Now you can run the `get_hash.py` script and pass the genesis hash to it:

```
idx         0
hash        dfd13d010fe147734cdc1f63f582fb8b6c5004ba86502dde43d0afed0521d07d6af8cee13504fd4428d275897b32d5c271f75c83bb636d097a6a4b39268a9321
prev hash   00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
merkle root 6ac6178030ed7fcd0e73df1993f754eb18611d5706d54a83479f713e029db2438220235b6f3c2a760441648482254c92456cec91a0c51f88c45ac34856f43a54
timestamp   1535744574006
transactions
  None->None : Blabidy blah; tweedly doo; shut the heck up, or I'll bash you
None

```

You can also create new blocks, albeit this is more for testing purposes:

``` sh
python tools/create_block.py
The hash of the new block is:
baf27521a2e7e7ebd31ce94ffbfe3f40c45fcad742b8aaaf39272c598c36e1a6db3ad8f9278a0b3cf7e10b633fe5e88ba82f1a337034d839e2f80b45e26b52f8
```

And then you can also check that the blockchain is valid:

``` sh
python tools/blockchain_validity.py
Current hash:
bf479d3202d4213c0087d31d7ace52a8062696cbb060aac6db9008e3b3ae9b8a0624117b6c18d58c1e3a8ad1acd59e3440726df5602cb310aa7e78d0ccb4842c
[OK] bf479d3202d4213c0087d31d7ace52a8062696cbb060aac6db9008e3b3ae9b8a0624117b6c18d58c1e3a8ad1acd59e3440726df5602cb310aa7e78d0ccb4842c
[OK] c51eaa8f26588d90371e3c16f0caf468ec9f109709552067989b0050d06438c9f30305b3b8ea2c962bbc799dff636994988d192c1e42f192ce8c2a5de0f48d17
[OK] 857e9ff49178b5ec5cfb67cb72d698558abccb4614a0bed3f9191b946a21c7fdef0e8d4cbcbf336ae7feec10bd9576b7a47312b600f058b167b9ea89f7ad84cd
[OK] 817c692ac51c25100dfa40739a6ec94c6a22dfcaeeea316c99f7785d7c2899c0586fb939283d5e2d282878ca315a344c2959e373204f2e7dbe5b3dc11e505b83
[OK] 4debe154677076d630120730ac79b8c377679f16743c283860f02d32b6e43fae289e9293b2ea932c8b49aa592ab56d196671516bff032c42072ce0e68f154a55
[OK] baf27521a2e7e7ebd31ce94ffbfe3f40c45fcad742b8aaaf39272c598c36e1a6db3ad8f9278a0b3cf7e10b633fe5e88ba82f1a337034d839e2f80b45e26b52f8
[OK] dfd13d010fe147734cdc1f63f582fb8b6c5004ba86502dde43d0afed0521d07d6af8cee13504fd4428d275897b32d5c271f75c83bb636d097a6a4b39268a9321
The blockchain is valid
Height: 7
```

Finally, there's a script for creating a wallet:

``` sh
$ python tools/wallet.py -h
Wallet usage:
-c <name> Create a new wallet
-o <path> Open an existing wallet
-m <ip address> Create a master node
-p <password> The password to encrypt/decrypt the wallet with
-h Show this help message
$ python tools/wallet.py -c "Test Wallet"
Wallet 'Test Wallet' was created in ~/.dimos
$ ls ~/.dimos
 blockchain  'Test Wallet.wallet'
$ python tools/wallet.py -o ~/.dimos/Test\ Wallet.wallet
Public: 02adc56f757b7a91a3f35e3262656f0dd9fc935d0f5e92cc8f5a8f7851e3a6e1f0
Private: 2acac190d80a1a480c563e99d81e2038272de80b50f1ce54f973e5d22219d283
```

#### Warning

This software is far from complete and improved every day.

