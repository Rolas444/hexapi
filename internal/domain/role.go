package domain

type Role struct {
	Id     uint    `json:"id"`
	Name   string  `json:"name"`
	Desc   *string `json:"desc"`
	Status int     `json:"status"`
}
