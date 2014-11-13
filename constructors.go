package gopengraph

import (
	"github.com/PuerkitoBio/goquery"
)

// Build a new instance of the GopenGraph struct based on an already scrapped site
func New(doc *goquery.Document) *GopenGraph {
	mg := new(GopenGraph)
	mg.MandatoryAttrs = []string{"og:title", "og:type", "og:image", "og:url"}
	mg.PopulateOgTags(doc)
	mg.PopulateAttrs(doc)

	return mg
}

// Build a new instance of the GopenGraph Struct with an url string
func NewFromUrl(url string) (*GopenGraph, error) {
	doc, err := goquery.NewDocument(url)

	return New(doc), err
}
