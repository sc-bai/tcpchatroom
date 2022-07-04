package comm

const (
	CodeRegister = 1 + iota
	CodeLogin
	CodeSms
)

type LoginMessage struct {
	UserId     string `json:"userid"`
	UserName   string `json:"username"`
	UserPasswd string `json:"userpasswd"`
}

type SmsMessage struct {
}

type Msg struct {
	Code uint8  `json:"code"`
	Data string `json:"data"`
}
