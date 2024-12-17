package domain

type Module struct {
	Id       uint    `json:"id"`
	Name     string  `json:"name"`
	Desc     *string `json:"desc"`
	Status   int     `json:"status"`
	ParentId uint    `json:"parentId"`
}
