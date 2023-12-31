package types

type Storage interface {
	Push([]byte) (int, error)
	Get(int) ([]byte, error)
}
