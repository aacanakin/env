package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TagOptionsTestSuite struct {
	suite.Suite
}

func (s *TagOptionsTestSuite) TestParseTagEmpty() {
	tag := ""
	name, opts := parseTag(tag)

	assert.Empty(s.T(), name)
	assert.Empty(s.T(), opts)
}

func (s *TagOptionsTestSuite) TestParseTagOptsEmpty() {
	tag := "SAMPLE_ENV_VAR"
	name, opts := parseTag(tag)

	assert.Empty(s.T(), opts)
	assert.Equal(s.T(), tag, name)
}

func (s *TagOptionsTestSuite) TestParseTagOptsNonEmpty() {
	tag := "SAMPLE_ENV_VAR,omitempty,file"
	name, opts := parseTag(tag)

	assert.Equal(s.T(), tagOptions("omitempty,file"), opts)
	assert.Equal(s.T(), "SAMPLE_ENV_VAR", name)
}

func (s *TagOptionsTestSuite) TestTagContainsFirst() {
	tag := "SAMPLE_ENV_VAR,omitempty,file"
	_, opts := parseTag(tag)

	assert.True(s.T(), opts.Contains("omitempty"))
}

func (s *TagOptionsTestSuite) TestTagContainsSecond() {
	tag := "SAMPLE_ENV_VAR,omitempty,file"
	_, opts := parseTag(tag)

	assert.True(s.T(), opts.Contains("file"))
}

func (s *TagOptionsTestSuite) TestTagNotContains() {
	tag := "SAMPLE_ENV_VAR,omitempty,file"
	_, opts := parseTag(tag)

	assert.False(s.T(), opts.Contains("notcontains"))
}

func (s *TagOptionsTestSuite) TestTagNotContainsEmpty() {
	tag := "SAMPLE_ENV_VAR"
	_, opts := parseTag(tag)

	assert.False(s.T(), opts.Contains("notcontains"))
}

func TestTagOptionsTestSuite(t *testing.T) {
	suite.Run(t, new(TagOptionsTestSuite))
}
