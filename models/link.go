package models

import (
	"main/db"
)

// Link ...
type Link struct {
	NumID int    `db:"num_id" json:"num_id"`
	URL   string `db:"url" json:"url"`
}

// LinkModel ...
type LinkModel struct{}

// Create ...
func (m LinkModel) Create(url string) (link Link, err error) {
	err = db.GetDB().QueryRow("INSERT INTO urls VALUES(DEFAULT, $1) RETURNING num_id", url).Scan(&link.NumID)

	return link, err
}

// Find ...
func (m LinkModel) Find(numID int) (link Link, err error) {
	err = db.GetDB().SelectOne(&link, "SELECT num_id, url FROM urls WHERE num_id=$1", numID)

	return link, err
}
