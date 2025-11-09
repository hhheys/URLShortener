package model

type Link struct {
	Id        int64  `json:"id"`
	Link      string `json:"link"`
	ShortLink string `json:"short_link"`
}

func NewLink(id int64, linkString string, shortLink string) *Link {
	return &Link{
		Id:        id,
		Link:      linkString,
		ShortLink: shortLink,
	}
}
