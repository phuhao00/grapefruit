package tcp

import (
	"fmt"
	"io"
)

var ReadZeroLenMsg = fmt.Errorf("%v readLen:%v", io.EOF, 0)
