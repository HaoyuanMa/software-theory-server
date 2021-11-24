package protocol

type Login struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}
