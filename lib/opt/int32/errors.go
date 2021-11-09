package optint32

import (
	"errors"
)

var (
	errNilReceiver = errors.New("optint32: nil receiver")
	errNilSource   = errors.New("optint32: nil source")
	errNothing     = errors.New("optint32: nothing")
)
