package models

type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}
