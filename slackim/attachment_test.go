package slackim

import (
	"testing"

	"github.com/myteksi/go/automation/gandalf/engine/glamdring/trand"
	"github.com/stretchr/testify/require"
)

const (
	randStringLength = 20
)

func TestNewAttachment(t *testing.T) {
	x := NewAttachment("Hello")
	require.IsType(t, &Attachment{}, x)
	require.Equal(t, "Hello", x.Fallback)
	require.NotNil(t, x.Timestamp)
}

func TestAttachment_Setters(t *testing.T) {
	a := &Attachment{}

	str := trand.AlphaNumericString(randStringLength)
	require.Equal(t, "", a.Color)
	a.SetColor(str)
	require.Equal(t, str, a.Color)

	str = trand.AlphaNumericString(randStringLength)
	require.Equal(t, "", a.Pretext)
	a.SetPretext(str, false)
	require.Equal(t, str, a.Pretext)

	str1 := trand.AlphaNumericString(randStringLength)
	str2 := trand.AlphaNumericString(randStringLength)
	str3 := trand.AlphaNumericString(randStringLength)
	require.Equal(t, "", a.AuthorName)
	require.Equal(t, "", a.AuthorLink)
	require.Equal(t, "", a.AuthorIcon)
	a.SetAuthor(str1, str2, str3, false)
	require.Equal(t, str1, a.AuthorName)
	require.Equal(t, str2, a.AuthorLink)
	require.Equal(t, str3, a.AuthorIcon)

	str1 = trand.AlphaNumericString(randStringLength)
	str2 = trand.AlphaNumericString(randStringLength)
	require.Equal(t, "", a.Title)
	require.Equal(t, "", a.TitleLink)
	a.SetTitle(str1, str2, false)
	require.Equal(t, str1, a.Title)
	require.Equal(t, str2, a.TitleLink)

	str = trand.AlphaNumericString(randStringLength)
	require.Equal(t, "", a.ImageURL)
	a.SetImage(str)
	require.Equal(t, str, a.ImageURL)

	str = trand.AlphaNumericString(randStringLength)
	require.Equal(t, "", a.ThumbURL)
	a.SetThumbnail(str)
	require.Equal(t, str, a.ThumbURL)

	str1 = trand.AlphaNumericString(randStringLength)
	str2 = trand.AlphaNumericString(randStringLength)
	require.Equal(t, "", a.Footer)
	require.Equal(t, "", a.FooterIcon)
	a.SetFooter(str1, str2)
	require.Equal(t, str1, a.Footer)
	require.Equal(t, str2, a.FooterIcon)
}

func TestAttachment_checkMarkdown(t *testing.T) {
	a := &Attachment{}
	a.SetPretext("Hello", false)
	require.Empty(t, a.MarkdownIn)

	a.SetPretext("Hello", true)
	require.Contains(t, a.MarkdownIn, "pretext")

	a.SetPretext("Hello", false)
	require.Empty(t, a.MarkdownIn)
}
