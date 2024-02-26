package usercontroller

type InputReegister struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type InputLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
