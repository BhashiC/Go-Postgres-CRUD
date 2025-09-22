package models

type Book struct {
	ID          uint      `gorm:"primary key;autoIncrement" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	AuthorID    uint      `json:"author_id"`
	Author      Author    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"author"`
	PublisherID uint      `json:"publisher_id"`
	Publisher   Publisher `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"publisher"`
}
