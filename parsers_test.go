package env

import (
	"os"
	"testing"

	"github.com/fatih/structs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ParsersTestSuite struct {
	suite.Suite
}

func (s *ParsersTestSuite) TestParseBool() {
	type config struct {
		BoolField bool
	}

	var c config
	sx := structs.New(&c)

	val := "true"
	err := parseBool(sx.Field("BoolField"), val)

	assert.Nil(s.T(), err)
	assert.True(s.T(), c.BoolField)
}

func (s *ParsersTestSuite) TestParseBoolErr() {
	type config struct {
		BoolField bool
	}

	var c config
	sx := structs.New(&c)

	val := "tr"
	err := parseBool(sx.Field("BoolField"), val)

	assert.NotNil(s.T(), err)
	assert.False(s.T(), c.BoolField)
}

func (s *ParsersTestSuite) TestParseString() {
	type config struct {
		StrField string
	}

	var c config
	sx := structs.New(&c)

	val := "sampleValue"
	err := parseString(sx.Field("StrField"), val)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), val, c.StrField)
}

func (s *ParsersTestSuite) TestParseInt() {
	type config struct {
		IntField int
	}

	var c config
	sx := structs.New(&c)

	val := "123"
	err := parseInt(32)(sx.Field("IntField"), val)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 123, c.IntField)
}

func (s *ParsersTestSuite) TestParseIntInvalidTypeErr() {
	type config struct {
		IntField int
	}

	var c config
	sx := structs.New(&c)

	val := "123asd"
	err := parseInt(32)(sx.Field("IntField"), val)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), 0, c.IntField)
}

func (s *ParsersTestSuite) TestParseIntSetErr() {
	type config struct {
		IntField int
		StrField string
	}

	var c config
	sx := structs.New(&c)

	val := "123"
	err := parseInt(32)(sx.Field("StrField"), val)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), 0, c.IntField)
}

func (s *ParsersTestSuite) TestParseIntOverflowErr() {
	type config struct {
		Int8Field int8
	}

	var c config
	sx := structs.New(&c)

	val := "123123123123123123123123123123"
	err := parseInt(128)(sx.Field("Int8Field"), val)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), int8(0), c.Int8Field)
}

func (s *ParsersTestSuite) TestParseUint8() {
	type config struct {
		UintField uint8
	}

	var c config
	sx := structs.New(&c)

	val := "255"
	err := parseUint(8)(sx.Field("UintField"), val)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), uint8(255), c.UintField)
}

func (s *ParsersTestSuite) TestParseUint16() {
	type config struct {
		UintField uint16
	}

	var c config
	sx := structs.New(&c)

	val := "65535"
	err := parseUint(16)(sx.Field("UintField"), val)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), uint16(65535), c.UintField)
}

func (s *ParsersTestSuite) TestParseUint32() {
	type config struct {
		UintField uint32
	}

	var c config
	sx := structs.New(&c)

	val := "4294967295"
	err := parseUint(32)(sx.Field("UintField"), val)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), uint32(4294967295), c.UintField)
}

func (s *ParsersTestSuite) TestParseUint64() {
	type config struct {
		UintField uint64
	}

	var c config
	sx := structs.New(&c)

	val := "18446744073709551615"
	err := parseUint(64)(sx.Field("UintField"), val)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), uint64(18446744073709551615), c.UintField)
}

func (s *ParsersTestSuite) TestParseUintInvalidTypeErr() {
	type config struct {
		UintField bool
	}

	var c config
	sx := structs.New(&c)

	val := "18446744073709551615"
	err := parseUint(64)(sx.Field("UintField"), val)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), false, c.UintField)
}

func (s *ParsersTestSuite) TestParseUintInvalidValueErr() {
	type config struct {
		UintField uint
	}

	var c config
	sx := structs.New(&c)

	val := "18text"
	err := parseUint(32)(sx.Field("UintField"), val)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), uint(0), c.UintField)
}

func (s *ParsersTestSuite) TestParseFloat32() {
	type config struct {
		FloatField float32
	}

	var c config
	sx := structs.New(&c)

	val := "184.5"
	err := parseFloat(32)(sx.Field("FloatField"), val)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), float32(184.5), c.FloatField)
}

func (s *ParsersTestSuite) TestParseFloat64() {
	type config struct {
		FloatField float64
	}

	var c config
	sx := structs.New(&c)

	val := "1.7976931348623157"
	err := parseFloat(64)(sx.Field("FloatField"), val)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), float64(1.7976931348623157), c.FloatField)
}

func (s *ParsersTestSuite) TestParseFloatInvalidType() {
	type config struct {
		FloatField bool
	}

	var c config
	sx := structs.New(&c)

	val := "1.7976931348623157"
	err := parseFloat(64)(sx.Field("FloatField"), val)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), false, c.FloatField)
}

func (s *ParsersTestSuite) TestParseFloatInvalidVal() {
	type config struct {
		FloatField float32
	}

	var c config
	sx := structs.New(&c)

	val := "1.7976931348623157text"
	err := parseFloat(32)(sx.Field("FloatField"), val)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), float32(0), c.FloatField)
}

func (s *ParsersTestSuite) TestParseComplex64() {
	type config struct {
		ComplexField complex64
	}

	var c config
	sx := structs.New(&c)

	val := "1.7976931348623157"
	err := parseComplex(64)(sx.Field("ComplexField"), val)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), complex64(1.7976931+0i), c.ComplexField)
}

func (s *ParsersTestSuite) TestParseComplex128() {
	type config struct {
		ComplexField complex128
	}

	var c config
	sx := structs.New(&c)

	val := "1.7976931348623157e+308+1.7976931348623157e+308i"
	err := parseComplex(128)(sx.Field("ComplexField"), val)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), complex128(1.7976931348623157e+308+1.7976931348623157e+308i), c.ComplexField)
}

func (s *ParsersTestSuite) TestParseComplexInvalidType() {
	type config struct {
		ComplexField string
	}

	var c config
	sx := structs.New(&c)

	val := "1.7976931348623157e+308+1.7976931348623157e+308i"
	err := parseComplex(128)(sx.Field("ComplexField"), val)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), "", c.ComplexField)
}

func (s *ParsersTestSuite) TestParseComplexInvalidVal() {
	type config struct {
		ComplexField complex128
	}

	var c config
	sx := structs.New(&c)

	val := "text"
	err := parseComplex(128)(sx.Field("ComplexField"), val)

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), complex128(0+0i), c.ComplexField)
}

func (s *ParsersTestSuite) TestParseFieldDefault() {
	val := "strFieldVal"
	os.Setenv("STR_FIELD", val)
	defer os.Clearenv()

	type config struct {
		StrField string
	}

	var c config
	sx := structs.New(&c)

	err := parseField(sx.Field("StrField"))

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), val, c.StrField)
}

func (s *ParsersTestSuite) TestParseFieldCustomTag() {
	val := "strFieldVal"
	os.Setenv("STR_CUSTOM_FIELD", val)
	defer os.Clearenv()

	type config struct {
		StrField string `env:"STR_CUSTOM_FIELD"`
	}

	var c config
	sx := structs.New(&c)

	err := parseField(sx.Field("StrField"))

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), val, c.StrField)
}

func (s *ParsersTestSuite) TestParseFieldUnsupportedType() {
	val := "strFieldVal"
	os.Setenv("STR_FIELD", val)
	defer os.Clearenv()

	type config struct {
		StrField *string
	}

	var c config
	sx := structs.New(&c)

	err := parseField(sx.Field("StrField"))

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), (*string)(nil), c.StrField)
}

func (s *ParsersTestSuite) TestParseFieldParserErr() {
	val := "strFieldVal"
	os.Setenv("INT_FIELD", val)
	defer os.Clearenv()

	type config struct {
		IntField int
	}

	var c config
	sx := structs.New(&c)

	err := parseField(sx.Field("IntField"))

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), 0, c.IntField)
}

func (s *ParsersTestSuite) TestParseFieldStructDefault() {
	val := "123"
	os.Setenv("SUB_INT_FIELD", val)
	defer os.Clearenv()

	type subconfig struct {
		SubIntField int
	}

	type config struct {
		SubConfig subconfig
	}

	var c config
	sx := structs.New(&c)

	err := parseField(sx.Field("SubConfig"))

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 123, c.SubConfig.SubIntField)
}

func (s *ParsersTestSuite) TestParseFieldStructEmbeddedDefault() {
	val := "123"
	os.Setenv("SUB_INT_FIELD", val)
	defer os.Clearenv()

	type subconfig struct {
		SubIntField int
	}

	type config struct {
		subconfig
	}

	var c config
	sx := structs.New(&c)

	err := parseField(sx.Field("SubIntField"))

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 123, c.SubIntField)
}

func (s *ParsersTestSuite) TestParseDefault() {
	strVal := "strVal"
	intVal := "123"

	os.Setenv("STR_FIELD", strVal)
	os.Setenv("INT_FIELD", intVal)
	os.Setenv("SUB_BOOL_FIELD", "t")
	defer os.Clearenv()

	type subconfig struct {
		SubBoolField bool
	}

	type config struct {
		StrField string
		IntField int
		subconfig
	}

	var c config
	err := Parse(&c)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), strVal, c.StrField)
	assert.Equal(s.T(), 123, c.IntField)
	assert.True(s.T(), c.SubBoolField)
}

func (s *ParsersTestSuite) TestParseCustomTag() {
	strVal := "strVal"
	intVal := "123"

	os.Setenv("STR_CUSTOM_FIELD", strVal)
	os.Setenv("INT_FIELD", intVal)
	os.Setenv("SUB_BOOL_CUSTOM_FIELD", "t")
	defer os.Clearenv()

	type subconfig struct {
		SubBoolField bool `env:"SUB_BOOL_CUSTOM_FIELD"`
	}

	type config struct {
		StrField string `env:"STR_CUSTOM_FIELD"`
		IntField int
		subconfig
	}

	var c config
	err := Parse(&c)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), strVal, c.StrField)
	assert.Equal(s.T(), 123, c.IntField)
	assert.True(s.T(), c.SubBoolField)
}

func (s *ParsersTestSuite) TestParseNonStructErr() {
	type config int

	var c config

	err := Parse(&c)

	assert.NotNil(s.T(), err)
}

func (s *ParsersTestSuite) TestParseFieldErr() {

	os.Setenv("STR_FIELD", "strVal")
	os.Setenv("BOOL_FIELD", "text")

	type config struct {
		StrField  string
		BoolField bool
	}

	var c config
	err := Parse(&c)

	assert.NotNil(s.T(), err)
}

func (s *ParsersTestSuite) TestParseSubStructFieldErr() {

	os.Setenv("STR_FIELD", "strVal")
	os.Setenv("SUB_BOOL_FIELD", "text")

	type subconfig struct {
		SubBoolField bool `env:"SUB_BOOL_FIELD"`
	}

	type config struct {
		StrField string
		subconfig
	}

	var c config
	err := Parse(&c)

	assert.NotNil(s.T(), err)
}

func TestParsersTestSuite(t *testing.T) {
	suite.Run(t, new(ParsersTestSuite))
}
