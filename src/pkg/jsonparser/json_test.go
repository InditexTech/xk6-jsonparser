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

func TestUnmarshal(t *testing.T) {
	t.Parallel()
	r := unmarshal(`{"operation": "get", "key": "example"}`)
	fmt.Println(r)
	require.NotNil(t, r)
}

func TestJSON_Unmarshal_should_not_return_error(t *testing.T) {
	t.Parallel()
	assert.NotPanics(t, func() { unmarshal(`{"operation": "get", `) }, "The code throw panic")
}

func TestJSON_Unmarshal_should_return_nil(t *testing.T) {
	t.Parallel()
	r := unmarshal(`{"operation": "get", `)
	fmt.Println(r)
	assert.Nil(t, r, "The code not return nil")
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

func TestJSON_Marshal_should_not_return_error(t *testing.T) {
	var unsupportedValues = []any{
		math.NaN(),
	}
	t.Parallel()
	assert.NotPanics(t, func() { marshal(unsupportedValues[0]) }, "The code throw panic")
}

func TestJSON_Marshal_should_return_empty_string(t *testing.T) {
	var unsupportedValues = []any{
		math.NaN(),
	}
	t.Parallel()
	r := marshal(unsupportedValues[0])
	require.Equal(t, r, "", "The code not return empty string")
}
