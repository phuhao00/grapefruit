// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package po

const TableNamePhoto = "photos"

// Photo mapped from table <photos>
type Photo struct {
	ID          int64  `gorm:"column:Id" json:"Id"`
	Desc        string `gorm:"column:desc" json:"desc"`
	URL         string `gorm:"column:url" json:"url"`
	UpdatedTime string `gorm:"column:updated_time" json:"updated_time"`
	CreatedTime string `gorm:"column:created_time" json:"created_time"`
	LikeTimes   int64  `gorm:"column:like_times" json:"like_times"`
}

// TableName Photo's table name
func (*Photo) TableName() string {
	return TableNamePhoto
}