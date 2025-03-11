package json

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Employee struct {
	Name    string
	Surname string
	Age     int
}

func TestJSON_Unmarshal(t *testing.T) {
	t.Parallel()
	expected := map[string]interface{}{
		"operation": "get",
		"key":       "example",
	}
	r := unmarshal(`{"operation": "get", "key": "example"}`)
	fmt.Println(r)
	assert.Equal(t, expected, r)
}

func TestJSON_Unmarshal_error(t *testing.T) {
	t.Parallel()
	var result any

	assert.NotPanics(t, func() {
		result = unmarshal(`{"operation": "get", `)
	}, "The code throw panic")
	assert.Nil(t, result, "The code does not return nil")
}

func TestJSON_Marshal(t *testing.T) {
	object := Employee{
		Name:    "SRE",
		Surname: "Inditex",
		Age:     30,
	}
	t.Parallel()
	r := marshal(object)
	fmt.Println(r)
	require.NotNil(t, r)
}

func TestJSON_Marshal_ScapeHtml(t *testing.T) {
	object := Employee{
		Name:    "SRE",
		Surname: "https://www.inditex.com/&lt;script&gt;alert('XSS')&lt;/script&gt;",
		Age:     30,
	}

	t.Parallel()
	r := marshal(object)
	fmt.Println(r)
	require.NotNil(t, r)
	if !strings.Contains(r, "https://www.inditex.com/&lt;script&gt;alert('XSS')&lt;/script&gt;") {
		t.Errorf("Result does not contain the required string")
	}
}

func TestJSON_Marshal_error(t *testing.T) {
	t.Parallel()
	var unsupportedValues = []any{
		math.NaN(),
	}

	var result string
	assert.NotPanics(t, func() {
		result = marshal(unsupportedValues[0])
	}, "The code throw panic")
	assert.Equal(t, "", result)
}
