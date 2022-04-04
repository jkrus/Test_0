package nats_subscr

import "github.com/pkg/errors"

const (
	errOrderValidateMsg = "order validate"
)

func ErrOrderValidate(w error) error {
	return errors.Wrap(w, errOrderValidateMsg)
}
