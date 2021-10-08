package model

import "strconv"

type ErrNotFound struct {
	ErrCode int
}

func (e *ErrNotFound) Error() string {
	str := strconv.Itoa(e.ErrCode)
	return str
}
