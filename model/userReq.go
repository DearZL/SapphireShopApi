package model

type UserListReq struct {
	Page     uint32
	PageSize uint32
}

type UserRegReq struct {
	Email    string `json:"email,omitempty" bind:"require"`
	UserName string `json:"userName,omitempty" bind:"require"`
	Password string `json:"password,omitempty" bind:"require"`
	Code     string `json:"code,omitempty" bind:"require"`
}
