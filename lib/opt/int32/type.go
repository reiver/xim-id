package optint32

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// OptInt32 is an implementation of a string option-type.
type OptInt32 struct {
	value int32
	loaded bool
}

// Nothing returns a string option-type that has no value.
func Nothing() OptInt32 {
	return OptInt32{}
}

// Something return a string option-type with a value.
func Something(value int32) OptInt32 {
	return OptInt32{
		loaded: true,
		value: value,
	}
}

// GoString makes OptInt32 fit the fmt.GoStringer interface.
func (receiver OptInt32) GoString() string {
	if Nothing() == receiver {
		return "optstr.Nothing()"
	}

	return fmt.Sprintf("optstr.Something(%q)", receiver.value)
}

// MarshalText makes OptInt32 fit the encoding.TextMarshaler interface.
func (receiver OptInt32) MarshalText() (text []byte, err error) {
	if Nothing() == receiver {
		return nil, errNothing
	}

	return []byte(receiver.String()), nil
}

func (receiver OptInt32) Return() (int32, bool) {
	if Nothing() == receiver {
		return 0, false
	}

	return receiver.value, true
}

// Scan makes OptInt32 fit the sql.Scan interfaces.
func (receiver *OptInt32) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch casted := src.(type) {
	case OptInt32:
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

// Set makes OptInt32 fit the flag.Value interfaces.
func (receiver *OptInt32) Set(value string) error {
	if nil == receiver {
		return errNilReceiver
	}

	i64, err := strconv.ParseInt(value, 10, 0)
	if nil != err {
		return err
	}
	var i int32 = int32(i64)

	*receiver = Something(i)
	return nil
}

// String makes OptInt32 fit the fmt.Stinger, flag.Value interfaces.
func (receiver OptInt32) String() string {
	if Nothing() == receiver {
		return ""
	}

	return strconv.FormatInt(int64(receiver.value), 10)
}

// UnmarshalText makes OptInt32 fit the encoding.TextUnmarshaler interface.
func (receiver *OptInt32) UnmarshalText(text []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == text {
		return errNilSource
	}

	receiver.Set(string(text))
	return nil
}

// Value makes OptInt32 fit the database/sql/driver.Valuer interface.
func (receiver OptInt32) Value() (driver.Value, error) {
	if Nothing() == receiver {
		return nil, errNothing
	}

	return receiver.value, nil
}
