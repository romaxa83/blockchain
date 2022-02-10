package blockchain

// выход, хранит данные (монеты)
type TxOutput struct {
	Value  int
	PubKey string // блокируется ключом
}

// вход, ссылается на предыдущий выход
type TxInput struct {
	ID  []byte // идентификатор транзакции выхода
	Out int    // хранит индекс выхода данной транзакции
	Sig string
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
