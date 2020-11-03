package link

type link struct {
	id, link string
}

var database []link

func copyDatabase(source []link) {
	database = make([]link, len(source))
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

func AddLink(id, url string) {
	if getLink(id) == "" {
	database = append(database, link{id, url})
	}
}

func getLink(id string) string {
	for _,v := range database {
		if v.id == id {
			return v.link
		}
	}
	return ""
}