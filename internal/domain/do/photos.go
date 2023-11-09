package do

import "time"

type Photos struct {
	Id          uint64
	Url         string
	CreateTime  time.Time
	UpdateTime  time.Time
	Description string
	LikeTimes   int64
}
