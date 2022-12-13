package e

const (
	SUCESS = 200
	ERROR  = 500
)

var MsgFlags = map[int]string{
	SUCESS: "ok",
	ERROR:  "failed",
}

func getMsg(code int) string {
	return MsgFlags[code]
}
