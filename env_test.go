package env

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	suite.Suite
}

func (s *ParserTestSuite) TestParseStruct() {

	os.Setenv("BOOL_FIELD", "true")
	os.Setenv("STR_FIELD", "strFieldValue")
	os.Setenv("INT_FIELD", "123")
	os.Setenv("INT64_FIELD", "9223372036854775807")
	os.Setenv("SUB_STR_FIELD", "subStrFieldValue")

	type SubConfig struct {
		SubStrField string `env:"SUB_STR_FIELD"`
	}

	type Config struct {
		BoolField bool   `env:"BOOL_FIELD"`
		StrField  string `env:"STR_FIELD"`
		// IntField       int    `env:"INT_FIELD"`
		// Int64Field     int64  `env:"INT64_FIELD"`
		SubConfigField SubConfig
	}

	var c Config

	err := Parse(&c)

	fmt.Printf("config: %+v\n", c)
	fmt.Printf("err: %s\n", err)

	assert.Nil(s.T(), err)

}

func TestParserTestSuite(t *testing.T) {
	suite.Run(t, new(ParserTestSuite))
}
