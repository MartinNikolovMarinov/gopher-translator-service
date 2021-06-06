package errors

import (
	"fmt"
)

type HttpErr struct {
	Status  int    `json:"status"`
	Details string `json:"details"`
	Err     error  `json:"-"`
}

func NewHttpErr(status int, details string, err error) error {
	return &HttpErr{Status: status, Details: details, Err: err}
}

func (e *HttpErr) Error() string {
	return fmt.Sprintf("Status: %d Detals: %s Err: %+v", e.Status, e.Details, e.Err)
}
