package models

type Publisher struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"size:255;not null;unique"`
	Books []Book // has many relationship
}
