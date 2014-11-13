package gopengraph

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

var doc *goquery.Document

func Doc() *goquery.Document {
	if doc == nil {
		doc = loadDoc("page")
	}
	return doc
}

func loadDoc(page string) *goquery.Document {
	var f *os.File
	var e error

	if f, e = os.Open(fmt.Sprintf("./testdata/%s.html", page)); e != nil {
		panic(e.Error())
	}
	defer f.Close()

	var node *html.Node
	if node, e = html.Parse(f); e != nil {
		panic(e.Error())
	}

	doc := goquery.NewDocumentFromNode(node)
	doc.Url, _ = url.Parse("http://example.com")

	return doc
}

func TestPageTitle(t *testing.T) {
	mg := New(Doc())

	if mg.Title != "Page Title" {
		t.Error("Expected <title> tag to have the content 'Page Title'.")
	}
}

func TestPageDescription(t *testing.T) {
	mg := New(Doc())

	if mg.Description != "Page description" {
		t.Error("Expected <meta> description tag to have the content 'Page description'.")
	}
}

func TestPageOgTags(t *testing.T) {
	mg := New(Doc())

	if mg.OgAttrs["og:type"] != "pageType" {
		t.Error("Expected 'og:type' tag to have the content pageType.")
	}

	if mg.OgAttrs["og:url"] != "pageUrl" {
		t.Error("Expected 'og:url' tag to have the content pageUrl.")
	}

	if mg.OgAttrs["og:title"] != "pageTitle" {
		t.Error("Expected 'og:title' tag to have the content pageTitle.")
	}

	if mg.OgAttrs["og:description"] != "pageDescription" {
		t.Error("Expected 'og:description' tag to have the content pageDescription.")
	}

	if mg.OgAttrs["og:image"] != "/assets/image.jpg" {
		t.Error("Expected 'og:image' tag to have the content /assets/image.jpg.")
	}
}
