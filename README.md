blockkit
===

`blockkit` provides structs can convert to Slakc Block Kit JSON with json.Marshal.

This project only implemented subset of Block Kit currently.

Install
---

```sh
go get github.com/loveratory/slack-blockkit
```

Example
---

See [./cmd/example](./cmd/example).

```
$ go run ./cmd/example
[
  {
    "elements": [
      {
        "image_url": "https://github.com/otofune.png",
        "alt_text": "icon",
        "type": "image"
      },
      {
        "text": "@otofune | 2020/05/08 07:54",
        "type": "plain_text"
      }
    ],
    "type": "context"
  },
  {
    "text": {
      "text": "Hello world~",
      "type": "plain_text"
    },
    "type": "section"
  },
  {
    "type": "divider"
  }
]
```
