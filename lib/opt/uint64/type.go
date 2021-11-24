package optuint64

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// OptUint64 is an implementation of a string option-type.
type OptUint64 struct {
	value uint64
	loaded bool
}

// Nothing returns a string option-type that has no value.
func Nothing() OptUint64 {
	return OptUint64{}
}

// Something return a string option-type with a value.
func Something(value uint64) OptUint64 {
	return OptUint64{
		loaded: true,
		value: value,
	}
}

func (receiver OptUint64) ForcedReturn(value uint64) uint64 {
	if Nothing() == receiver {
		return value
	}

	return receiver.value
}

// GoString makes OptUint64 fit the fmt.GoStringer interface.
func (receiver OptUint64) GoString() string {
	if Nothing() == receiver {
		return "optstr.Nothing()"
	}

	return fmt.Sprintf("optstr.Something(%q)", receiver.value)
}

// MarshalText makes OptUint64 fit the encoding.TextMarshaler interface.
func (receiver OptUint64) MarshalText() (text []byte, err error) {
	if Nothing() == receiver {
		return nil, errNothing
	}

	return []byte(receiver.String()), nil
}

func (receiver OptUint64) Return() (uint64, bool) {
	if Nothing() == receiver {
		return 0, false
	}

	return receiver.value, true
}

// Scan makes OptUint64 fit the sql.Scan interfaces.
func (receiver *OptUint64) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch casted := src.(type) {
	case OptUint64:
		*receiver = casted
		return nil
	case string:
		return receiver.Set(casted)
	case []byte:
		s := string(casted)
		return receiver.Set(s)
	default:
		return fmt.Errorf("Cannot scan something of type %T into an %T.", src, *receiver)
	}
}

// Set makes OptUint64 fit the flag.Value interfaces.
func (receiver *OptUint64) Set(value string) error {
	if nil == receiver {
		return errNilReceiver
	}

	i64, err := strconv.ParseInt(value, 10, 0)
	if nil != err {
		return err
	}
	var i uint64 = uint64(i64)

	*receiver = Something(i)
	return nil
}

// String makes OptUint64 fit the fmt.Stinger, flag.Value interfaces.
func (receiver OptUint64) String() string {
	if Nothing() == receiver {
		return ""
	}

	return strconv.FormatInt(int64(receiver.value), 10)
}

// UnmarshalText makes OptUint64 fit the encoding.TextUnmarshaler interface.
func (receiver *OptUint64) UnmarshalText(text []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == text {
		return errNilSource
	}

	receiver.Set(string(text))
	return nil
}

// Value makes OptUint64 fit the database/sql/driver.Valuer interface.
func (receiver OptUint64) Value() (driver.Value, error) {
	if Nothing() == receiver {
		return nil, errNothing
	}

	return receiver.value, nil
}
