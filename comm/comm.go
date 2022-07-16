package comm

const (
	CodeRegister    = 1 + iota
	CodeRegisterRes // 服务器返回
	CodeLogin
	CodeLoginRes // 服务器返回
	CodeUserList
	CodeUserListRes
	CodeSms
)

const (
	MAXONLINEUSER = 1024
)

const (
	CodeSmsPrivateChat = 100 + iota
	CodeSmsPrivateChatRes
	CodeSmsGorup
	CodeSmsGroupRes
)

const (
	ServerSuccess = 200 + iota
	ServerFail
)

// 服务器返回消息
type ServerRes struct {
	Code uint8  `json:"code"` // 返回结果
	Msg  string `json:"msg"`
}

type RegistMsg struct {
	UserName   string `json:"username"`
	UserPasswd string `json:"userpasswd"`
}

type LoginMessage struct {
	UserId     string `json:"userid"`
	UserName   string `json:"username"`
	UserPasswd string `json:"userpasswd"`
}

type SmsPrivateChat struct {
	SendUserId   string `json:"userid"`
	SendUserName string `json:"sendusername"`
	RecvUserName string `json:"recvusername"` // 私聊名称
	ChatData     string `json:"data"`         // 私聊数据
}

type SmsMsg struct {
	Code    uint8  `json:"code"`
	SmsData string `json:"smsdata"`
}

// 通信消息结构体
type Msg struct {
	Code uint8  `json:"code"`
	Data string `json:"data"`
}
