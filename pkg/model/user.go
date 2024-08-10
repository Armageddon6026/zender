package model

import "time"

type UserTable struct {
	Id       int    `json:"id,omitempty"`
	Account  string `json:"account,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
	Date     string `json:"date,omitempty"`
	Role     int    `json:"role,omitempty"`
	Auth     bool   `json:"Auth,omitempty"`
}

type UserInfo struct {
	Account string    `json:"account,omitempty"`
	Name    string    `json:"name,omitempty"`
	Date    time.Time `json:"date,omitempty"`
	Role    int       `json:"role,omitempty"`
}

type UserInsert struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserUpdate struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserAuthUpdate struct {
	Role int  `json:"role,omitempty"`
	Auth bool `json:"Auth,omitempty"`
}

type UserList []UserInfo
