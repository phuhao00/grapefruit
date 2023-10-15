package tcp

type Message struct {
	Command uint64 `json:"command"`
	Data    []byte `json:"data"`
}
