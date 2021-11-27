package store

const (
	NotFound = 0
)

type Error struct {
	Code int
	Msg  string
}

func (e *Error) Error() *Error {
	return e
}
