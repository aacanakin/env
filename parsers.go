package env

import (
	"reflect"
	"strconv"
)

func parseString(v reflect.Value, val string) error {
	v.SetString(val)
	return nil
}

func parseInt(v reflect.Value, val string, bitSize int) error {
	intVal, err := strconv.ParseInt(val, 10, bitSize)
	if err != nil {
		return err
	}

	v.SetInt(intVal)
	return nil
}

func parseUInt(v reflect.Value, val string, bitSize int) error {
	intVal, err := strconv.ParseUint(val, 10, bitSize)
	if err != nil {
		return err
	}
	v.SetUint(intVal)
	return nil
}

func parseFloat(v reflect.Value, val string, bitSize int) error {
	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err != nil {
		return err
	}
	v.SetFloat(floatVal)
	return nil
}

func parseComplex(v reflect.Value, val string, bitSize int) error {
	complexVal, err := strconv.ParseComplex(val, bitSize)
	if err != nil {
		return err
	}
	v.SetComplex(complexVal)
	return nil
}

// Mapper represents the type of primitive value parsers
type Mapper map[reflect.Kind]func(v reflect.Value, val string) error

var mappers Mapper = Mapper{
	reflect.Bool: func(v reflect.Value, val string) error {
		boolVal, err := strconv.ParseBool(val)
		if err != nil {
			return err
		}
		v.SetBool(boolVal)
		return nil
	},
	reflect.String: func(v reflect.Value, val string) error {
		return parseString(v, val)
	},
	reflect.Int: func(v reflect.Value, val string) error {
		return parseInt(v, val, 32)
	},
	reflect.Int8: func(v reflect.Value, val string) error {
		return parseInt(v, val, 8)
	},
	reflect.Int16: func(v reflect.Value, val string) error {
		return parseInt(v, val, 16)
	},
	reflect.Int32: func(v reflect.Value, val string) error {
		return parseInt(v, val, 32)
	},
	reflect.Int64: func(v reflect.Value, val string) error {
		return parseInt(v, val, 64)
	},
	reflect.Uint: func(v reflect.Value, val string) error {
		return parseUInt(v, val, 32)
	},
	reflect.Uint8: func(v reflect.Value, val string) error {
		return parseUInt(v, val, 8)
	},
	reflect.Uint16: func(v reflect.Value, val string) error {
		return parseUInt(v, val, 16)
	},
	reflect.Uint32: func(v reflect.Value, val string) error {
		return parseUInt(v, val, 32)
	},
	reflect.Uint64: func(v reflect.Value, val string) error {
		return parseUInt(v, val, 64)
	},
	reflect.Float32: func(v reflect.Value, val string) error {
		return parseFloat(v, val, 32)
	},
	reflect.Float64: func(v reflect.Value, val string) error {
		return parseFloat(v, val, 64)
	},
	reflect.Complex64: func(v reflect.Value, val string) error {
		return parseComplex(v, val, 64)
	},
	reflect.Complex128: func(v reflect.Value, val string) error {
		return parseComplex(v, val, 128)
	},
}
