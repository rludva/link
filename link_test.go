package link

import (
	"testing"
)

type links []struct {
	id, link string
}

var database links

func copyDatabase(source links) {
	database = make(links, 5)
	copy(database, source)
}

func GetLink(id string) string {
	for _, v := range database {
		if v.id == id {
			return v.link
		}
	}
	return ""
}

func TestGetLink(t *testing.T) {
	type testCase struct {
		links    links
		id, link string
	}

	testCases := []testCase{
		{
			links{},
			"",
			"",
		},
		{
			links{
				{"123", "http://www.nutius.com"},
			},
			"123",
			"http://www.nutius.com",
		},
		{
			links{
				{"123", "http://www.nutius.com"},
				{"1234", "http://www.radmi.cz"},
			},
			"1234",
			"http://www.radmi.cz",
		},
	}

	for _, tt := range testCases {
		copyDatabase(tt.links)
		if link := GetLink(tt.id); link != tt.link {
			t.Errorf("GetLink(`%v`):`%v`, WANT: `%v`,\nDatabase: `%v`", tt.id, link, tt.link, database)
		}
	}
}
