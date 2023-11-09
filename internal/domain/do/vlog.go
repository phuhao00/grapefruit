package do

import "time"

type Vlog struct {
	Id          uint64    `json:"id"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
	LikeTimes   int64     `json:"likeTimes"`
}
