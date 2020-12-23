package resolvers

type Publication struct {
	Title string `json:"title"`
	URI   string `json:"uri"`
	Date  string `json:"date"`
}

func GetPublications() []Publication {
	return []Publication{
		{
			Title: "test title 1",
			URI:   "www.testuri1.com",
			Date:  "testdate1",
		},
		{
			Title: "test title 2",
			URI:   "www.testuri2.com",
			Date:  "testdate2",
		},
	}
}
