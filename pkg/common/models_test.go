package common

import (
	"io/ioutil"
	"regexp"
	"strings"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/deepmap/oapi-codegen/pkg/util"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

// Show that the generated code file has not been modified.
func TestCodegen(t *testing.T) {
	swagger, err := util.LoadSwagger("../../openapi/models.yaml")
	require.NoError(t, err)

	expected, err := codegen.Generate(swagger, "common", codegen.Options{
		GenerateTypes: true,
		SkipPrune:     true,
	})
	require.NoError(t, err)

	actual, err := ioutil.ReadFile("models.go")
	require.NoError(t, err)
	fileAsStr := string(actual)

	// Ignore unique versions/timestamps in the header and just compare everything after it.
	reg := regexp.MustCompilePOSIX(`.*DO NOT EDIT\.`)
	fileAsStr = strings.TrimSuffix(reg.ReplaceAllString(fileAsStr, ""), "\n")
	expected = reg.ReplaceAllString(expected, "")

	assert.Equal(t, expected, fileAsStr)
}
