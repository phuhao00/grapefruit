package do

import "time"

type Project struct {
	Name        string    `json:"name"`        //名称
	BeginTime   time.Time `json:"beginTime"`   //开始时间
	EndTime     time.Time `json:"endTime"`     //结束时间
	Description string    `json:"description"` //描述
	Skills      []string  `json:"skills"`      //技能，技术栈
	Works       []string  `json:"works"`       //职能，
}
