package slackim

// Field is the field type contained in an attachment
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// AddField adds a new field to the struct
func (a *Attachment) AddField(title, value string, short bool, containsMarkdown bool) *Attachment {
	a.Fields = append(a.Fields, &Field{
		Title: title,
		Value: value,
		Short: short,
	})
	a.checkMarkdown(containsMarkdown, "fields")
	return a
}

// SetAllFieldsShort sets all Fields' "short" value to true,
// making the Fields display horizontally
func (a *Attachment) SetAllFieldsShort() *Attachment {
	for _, f := range a.Fields {
		f.Short = true
	}
	return a
}
