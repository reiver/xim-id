package optuint64

import (
	"errors"
)

var (
	errNilReceiver = errors.New("optuint64: nil receiver")
	errNilSource   = errors.New("optuint64: nil source")
	errNothing     = errors.New("optuint64: nothing")
)
