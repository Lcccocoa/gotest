package models

type Todo struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
