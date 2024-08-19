package core

import (
	"encoding/gob"
	"fmt"
	"math/rand"

	"github.com/bishalbera/modular-blockchain/crypto"
	"github.com/bishalbera/modular-blockchain/types"
)

type TxType byte

const (
	TxTypeCollection TxType = iota
	TxTypeMint
)

type CollectionTx struct {
	Fee      int64
	MetaData []byte
}

type MintTx struct {
	Fee             int64
	NFT             types.Hash
	Collection      types.Hash
	MetaData        []byte
	CollectionOwner crypto.PublicKey
	Signature       crypto.Signature
}

type Transaction struct {
	// Only used for native NFT logic
	TxInner any
	// Any arbitary data for the vm
	Data      []byte
	From      crypto.PublicKey
	Signature *crypto.Signature
	To        crypto.PublicKey
	Value     uint64
	Nonce     int64
	// cached version of the tx data hash
	hash types.Hash
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		Data:  data,
		Nonce: rand.Int63n(1000000000000000),
	}
}

func (tx *Transaction) Hash(hasher Hasher[*Transaction]) types.Hash {
	if tx.hash.IsZero() {
		tx.hash = hasher.Hash(tx)
	}

	return tx.hash
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}
	tx.From = privKey.PublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}
	if !tx.Signature.Verify(tx.From, tx.Data) {
		return fmt.Errorf("invalid transaction signature")
	}
	return nil
}

func (tx *Transaction) Decode(dec Decoder[*Transaction]) error {
	return dec.Decode(tx)
}

func (tx *Transaction) Encode(enc Encoder[*Transaction]) error {
	return enc.Encode(tx)
}

func init() {
	gob.Register(CollectionTx{})
	gob.Register(MintTx{})
}
