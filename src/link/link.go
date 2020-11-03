package link

type link struct {
	id, link string
}

var database []link

func copyDatabase(source []link) {
	database = make([]link, len(source))
	copy(database, source)
}

// GetLink returns link for defined id..
func GetLink(id string) string {
	for _, v := range database {
		if v.id == id {
			return v.link
		}
	}
	return ""
}

// AddLink is adding a new short-link to the database.
func AddLink(id, url string) {
	if GetLink(id) == "" {
		database = append(database, link{id, url})
		return
	}
	// no test to force me add this call! setLink(id, url)
}
/*
func setLink(id, link string) {
	for _, v := range database {
		if id == v.id {
			v.link = link
		}
	}
}
*/