package slackim

type author struct {
	AuthorName string `json:"author_name,omitempty"` // eg. Bobby Tables
	AuthorLink string `json:"author_link,omitempty"` // eg. http://flickr.com/bobby/
	AuthorIcon string `json:"author_icon,omitempty"` // eg. http://flickr.com/icons/bobby.jpg
}

type title struct {
	Title     string `json:"title"`      // eg. Slack API Documentation
	TitleLink string `json:"title_link"` // eg. https://api.slack.com/
}

type image struct {
	ImageURL string `json:"image_url,omitempty"`
	ThumbURL string `json:"thumb_url,omitempty"`
}

type footer struct {
	Footer     string `json:"footer"`      // eg. Slack API
	FooterIcon string `json:"footer_icon"` // eg. https://platform.slack-edge.com/img/default_application_icon.png
}
