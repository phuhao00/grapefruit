// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package po

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID       int32  `gorm:"column:id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Pwd      string `gorm:"column:pwd" json:"pwd"`
	Token    string `gorm:"column:token" json:"token"`
	Category string `gorm:"column:category" json:"category"`
	Email    string `gorm:"column:email" json:"email"`
	Status   int32  `gorm:"column:status" json:"status"`
	Role     int32  `gorm:"column:role" json:"role"`
	Avatar   string `gorm:"column:avatar;comment:头像(链接)" json:"avatar"` // 头像(链接)
	Photos   string `gorm:"column:photos" json:"photos"`
	Vlogs    string `gorm:"column:vlogs" json:"vlogs"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
