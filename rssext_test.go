package feeds

import (
	"encoding/xml"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

var testRssFeedExtXML = RssFeedExtXml{
	XMLName:          xml.Name{Space: "", Local: "rss"},
	Version:          "2.0",
	ContentNamespace: "",
	Channel: &RssFeedExt{
		XMLName:        xml.Name{Space: "", Local: "channel"},
		Title:          "Lorem ipsum feed for an interval of 1 minutes",
		Link:           "http://example.com/",
		Description:    "This is a constantly updating lorem ipsum feed",
		Language:       "",
		Copyright:      "Michael Bertolacci, licensed under a Creative Commons Attribution 3.0 Unported License.",
		ManagingEditor: "",
		WebMaster:      "",
		PubDate:        "Tue, 30 Oct 2018 23:22:00 GMT",
		LastBuildDate:  "Tue, 30 Oct 2018 23:22:37 GMT",
		Category:       "",
		Generator:      "RSS for Node",
		Docs:           "",
		Cloud:          "",
		Ttl:            60,
		Rating:         "",
		SkipHours:      "",
		SkipDays:       "",
		Image:          (*RssImageExt)(nil),
		TextInput:      (*RssTextInputExt)(nil),
		Items: []*RssItemExt{
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:22:00+00:00",
				Link:        "http://example.com/test/1540941720",
				Description: "Exercitation ut Lorem sint proident.",
				Content:     (*RssContentExt)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosureExt)(nil),
				Guid:        &RssGuidExt{XMLName: xml.Name{Local: "guid"}, Id: "http://example.com/test/1540941720", IsPermaLink: "true"},
				PubDate:     "Tue, 30 Oct 2018 23:22:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:21:00+00:00",
				Link:        "http://example.com/test/1540941660",
				Description: "Ea est do quis fugiat exercitation.",
				Content:     (*RssContentExt)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosureExt)(nil),
				Guid:        &RssGuidExt{XMLName: xml.Name{Local: "guid"}, Id: "http://example.com/test/1540941660", IsPermaLink: "true"},
				PubDate:     "Tue, 30 Oct 2018 23:21:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:20:00+00:00",
				Link:        "http://example.com/test/1540941600",
				Description: "Ipsum velit cillum ad laborum sit nulla exercitation consequat sint veniam culpa veniam voluptate incididunt.",
				Content:     (*RssContentExt)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosureExt)(nil),
				Guid:        &RssGuidExt{XMLName: xml.Name{Local: "guid"}, Id: "http://example.com/test/1540941600", IsPermaLink: "true"},
				PubDate:     "Tue, 30 Oct 2018 23:20:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:19:00+00:00",
				Link:        "http://example.com/test/1540941540",
				Description: "Ullamco pariatur aliqua consequat ea veniam id qui incididunt laborum.",
				Content:     (*RssContentExt)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosureExt)(nil),
				Guid:        &RssGuidExt{XMLName: xml.Name{Local: "guid"}, Id: "http://example.com/test/1540941540", IsPermaLink: "true"},
				PubDate:     "Tue, 30 Oct 2018 23:19:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:18:00+00:00",
				Link:        "http://example.com/test/1540941480",
				Description: "Velit proident aliquip aliquip anim mollit voluptate laboris voluptate et occaecat occaecat laboris ea nulla.",
				Content:     (*RssContentExt)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosureExt)(nil),
				Guid:        &RssGuidExt{XMLName: xml.Name{Local: "guid"}, Id: "http://example.com/test/1540941480", IsPermaLink: "true"},
				PubDate:     "Tue, 30 Oct 2018 23:18:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:17:00+00:00",
				Link:        "http://example.com/test/1540941420",
				Description: "Do in quis mollit consequat id in minim laborum sint exercitation laborum elit officia.",
				Content:     (*RssContentExt)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosureExt)(nil),
				Guid:        &RssGuidExt{XMLName: xml.Name{Local: "guid"}, Id: "http://example.com/test/1540941420", IsPermaLink: "true"},
				PubDate:     "Tue, 30 Oct 2018 23:17:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:16:00+00:00",
				Link:        "http://example.com/test/1540941360",
				Description: "Irure id sint ullamco Lorem magna consectetur officia adipisicing duis incididunt.",
				Content:     (*RssContentExt)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosureExt)(nil),
				Guid:        &RssGuidExt{XMLName: xml.Name{Local: "guid"}, Id: "http://example.com/test/1540941360", IsPermaLink: "true"},
				PubDate:     "Tue, 30 Oct 2018 23:16:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:15:00+00:00",
				Link:        "http://example.com/test/1540941300",
				Description: "Sunt anim excepteur esse nisi commodo culpa laborum exercitation ad anim ex elit.",
				Content:     (*RssContentExt)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosureExt)(nil),
				Guid:        &RssGuidExt{XMLName: xml.Name{Local: "guid"}, Id: "http://example.com/test/1540941300", IsPermaLink: "true"},
				PubDate:     "Tue, 30 Oct 2018 23:15:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:14:00+00:00",
				Link:        "http://example.com/test/1540941240",
				Description: "Excepteur aliquip fugiat ex labore nisi.",
				Content:     (*RssContentExt)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosureExt)(nil),
				Guid:        &RssGuidExt{XMLName: xml.Name{Local: "guid"}, Id: "http://example.com/test/1540941240", IsPermaLink: "true"},
				PubDate:     "Tue, 30 Oct 2018 23:14:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:13:00+00:00",
				Link:        "http://example.com/test/1540941180",
				Description: "Id proident adipisicing proident pariatur aute pariatur pariatur dolor dolor in voluptate dolor.",
				Content:     (*RssContentExt)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosureExt)(nil),
				Guid:        &RssGuidExt{XMLName: xml.Name{Local: "guid"}, Id: "http://example.com/test/1540941180", IsPermaLink: "true"},
				PubDate:     "Tue, 30 Oct 2018 23:13:00 GMT",
				Source:      "",
			},
		},
	},
}

func TestRssExtUnmarshal(t *testing.T) {
	var xmlFeed RssFeedExtXml
	xmlFile, err := os.Open("testext.rss")
	if err != nil {
		panic("AHH file bad")
	}
	bytes, _ := io.ReadAll(xmlFile)
	if err := xml.Unmarshal(bytes, &xmlFeed); err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(testRssFeedExtXML, xmlFeed) {
		diffs := pretty.Diff(testRssFeedExtXML, xmlFeed)
		t.Log(pretty.Println(diffs))
		t.Error("object was not unmarshalled correctly")
	}

}
