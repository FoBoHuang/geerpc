package codec

import "io"

type Header struct {
	ServiceMethod string //服务名和方法名
	Seq           uint64 //服务端请求的序号
	Error         string //错误信息
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodeFunc func(closer io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodeFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodeFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
