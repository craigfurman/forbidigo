package forbidigo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseValidPattern(t *testing.T) {
	ptrn, err := parse(`fmt\.Errorf`)
	require.Nil(t, err)
	assert.Equal(t, `fmt\.Errorf`, ptrn.pattern.String())
}

func TestParseValidPatternThatUsesSquareBrackets(t *testing.T) {
	ptrn, err := parse(`[f]mt\.Errorf`)
	require.Nil(t, err)
	assert.Equal(t, `[f]mt\.Errorf`, ptrn.pattern.String())
}

func TestParseValidPatternWithCustomMessage(t *testing.T) {
	ptrn, err := parse(`#[Please don't use this!]fmt\.Println`)
	require.Nil(t, err)
	assert.Equal(t, `fmt\.Println`, ptrn.pattern.String())
	assert.Equal(t, "Please don't use this!", ptrn.msg)
}

func TestParseInvalidPattern_ReturnsError(t *testing.T) {
	_, err := parse(`fmt\`)
	assert.NotNil(t, err)
}
