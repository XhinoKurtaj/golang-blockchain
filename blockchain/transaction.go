package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

type Transaction struct {
	ID      []byte
	Inputs  []TxInputs
	Outputs []TxOutputs
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

//Coin base transaction
func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coin to %s", to)
	}

	txin := TxInputs{[]byte{}, -1, data}
	txout := TxOutputs{100, to}

	tx := Transaction{nil, []TxInputs{txin}, []TxOutputs{txout}}
	tx.SetID()

	return &tx
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

func NewTransaction(from, to string, amount int, chain *BlockChain) *Transaction {
	var inputs []TxInputs
	var outputs []TxOutputs

	acc, validOutputs := chain.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("Error: not enough funds")
	}

	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		Handle(err)

		for _, out := range outs {
			input := TxInputs{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxOutputs{amount, to})

	if acc > amount {
		outputs = append(outputs, TxOutputs{acc - amount, from})
	}
	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}
