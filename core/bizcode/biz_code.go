package bizcode

type BizCode int

const (
	Success      BizCode = 0
	ServerError  BizCode = 500
	ClientError  BizCode = 400
	DataNotfound BizCode = 4040

	TokenInvalid BizCode = 10001
	TokenExpire  BizCode = 10002

	UserNameExists  BizCode = 20001
	PhoneInvalid    BizCode = 20002
	UserNotExists   BizCode = 20003
	PasswordInvalid BizCode = 20004
)

var kv = map[BizCode]string{
	Success:         "success",
	ServerError:     "server error",
	ClientError:     "client error",
	DataNotfound:    "data not found error",
	TokenInvalid:    "token invalid",
	TokenExpire:     "token expire",
	UserNameExists:  "user name exists",
	PhoneInvalid:    "phone invalid",
	UserNotExists:   "user not exists",
	PasswordInvalid: "password invalid",
}

func (b BizCode) String() string {
	return kv[b]
}

func (b BizCode) Code() int {
	return int(b)
}
