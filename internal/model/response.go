package model

type Meta struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type LoginResponse struct {
	Meta  Meta   `json:"meta"`
	Token string `json:"token"`
}

type DataResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}
