package models

type Key struct {
	ID  int64  `db:"ID" json:"id"`
	URL string `db:"URL" json:"url"`
}
