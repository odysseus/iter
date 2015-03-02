package iter

type Able interface {
	Next() int
	HasNext() bool
}
