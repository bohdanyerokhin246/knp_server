package config

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
