package response

type Error struct {
	Message string
}

// error
func NewError(err error) *Error {
	return &Error{ Message: err.Error() }
}
