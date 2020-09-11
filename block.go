package blockkit

import (
	"encoding/json"
	"fmt"
)

// Block block type
type Block interface {
	// FIXME: もっと意味のあるインターフェースにしたい
	IsBlock() bool
}

// SectionBlock https://api.slack.com/reference/block-kit/blocks#section
type SectionBlock struct {
	Text TextObject `json:"text"`
}

// IsBlock is block?
func (sb SectionBlock) IsBlock() bool {
	return true
}

// MarshalJSON fills default values and run normal marshal
func (sb SectionBlock) MarshalJSON() ([]byte, error) {
	type Alias SectionBlock
	return json.Marshal(struct {
		Alias
		Type string `json:"type"`
	}{
		Alias: (Alias)(sb),
		Type:  "section",
	})
}

// ContextBlock https://api.slack.com/reference/block-kit/blocks#context
type ContextBlock struct {
	Elements []ContextBlockElement `json:"elements"`
}

// ContextBlockElement represents children of context block
type ContextBlockElement interface {
	AvailableAsContextBlockElement() bool
}

// IsBlock is block?
func (cb ContextBlock) IsBlock() bool {
	return true
}

// MarshalJSON validate and fills default values and run normal marshal
func (cb ContextBlock) MarshalJSON() ([]byte, error) {
	if len(cb.Elements) > 10 {
		return nil, fmt.Errorf("context can include 10 or less elements, but %d elements given", len(cb.Elements))
	}
	type Alias ContextBlock
	return json.Marshal(struct {
		Alias
		Type string `json:"type"`
	}{
		Alias: (Alias)(cb),
		Type:  "context",
	})
}

// DividerBlock https://api.slack.com/reference/block-kit/blocks#divider
type DividerBlock struct{}

// IsBlock is block?
func (db DividerBlock) IsBlock() bool {
	return true
}

// MarshalJSON fills default values and run normal marshal
func (db DividerBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type string `json:"type"`
	}{
		Type: "divider",
	})
}
