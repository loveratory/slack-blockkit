package main

import (
	"encoding/json"
	"log"
	"os"

	blockkit "github.com/loveratory/slack-blockkit"
)

func main() {
	blocks := blockkit.BlockKit{
		blockkit.ContextBlock{
			Elements: []blockkit.ContextBlockElement{
				blockkit.ImageElement{
					ImageURL: "https://github.com/otofune.png",
					Alt:      "icon",
				},
				blockkit.TextObject{
					Body: "@otofune | 2020/05/08 07:54",
				},
			},
		},
		blockkit.SectionBlock{
			Text: blockkit.TextObject{
				Body: "Hello world~",
			},
		},
		blockkit.DividerBlock{},
	}
	b, err := json.MarshalIndent(blocks, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	os.Stdout.Write(b)
}
