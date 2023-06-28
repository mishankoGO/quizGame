package reader

type Reader interface {
	ReadLine() ([]string, error)
	Close() error
}
