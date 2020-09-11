package blockkit

import "encoding/json"

// ImageElement https://api.slack.com/reference/block-kit/block-elements#image
type ImageElement struct {
	ImageURL string `json:"image_url"`
	Alt      string `json:"alt_text"`
}

// MarshalJSON fills default values and call original marshal
func (ie ImageElement) MarshalJSON() ([]byte, error) {
	type Alias ImageElement
	return json.Marshal(struct {
		Alias
		Type string `json:"type"`
	}{
		Alias: (Alias)(ie),
		Type:  "image",
	})
}

// AvailableAsContextBlockElement represents object can use in ContextBlock.Elements
func (ie ImageElement) AvailableAsContextBlockElement() bool {
	return true
}
