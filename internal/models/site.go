package models

type Page struct {
	Path     string `json:"path,omitempty"`
	PageName string `json:"pageName,omitempty"`
	ForRole  string `json:"forRole,omitempty"`
}
type Menu struct {
	Name  string `json:"name,omitempty"`
	Path  string `json:"path,omitempty"`
	Group string `json:"group,omitempty"`
}

type MenuItem struct {
	ID       uint       `gorm:"primaryKey" json:"id"`
	Name     string     `json:"name"`
	Link     string     `json:"link"`
	Roles    string     `json:"roles"`
	Submenu  []MenuItem `gorm:"foreignKey:ParentID" json:"submenu,omitempty"`
	ParentID *uint      `gorm:"index" json:"-"`
}
