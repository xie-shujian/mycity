package e

type City struct {
	ID          int64  `db:"ID" json:"id"`
	Name        string `db:"Name" json:"name"`
	CountryCode string `db:"CountryCode" json:"country_code"`
	District    string `db:"District" json:"district"`
	Population  int64  `db:"Population" json:"population"`
}

type PageCity struct {
	Cities []City `json:"cities"`
	Page   Page   `json:"page"`
}
