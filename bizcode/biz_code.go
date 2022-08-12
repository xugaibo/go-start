package bizcode

type BizCode int

const (
	Success     BizCode = 0
	ServerError BizCode = 500
	ClientError BizCode = 400
)

var kv = map[BizCode]string{
	Success:     "success",
	ServerError: "server error",
	ClientError: "client error",
}

func (b BizCode) String() string {
	return kv[b]
}

func (b BizCode) Code() int {
	return int(b)
}
