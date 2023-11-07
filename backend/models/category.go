package models

type Category struct {
	ID       uint   `json:"id" gorm:"unique; not null"`
	Category string `json:"category_name"`
}
type SetNewName struct {
	Current string `json:"current"`
	New     string `json:"new"`
}
