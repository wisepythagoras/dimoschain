<p align="center">
  <img src="/images/logo-small.png" alt="Logo of the Dimos project"
    title="Logo of the Dimos project" width="200" height="200">
</p>

<h1 align="center">Dimos</h1>

<p align="center">
    <a href="https://goreportcard.com/report/github.com/wisepythagoras/dimoschain">
        <img src="https://goreportcard.com/badge/github.com/wisepythagoras/dimoschain" />
    </a>
    <a href="https://github.com/wisepythagoras/dimoschain/issues">
        <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"/>
    </a>
    <a href="https://github.com/wisepythagoras/dimoschain/blob/master/LICENSE">
        <img src="https://img.shields.io/badge/FOSS-100%25-4c1"/>
    </a>
</p>

Dimos (dee-mos, Δήμος: municipality, citizens) is a cryptocurrency based on a *Proof-of-Randomized-Delegation* consensus algorithm.

The purpose of this cryptocurrency is to demonstrate that cryptocurrencies don't have to be profitable only for those who have the capital to invest in mining equipment or in units of crypto to stake. In the same way, it eliminates the need to maintain power-hungry miners (including ASIC) and creates a more egalitarian financial platform with fast transaction times.

The whitepaper is on its way.

## Getting Started

Install the dependencies and build the package:

``` sh
make install-deps && make && make tests
```

I don't currently use Go Modules for a few reason, first of which is that I couldn't find the specific versions of the dependencies that I wanted. It also nuked my environment when I tried using it. So until I'm forced to use it, I'll use Go Deps.

Now you can create the blockchain by initializing it.

``` sh
./bin/create-genesis
```

This should output something like this:

```
badger 2019/03/21 11:41:01 INFO: All 0 tables opened in 0s
2019/03/21 11:41:01 Merkle Root: 97035fc67b851a86956f619f28c8ec7fce3b65e02a53ea175e4b7050ada4b2d4f7e53c34734321d59655a6647512ff91 <nil>
2019/03/21 11:41:01 Genesis Hash: 8283d0e96d642cc2d390889097c98e2bf9b6bf249a0d1374810e240425e0fd0494def1da5f3cc9020a11afac6a3c92fe <nil>
2019/03/21 11:41:01 Created genesis
```

Now you can run the `./bin/dimos-get-block` script and pass the genesis hash to it or just add the flag `-current`:

```
Usage of ./bin/dimos-get-block:
  -current
        Get the current block
  -hash string
        The hash of the block
```

Once you create a few blocks with `./bin/test-block` you can check the validity of the blockchain by running the following command:

``` sh
./bin/dimos-check-validity
```

This will go through all of your blocks and verify them.

#### Warning

This software is far from complete and improved every day.
