package message

const (
	LoginMesType       = "LoginMes"
	LoginResMesType    = "LoginResMes"
	RegisterMesType    = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type RegisterMes struct {
	User User `json:"user"` //User结构体类型，可调用其中字段
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
