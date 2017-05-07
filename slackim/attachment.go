package slackim

import (
	"time"
)

// Attachment is the attachment type for AttachmentMessage
type Attachment struct {
	BasicMessage
	author
	title
	image
	footer

	MarkdownIn []string `json:"mrkdwn_in,omitemptys"`
	Fallback   string   `json:"fallback"`          // Required plain-text summary of the attachment
	Color      string   `json:"color,omitempty"`   // Sidebar colour
	Pretext    string   `json:"pretext,omitempty"` // Optional text that appears above the attachment block
	Fields     []*Field `json:"fields,omitempty"`
	Timestamp  int64    `json:"ts"`
}

// NewAttachment begins a new attachment
func NewAttachment(fallbackText string) *Attachment {
	return &Attachment{
		Fallback:  fallbackText,
		Timestamp: time.Now().Unix(),
	}
}

// SetColor sets the Color of the border along the left side of the message attachment.
// The input can either be one of "good", "warning", "danger", or any hex color code (eg. #439FE0).
func (a *Attachment) SetColor(color string) *Attachment {
	a.Color = color
	return a
}

// SetPretext sets the optional text that appears above the message attachment block.
func (a *Attachment) SetPretext(text string, containsMarkdown bool) *Attachment {
	a.Pretext = text
	a.checkMarkdown(containsMarkdown, "pretext")
	return a
}

// SetAuthor displays a small section at the top of a message attachment with author information.
//   name : Small text used to display the author's name.
//   link : A valid URL that will hyperlink the author_name text mentioned above. Will only work if author_name is present.
//   icon : A valid URL that displays a small 16x16px image to the left of the author_name text. Will only work if author_name is present.
func (a *Attachment) SetAuthor(name, link, icon string, containsMarkdown bool) *Attachment {
	a.AuthorName = name
	a.AuthorLink = link
	a.AuthorIcon = icon
	a.checkMarkdown(containsMarkdown, "author_name")
	return a
}

// SetTitle sets the larger, bold text near the top of a message attachment.
// By passing a valid URL in the link parameter, the title text will be hyperlinked.
func (a *Attachment) SetTitle(title, link string, containsMarkdown bool) *Attachment {
	a.Title = title
	a.TitleLink = link
	a.checkMarkdown(containsMarkdown, "title")
	return a
}

// SetImage sets an image to be displayed inside a message attachment.
func (a *Attachment) SetImage(imageURL string) *Attachment {
	a.ImageURL = imageURL
	return a
}

// SetThumbnail sets an thumbnail image to be displayed on the right side, within a message attachment.
func (a *Attachment) SetThumbnail(thumbnailURL string) *Attachment {
	a.ThumbURL = thumbnailURL
	return a
}

// SetFooter sets the smaller, light text near the bottom of a message attachment.
// By passing a valid icon URL in the iconURL parameter, an icon will be displayed beside the text.
func (a *Attachment) SetFooter(text, iconURL string) *Attachment {
	a.Footer = text
	a.FooterIcon = iconURL
	return a
}

//============================================================ UNEXPORTED

// checkMarkdown checks that the given field is included/excluded correctly
func (a *Attachment) checkMarkdown(containsMarkdown bool, fieldName string) {
	if containsMarkdown {
		pos := findInList(a.MarkdownIn, fieldName)
		if pos < 0 {
			a.MarkdownIn = append(a.MarkdownIn, fieldName) // add item to a.MarkdownIn
		}

	} else {
		pos := findInList(a.MarkdownIn, fieldName)
		if len(a.MarkdownIn) == 1 && pos == 0 {
			a.MarkdownIn = []string{}
		} else if pos != -1 {
			a.MarkdownIn = append(a.MarkdownIn[:pos], a.MarkdownIn[pos:]...) // remove item in pos
		}
	}
}
