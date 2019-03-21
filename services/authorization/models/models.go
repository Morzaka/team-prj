package models

type User struct {
	Signin Signin `json:"signin"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Role     string `json:"role"`
}

type Signin struct{
	Login string `json:"login"`
	Password string `json:"password"`
}

