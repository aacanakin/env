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
	field.Set(boolVal)
	return nil
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

func parseUInt(bitSize int) FieldParser {
	return func(field *structs.Field, val string) error {
		uintVal, err := strconv.ParseUint(val, 10, bitSize)
		if err != nil {
			return err
		}
		return field.Set(uintVal)
	}
}

func parseFloat(bitSize int) FieldParser {
	return func(field *structs.Field, val string) error {
		floatVal, err := strconv.ParseFloat(val, bitSize)
		if err != nil {
			return err
		}
		return field.Set(floatVal)
	}
}

func parseComplex(bitSize int) FieldParser {
	return func(field *structs.Field, val string) error {
		complexVal, err := strconv.ParseComplex(val, bitSize)
		if err != nil {
			return err
		}
		field.Set(complexVal)
		return nil
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
	reflect.Uint:       parseUInt(32),
	reflect.Uint8:      parseUInt(8),
	reflect.Uint16:     parseUInt(16),
	reflect.Uint32:     parseUInt(32),
	reflect.Uint64:     parseUInt(64),
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
		// fmt.Printf("<Field name=%s, kind=%s tag=%s />\n", field.Name(), field.Kind().String(), field.Tag(`env`))

		err := parseField(field)
		if err != nil {
			return err
		}
	}

	return nil
}
