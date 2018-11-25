package protocol

const (
	Request = iota
	Response
)

const (
	MagicSize     = 2 // 魔术: 16bit
	FeaturesSize  = 1  // 功能: 5bit
	VersionSize   = 1  // 版本号: 3bit
	StatusSize    = 1  // 状态: 3bit
	SerializeSize = 1  // 序列化方式: 4bit
	ReserveSize   = 1  // 保留: 3bit
	RequestIdSize = 8 // 请求号: 64bit
	HeaderSize    = MagicSize + FeaturesSize + VersionSize +StatusSize + SerializeSize + ReserveSize +RequestIdSize
	Magic         = 0xdabb
)

type Header struct {
	MagicNum  uint16
	Features  uint8
	Version   uint8
	Status    uint8
	Serialize uint8
	Reserve   uint8
	RequestId uint64
}

func NewHeader(msgType int, proxy bool, serializeType int, status int, requestId uint64) *Header {
	var (
		features uint8 = 0x00 // 功能位
	)

	// 设置代理
	if proxy {
		features |= 0x2
	}

	// 消息类型
	if msgType == Request {
		features &= 0xfe
	} else {
		features |= 0x1
	}

	return &Header{
		MagicNum:  Magic,
		Features:  features,
		Status:    uint8(status),
		Serialize: uint8(serializeType),
		RequestId: requestId,
	}
}
