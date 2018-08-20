package e

var MsgFlags = map[int]string{
	StatusOK: "200",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
