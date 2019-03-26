package models
//User struct contains user data
type User struct {
	Signin Signin `json:"signin"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Role     string `json:"role"`
}
//Signin contains data for logging in
type Signin struct{
	Login string `json:"login"`
	Password string `json:"password"`
}

