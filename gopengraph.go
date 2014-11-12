package gopengraph

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

type GopenGraph struct {
	OgAttrs        map[string]string
	Title          string
	Description    string
	MandatoryAttrs []string
}

// Check if all MandatoryAttrs are present
func (m *GopenGraph) IsValid() bool {
	for _, attr := range m.MandatoryAttrs {
		_, isPresent := m.OgAttrs[attr]

		if !isPresent {
			return false
		}
	}

	return true
}

// Store all Open Graph tags (http://ogp.me/)
func (m *GopenGraph) PopulateOgTags(doc *goquery.Document) {
	op := make(map[string]string)

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		prop, _ := s.Attr("property")
		match, _ := regexp.MatchString("^og:(.+)$", prop)

		if match {
			val, _ := s.Attr("content")
			op[prop] = val
		}

	})

	m.OgAttrs = op
}

// Store the page Title and Description
func (m *GopenGraph) PopulateAttrs(doc *goquery.Document) {
	m.Title = doc.Find("title").Text()
	m.Description, _ = doc.Find("meta[name='description']").Attr("content")
}

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
