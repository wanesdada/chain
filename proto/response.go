package proto

type BaseResponseBody struct {
	Ret  int
	Msg  string
	Err  string
	Data interface{}
}
