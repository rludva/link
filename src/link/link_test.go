package link

import (
	"testing"
)

func TestGetLink(t *testing.T) {
	type testCase struct {
		links    []link
		id, link string
	}

	testCases := []testCase{
		{
			[]link{},
			"",
			"",
		},
		{
			[]link{
				{"123", "http://www.nutius.com"},
			},
			"123",
			"http://www.nutius.com",
		},
		{
			[]link{
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

func TestAddLink(t *testing.T) {
	type testCase struct {
		id, link                       string
		initialDatabase, finalDatabase []link
	}

	testCases := []testCase{
		{
			"a",
			"http://www.abc.com",
			[]link{},
			[]link{
				{"a", "http://www.avc.com"},
			},
		},
		{
			"existing item",
			"http://www.existing.com",
			[]link{
				{"existing item", "http://www.existing.com"},
			},
			[]link{
				{"existing item", "http://www.existing.com"},
			},
		},
		{
			"existing item",
			"http://www.existing-with-differen.content.com",
			[]link{
				{"existing item", "http://www.existing.com"},
			},
			[]link{
				{"existing item", "http://www.existing-with-differen.content.com"},
			},
		},
	}

	for _, tt := range testCases {
		copyDatabase(tt.initialDatabase)
		AddLink(tt.id, tt.link)
		if len(tt.finalDatabase) != len(database) {
			t.Errorf("initialDatabase: len(): %v, %v\n", len(tt.initialDatabase), tt.initialDatabase)
			t.Errorf("finalDatabase: len(): %v, %v\n", len(tt.finalDatabase), tt.finalDatabase)
			t.Errorf("database: len(): %v, %v\n", len(database), database)
		}
	}
}
