package github.com/jroden2/dumb

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToReverseCamelCase(t *testing.T) {
	t.Run("Should return successful with rEvErSeCaMeLcAsE", func(t *testing.T) {
		input := "reversecamelcase"
		want := "rEvErSeCaMeLcAsE"

		got := ToReverseCamelCase(input)
		fmt.Println(got)
		assert.Equal(t, want, got)
	})
}

func TestToCamelCase(t *testing.T) {
	t.Run("Should return successful with CaMeLcAsE", func(t *testing.T) {
		input := "camelcase"
		want := "CaMeLcAsE"

		got := ToCamelCase(input)
		fmt.Println(got)
		assert.Equal(t, want, got)
	})
}

