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
				{"a", "http://www.abc.com"},
			},
		},
		{
			"existing",
			"http://www.existing.com",
			[]link{
				{"existing", "http://www.existing.com"},
			},
			[]link{
				{"existing", "http://www.existing.com"},
			},
		},
		{
			"donald",
			"http://www.donald-trump.com",
			[]link{
				{"donald", "http://www.trump.com"},
			},
			[]link{
				{"donald", "http://www.donald-trump.com"},
			},
		},
	}

	for _, tt := range testCases {
		copyDatabase(tt.finalDatabase)
		expectedValue := GetLink(tt.id)

		copyDatabase(tt.initialDatabase)
		AddLink(tt.id, tt.link)
		value := GetLink(tt.id)
		if value != expectedValue {
			t.Errorf("initialDatabase: len(): %v, %v\n", len(tt.initialDatabase), tt.initialDatabase)
			t.Errorf("expected finalDatabase: len(): %v, %v\n", len(tt.finalDatabase), tt.finalDatabase)
			t.Errorf("cuurentDatabase: len(): %v, %v\n", len(database), database)
			t.Errorf("---")
		}
	}
}

func TestRemoveLink(t *testing.T) {
	type testCase struct {
		id                             string
		initialDatabase, finalDatabase []link
	}

	testCases := []testCase{
		{
			"a",
			[]link{
				{"a", "http://www.abc.com"},
			},
			[]link{},
		},
	}

	for _, tt := range testCases {
		copyDatabase(tt.initialDatabase)
		DeleteLink(tt.id)
		value := GetLink(tt.id)

		if value != "" {
			t.Errorf("initialDatabase: len(): %v, %v\n", len(tt.initialDatabase), tt.initialDatabase)
			t.Errorf("expected finalDatabase: len(): %v, %v\n", len(tt.finalDatabase), tt.finalDatabase)
			t.Errorf("cuurentDatabase: len(): %v, %v\n", len(database), database)
			t.Errorf("---")
		}
	}
}
