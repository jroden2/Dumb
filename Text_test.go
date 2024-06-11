package github.com/jroden2/dumb

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToCamelCase(t *testing.T) {
  t.Run("Should return successful with cAmElCaSe", func(t *testing.T)){
    input := "camelcase"
    want := "cAmElCaSe"
    
    got := ToCamelCase(input)
    fmt.Println(got)
		assert.Equal(t, want, got)
  }
}
