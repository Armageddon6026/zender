package model

type GroupTable struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Note string `json:"note,omitempty"`
}

type GroupInfo struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Note string `json:"note,omitempty"`
}

type GroupList []GroupInfo

type GroupUpdate struct {
	Name string `json:"name,omitempty"`
	Note string `json:"note,omitempty"`
}
