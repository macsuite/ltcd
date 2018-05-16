rpctest
=======

[![Build Status](http://img.shields.io/travis/macsuite/macd.svg)](https://travis-ci.org/macsuite/macd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/macsuite/macd/integration/rpctest)

Package rpctest provides a macd-specific RPC testing harness crafting and
executing integration tests by driving a `macd` instance via the `RPC`
interface. Each instance of an active harness comes equipped with a simple
in-memory HD wallet capable of properly syncing to the generated chain,
creating new addresses, and crafting fully signed transactions paying to an
arbitrary set of outputs.

This package was designed specifically to act as an RPC testing harness for
`macd`. However, the constructs presented are general enough to be adapted to
any project wishing to programmatically drive a `macd` instance of its
systems/integration tests.

## Installation and Updating

```bash
$ go get -u github.com/macsuite/macd/integration/rpctest
```

## License

Package rpctest is licensed under the [copyfree](http://copyfree.org) ISC
License.

=======
rpctest
=======

[![Build Status](http://img.shields.io/travis/macsuite/macd.svg)]
(https://travis-ci.org/macsuite/macd) [![ISC License]
(http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)]
(http://godoc.org/github.com/macsuite/macd/rpctest)

Package rpctest provides a macd-specific RPC testing harness crafting and
executing integration tests by driving a `macd` instance via the `RPC`
interface. Each instance of an active harness comes equipped with a simple
in-memory HD wallet capable of properly syncing to the generated chain,
creating new addresses, and crafting fully signed transactions paying to an
arbitrary set of outputs. 

This package was designed specifically to act as an RPC testing harness for
`macd`. However, the constructs presented are general enough to be adapted to
any project wishing to programmatically drive a `macd` instance of its
systems/integration tests. 

## Installation and Updating

```bash
$ go get -u github.com/macsuite/macd/rpctest
```

## License


Package rpctest is licensed under the [copyfree](http://copyfree.org) ISC
License.

