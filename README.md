# Gopengraph

GopenGraph is a very simple Go library for parsing [Open Graph](http://ogp.me/) protocol and Metadata information from web sites.

## Installation

```bash
$ go get github.com/roperzh/gopengraph
```

## Usage

```go

import (
  "fmt"

  "github.com/roperzh/gopengraph"
)

// Directly from an URL
pageFromUrl := gopengraph.NewFromUrl("https://github.com/")

fmt.Printf("Title: %s, Description: %s", pageFromUrl.Title, pageFromUrl.Description)
// => Title: Build software better, together, Description: GitHub is ..

// From a *goquery.Document
doc, err := goquery.NewDocument(url)
pageFromDocument := gopengraph.New(doc)

fmt.Printf("og:site_name: %s, og:url: %s", pageFromDocument.OgAttrs["og:site_name"], pageFromDocument.OgAttrs["og:url"])
// => og:site_name: GitHub, og:url: https://github.com

```
## Contributing


1- Fork it

2- Create your feature branch (git checkout -b my-new-feature)

3- Commit your changes (git commit -am 'Add some feature')

4- Push to the branch (git push origin my-new-feature)

5- Create new Pull Request

## License

MIT License
