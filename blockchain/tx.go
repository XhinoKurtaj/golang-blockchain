package blockchain

type TxOutputs struct {
	Value  int
	PubKey string
}

type TxInputs struct {
	ID  []byte
	Out int
	Sig string
}

func (in *TxInputs) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutputs) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
