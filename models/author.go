package models

type Author struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"size:255;not null"`
	Bio   string
	Books []Book // has many relationship
}
