// Package blockkit provides structs can convert to Slakc Block Kit JSON with json.Marshal
package blockkit

import (
	"encoding/json"
	"fmt"
)

// BlockKit represents block kit root
type BlockKit []Block

// MarshalJSON validate and run normal marshal
func (bk BlockKit) MarshalJSON() ([]byte, error) {
	type Alias BlockKit
	if len(bk) > 50 {
		return nil, fmt.Errorf("block kit can include 50 or less blocks, but %d blocks given", len(bk))
	}

	return json.Marshal((Alias)(bk))
}
