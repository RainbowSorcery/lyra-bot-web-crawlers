package e

const (
	SUCCESS = 200
	ERROR   = 500
)

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "failed",
}

func getMsg(code int) string {
	return MsgFlags[code]
}
