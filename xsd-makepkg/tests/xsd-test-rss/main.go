package main

import (
	"encoding/xml"

	"github.com/metaleap/go-xsd/xsd-makepkg/tests"

	util "github.com/metaleap/go-util"

	rss "github.com/metaleap/go-xsd-pkg/thearchitect.co.uk/schemas/rss-2_0.xsd_go"
)

type RssDoc struct {
	XMLName xml.Name `xml:"rss"`
	rss.TxsdRss
}

func main() {
	var (
		dirBasePath  = util.GopathSrcGithub("metaleap", "go-xsd", "xsd-makepkg", "tests", "xsd-test-rss")
		makeEmptyDoc = func() interface{} { return &RssDoc{} }
	)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
}