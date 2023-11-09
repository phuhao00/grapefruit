package do

import "time"

type Photos struct {
	Id          uint64    `json:"id"`
	Url         string    `json:"url"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
	Description string    `json:"description"`
	LikeTimes   int64     `json:"likeTimes"`
}
