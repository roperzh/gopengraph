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

func (m *GopenGraph) IsValid() bool {
	for _, attr := range m.MandatoryAttrs {
		_, isPresent := m.OgAttrs[attr]

		if !isPresent {
			return false
		}
	}

	return true
}

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

func (m *GopenGraph) PopulateAttrs(doc *goquery.Document) {
	m.Title = doc.Find("title").Text()
	m.Description, _ = doc.Find("meta[name='description']").Attr("content")
}

func New(doc *goquery.Document) *GopenGraph {
	mg := new(GopenGraph)
	mg.MandatoryAttrs = []string{"og:title", "og:type", "og:image", "og:url"}
	mg.PopulateOgTags(doc)
	mg.PopulateAttrs(doc)

	return mg
}

func NewFromUrl(url string) (*GopenGraph, error) {
	doc, err := goquery.NewDocument(url)

	return New(doc), err
}
