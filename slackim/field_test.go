package slackim

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAttachment_AddField(t *testing.T) {
	a := &Attachment{}
	require.Empty(t, a.Fields)

	a.AddField("HELLO", "world", false, false)
	require.Equal(t, "HELLO", a.Fields[0].Title)
	require.Equal(t, "world", a.Fields[0].Value)
}

func TestAttachment_SetFieldsHorizontal(t *testing.T) {
	a := &Attachment{}
	a.AddField("Test", "horizontal", false, false)
	require.False(t, a.Fields[0].Short)
	a.SetAllFieldsShort()
	require.True(t, a.Fields[0].Short)
}
