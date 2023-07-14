// shoosh really want a global constants
// between backend and frontend
package constant

import (
	"encoding/json"
)

type constants struct {
	RegMaxPasswordLen int
	RegMaxNicknameLen int
	RegMaxEmailLen    int
	MaxHttpBodyLen    int64
	MaxDeviceIdLen    int
	MaxPowLen         int
}

var (
	inited bool

	CJson []byte
	C     = constants{
		RegMaxPasswordLen: 64,
		RegMaxNicknameLen: 64,
		RegMaxEmailLen:    256,
		MaxHttpBodyLen:    16384,
		MaxDeviceIdLen:    4096,
		MaxPowLen:         32,
	}
)

func init() {
	if inited {
		return
	}
	inited = true

	CJson, _ = json.Marshal(C)
}
