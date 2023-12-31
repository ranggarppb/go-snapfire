package types

type Publisher interface {
	Start(chan<- Message) error
}
