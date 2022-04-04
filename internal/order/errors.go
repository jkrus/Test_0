package order

import "github.com/pkg/errors"

const (
	errOrderDecodeMsg   = "order decode"
	errOrderNotFoundMsg = "order not found by uid"
)

func ErrOrderNotFound(w error) error {
	return errors.Wrap(w, errOrderNotFoundMsg)
}

func ErrOrderDecode(w error) error {
	return errors.Wrap(w, errOrderDecodeMsg)
}
