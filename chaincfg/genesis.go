// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/macsuite/macd/chaincfg/chainhash"
	"github.com/macsuite/macd/wire"
)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
    0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x40, 0x64, 0x6c, 0x6f, 0x20, 0x73, 0x72, 0x61, 0x65,
    0x79, 0x20, 0x30, 0x35, 0x20, 0x77, 0x6f, 0x6e, 0x20, 0x73, 0x69, 0x20, 0x2c, 0x72, 0x65, 0x74,
    0x68, 0x67, 0x69, 0x66, 0x20, 0x73, 0x75, 0x6f, 0x6d, 0x61, 0x66, 0x20, 0x2c, 0x65, 0x6b, 0x73,
    0x61, 0x4d, 0x20, 0x79, 0x72, 0x6e, 0x65, 0x48, 0x20, 0x34, 0x31, 0x30, 0x32, 0x2f, 0x6e, 0x61,
    0x4a, 0x2f, 0x36, 0x30, 0x20, 0x6c, 0x65, 0x67, 0x65, 0x69, 0x70, 0x73, 0x73, 0x65, 0x67, 0x61,
    0x54, 0x20, 0x72, 0x65, 0x44,

			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
				0x41, 0x4, 0x1, 0x84, 0x71, 0xf, 0xa6, 0x89,
				0xad, 0x50, 0x23, 0x69, 0xc, 0x80, 0xf3, 0xa4,
				0x9c, 0x8f, 0x13, 0xf8, 0xd4, 0x5b, 0x8c, 0x85,
				0x7f, 0xbc, 0xbc, 0x8b, 0xc4, 0xa8, 0xe4, 0xd3,
				0xeb, 0x4b, 0x10, 0xf4, 0xd4, 0x60, 0x4f, 0xa0,
				0x8d, 0xce, 0x60, 0x1a, 0xaf, 0xf, 0x47, 0x2,
				0x16, 0xfe, 0x1b, 0x51, 0x85, 0xb, 0x4a, 0xcf,
				0x21, 0xb1, 0x79, 0xc4, 0x50, 0x70, 0xac, 0x7b,
				0x3, 0xa9, 0xac,
			},
		},
	},
	LockTime: 0,
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var genesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x72, 0x31, 0xe8, 0x03, 0x77, 0x96, 0xde, 0x41,
 0x2d, 0x35, 0x0e, 0x50, 0x9a, 0x56, 0xed, 0x4f,
 0xcc, 0xdd, 0xb2, 0x0c, 0xfd, 0xe1, 0xde, 0xbf,
 0x1c, 0x47, 0xa5, 0xce, 0x9b, 0x87, 0x1f, 0x6a,
 })

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x5c, 0x4a, 0x38, 0xbc, 0xc3, 0xd0, 0xa4, 0xa4,
 0x6f, 0x01, 0x86, 0xb2, 0x88, 0x2a, 0x80, 0x1c,
 0x3c, 0x19, 0xe8, 0xd9, 0x0e, 0x9d, 0x29, 0x66,
 0x14, 0xe7, 0xf3, 0x63, 0x10, 0xe4, 0xa9, 0x36,
 })

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: genesisMerkleRoot,        // 36a9e41063f3e71466299d0ed9e8193c1c802a88b286016fa4a4d0c3bc384a5c
		Timestamp:  time.Unix(1389040865, 0),
		Bits:       0x1e0ffff0,
		Nonce:      3716037,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// regTestGenesisHash is the hash of the first block in the block chain for the
// regression test network (genesis block).
var regTestGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x2b, 0x91, 0x37, 0xb1, 0xd1, 0xd5, 0x99, 0x0f,
 0x6e, 0x69, 0x9f, 0x18, 0x72, 0xf2, 0x98, 0x9a,
 0x1e, 0x8c, 0xa1, 0xc2, 0xde, 0xcd, 0xa4, 0x08,
 0x33, 0xbf, 0x65, 0x88, 0x69, 0x4f, 0x21, 0x62,
})

// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the regression test network.  It is the same as the merkle root for
// the main network.
var regTestGenesisMerkleRoot = genesisMerkleRoot

// regTestGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the regression test network.
var regTestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: regTestGenesisMerkleRoot, // 36a9e41063f3e71466299d0ed9e8193c1c802a88b286016fa4a4d0c3bc384a5c
		Timestamp:  time.Unix(1296688602, 0), //
		Bits:       0x207fffff,               //
		Nonce:      0,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// testNet4GenesisHash is the hash of the first block in the block chain for the
// test network (version 4).
var testNet4GenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x5f, 0x89, 0x7e, 0x54, 0x55, 0xf5, 0xde, 0xdb,
 0x02, 0x46, 0x53, 0x29, 0xe3, 0xc6, 0xcc, 0xe1,
 0xbf, 0x0a, 0x07, 0x0f, 0xbd, 0x36, 0xcc, 0x41,
 0x29, 0x9a, 0xc4, 0x1c, 0x48, 0x9c, 0x05, 0x72,
})

// testNet3GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 3).  It is the same as the merkle root
// for the main network.
var testNet3GenesisMerkleRoot = genesisMerkleRoot

// testNet4GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 4).  It is the same as the merkle root
// for the main network.
var testNet4GenesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x5c, 0x4a, 0x38, 0xbc, 0xc3, 0xd0, 0xa4, 0xa4,
 0x6f, 0x01, 0x86, 0xb2, 0x88, 0x2a, 0x80, 0x1c,
 0x3c, 0x19, 0xe8, 0xd9, 0x0e, 0x9d, 0x29, 0x66,
 0x14, 0xe7, 0xf3, 0x63, 0x10, 0xe4, 0xa9, 0x36,
})

// testNet4GenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 4).
var testNet4GenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},          // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: testNet4GenesisMerkleRoot, // 36a9e41063f3e71466299d0ed9e8193c1c802a88b286016fa4a4d0c3bc384a5c
		Timestamp:  time.Unix(1473357600, 0),
		Bits:       0x1e0ffff0,
		Nonce:      5653466,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// simNetGenesisHash is the hash of the first block in the block chain for the
// simulation test network.
var simNetGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x2b, 0x91, 0x37, 0xb1, 0xd1, 0xd5, 0x99, 0x0f,
 0x6e, 0x69, 0x9f, 0x18, 0x72, 0xf2, 0x98, 0x9a,
 0x1e, 0x8c, 0xa1, 0xc2, 0xde, 0xcd, 0xa4, 0x08,
 0x33, 0xbf, 0x65, 0x88, 0x69, 0x4f, 0x21, 0x62,
})

// simNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the simulation test network.  It is the same as the merkle root for
// the main network.
var simNetGenesisMerkleRoot = genesisMerkleRoot

// simNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the simulation test network.
var simNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: simNetGenesisMerkleRoot,  // 36a9e41063f3e71466299d0ed9e8193c1c802a88b286016fa4a4d0c3bc384a5c
		Timestamp:  time.Unix(1296688602, 0), //
		Bits:       0x207fffff,               //
		Nonce:      0,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}
