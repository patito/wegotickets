package scrape

import "github.com/patito/wegotickets/html"

const (
	// dateAttr - Date Information
	dateAttr = ".event-information " + html.TagH4

	// venueAttr - Venue Informetion
	venueAttr = ".event-information " + html.TagH2

	// artistAttr - Artist information
	artistAttr = ".event-information " + html.TagH1

	// linksAttr - to walk trought all event links
	linksAttr = ".chatterbox-margin"

	// eventLinkAttr - Get the specific link of an event
	eventLinkAttr = html.TagH2 + html.TagA
)
