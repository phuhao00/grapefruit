package do

import "time"

type Vlog struct {
	Id          uint64
	Url         string
	Description string
	CreateTime  time.Time
	UpdateTime  time.Time
	LikeTimes   int64
}
