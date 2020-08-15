package env

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/fatih/structs"
)

// FieldParser represents a generic field parser
type FieldParser func(field *structs.Field, val string) error

func parseBool(field *structs.Field, val string) error {
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return err
	}
	return field.Set(boolVal)
}

func parseString(field *structs.Field, val string) error {
	return field.Set(val)
}

func parseInt(bitSize int) FieldParser {
	return func(field *structs.Field, val string) error {

		if field.Kind() == reflect.Int {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return err
			}
			return field.Set(intVal)
		}

		intVal, err := strconv.ParseInt(val, 10, bitSize)
		if err != nil {
			return err
		}
		return field.Set(intVal)
	}
}

func parseUint(bitSize int) FieldParser {
	return func(field *structs.Field, val string) error {
		uintVal, err := strconv.ParseUint(val, 10, bitSize)
		if err != nil {
			return err
		}

		switch field.Kind() {
		case reflect.Uint8:
			return field.Set(uint8(uintVal))
		case reflect.Uint16:
			return field.Set(uint16(uintVal))
		case reflect.Uint32:
			return field.Set(uint32(uintVal))
		case reflect.Uint64:
			return field.Set(uintVal)
		default:
			field.Set(0)
			return fmt.Errorf("Unsupported type while parsing uint kind: %s", field.Kind())
		}

	}
}

func parseFloat(bitSize int) FieldParser {
	return func(field *structs.Field, val string) error {
		floatVal, err := strconv.ParseFloat(val, bitSize)
		if err != nil {
			return err
		}

		switch field.Kind() {
		case reflect.Float32:
			return field.Set(float32(floatVal))
		case reflect.Float64:
			return field.Set(float64(floatVal))
		default:
			return fmt.Errorf("Unsupported type while parsing float kind: %s", field.Kind())
		}
	}
}

func parseComplex(bitSize int) FieldParser {
	return func(field *structs.Field, val string) error {
		complexVal, err := strconv.ParseComplex(val, bitSize)
		if err != nil {
			return err
		}

		switch field.Kind() {
		case reflect.Complex64:
			return field.Set(complex64(complexVal))
		case reflect.Complex128:
			return field.Set(complex128(complexVal))
		default:
			return fmt.Errorf("Unsupported type while parsing complex kind: %s", field.Kind())
		}
	}
}

// Mapper represents a map of the type of primitive value parsers
type Mapper map[reflect.Kind]func(field *structs.Field, val string) error

var mappers = Mapper{
	reflect.String:     parseString,
	reflect.Bool:       parseBool,
	reflect.Int:        parseInt(32),
	reflect.Int8:       parseInt(8),
	reflect.Int16:      parseInt(16),
	reflect.Int32:      parseInt(32),
	reflect.Int64:      parseInt(64),
	reflect.Uint:       parseUint(32),
	reflect.Uint8:      parseUint(8),
	reflect.Uint16:     parseUint(16),
	reflect.Uint32:     parseUint(32),
	reflect.Uint64:     parseUint(64),
	reflect.Float32:    parseFloat(32),
	reflect.Float64:    parseFloat(64),
	reflect.Complex64:  parseComplex(64),
	reflect.Complex128: parseComplex(128),
}

// Parser represents a parser interface that contains a parse function
type Parser interface {
	Parse(s interface{}) error
}

func parseField(field *structs.Field) error {

	if field.Kind() == reflect.Struct {
		for _, f := range field.Fields() {
			return parseField(f)
		}
	}

	tagVal := field.Tag("env")
	defaultTagVal := fmt.Sprintf("%s", strings.Join(camelToSlice(field.Name()), "_"))
	if tagVal == "" {
		tagVal = defaultTagVal
	}

	tag, _ := parseTag(tagVal)

	envVar := strings.ToUpper(tag)
	envVal, _ := os.LookupEnv(envVar) // TODO: Validate required handling

	fx := mappers[field.Kind()]
	if fx != nil {
		err := fx(field, envVal)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Unsupported type '%s' for field '%s'", field.Kind(), field.Name())
	}

	return nil
}

// Parse parses struct and maps environment variables to struct values
func Parse(c interface{}) error {

	isStruct := structs.IsStruct(c)
	if !isStruct {
		return errors.New("Types other than structs are not allowed")
	}

	s := structs.New(c)

	for _, field := range s.Fields() {
		err := parseField(field)
		if err != nil {
			return err
		}
	}

	return nil
}
