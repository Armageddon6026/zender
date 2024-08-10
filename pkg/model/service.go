package model

type ServiceTable struct {
	Sn         int    `json:"sn,omitempty"`
	Uuid       string `json:"uuid,omitempty"`
	Name       string `json:"name,omitempty"`
	PrivateKey string `json:"privateKey,omitempty"`
	GroupId    int    `json:"groupId,omitempty"`
}

type ServiceInfo struct {
	Uuid       string `json:"uuid,omitempty"`
	Name       string `json:"name,omitempty"`
	PrivateKey string `json:"privateKey,omitempty"`
	GroupId    int    `json:"groupId,omitempty"`
}

type ServiceInsert struct {
	Uuid       string `json:"uuid,omitempty"`
	Name       string `json:"name,omitempty"`
	PrivateKey string `json:"privateKey,omitempty"`
	GroupId    int    `json:"groupId,omitempty"`
}

type ServiceUpdate struct {
	Name    string `json:"name,omitempty"`
	GroupId int    `json:"groupId,omitempty"`
}

type ServciceList []ServiceInfo
