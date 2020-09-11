package blockkit

import "encoding/json"

// TextObject https://api.slack.com/reference/block-kit/composition-objects#text
type TextObject struct {
	IsMrkdwn bool `json:"-"`
	// Body is text in json, renaming because duplicate with object name
	Body string `json:"text"`
}

// MarshalJSON fills default values and run normal marshal
func (to TextObject) MarshalJSON() ([]byte, error) {
	type Alias TextObject
	ot := "plain_text"
	if to.IsMrkdwn {
		ot = "mrkdwn"
	}
	return json.Marshal(struct {
		Alias
		Type string `json:"type"`
	}{
		Alias: (Alias)(to),
		Type:  ot,
	})
}

// AvailableAsContextBlockElement represents object can use in ContextBlock.Elements
func (to TextObject) AvailableAsContextBlockElement() bool {
	return true
}
