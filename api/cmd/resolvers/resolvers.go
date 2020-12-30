package resolvers

type Publication struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URI   string `json:"uri"`
	Date  string `json:"date"`
}

var test1 = Publication{
	ID:    01,
	Title: "test title 1",
	URI:   "www.testuri1.com",
	Date:  "testdate1",
}

var test2 = Publication{
	ID:    02,
	Title: "test title 2",
	URI:   "www.testuri2.com",
	Date:  "testdate2",
}

var Publications = []Publication{test1, test2}

func GetPublications() []Publication {
	return Publications
}

func GetPublication(id int) Publication {
	for _, pub := range Publications {
		if pub.ID == id {
			return pub
		}
	}

	return Publication{}
}

func CreatePublication(id int, title, uri, dateAdded string) Publication {
	toReturn := Publication{
		ID:    id,
		Title: title,
		URI:   uri,
		Date:  dateAdded,
	}

	Publications = append(Publications, toReturn)

	return toReturn
}
