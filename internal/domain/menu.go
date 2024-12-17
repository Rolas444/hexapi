package domain

type Menu struct {
	Id       uint `json:"id"`
	IdModule uint
	Name     string `json:"name"`
	Path     string
	Desc     *string `json:"desc"`
	Status   int     `json:"status"`
	ParentId uint    `json:"parentId"`
}
