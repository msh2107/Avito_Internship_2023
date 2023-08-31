package models

type Segment struct {
	ID   int    `db:"id" json:"ID"`
	Slug string `db:"slug" json:"Slug"`
}
