// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package po

const TableNameJob = "job"

// Job mapped from table <job>
type Job struct {
	ID        int32   `gorm:"column:id" json:"id"`
	Name      string  `gorm:"column:name" json:"name"`
	Desc      string  `gorm:"column:desc;comment:职位描述" json:"desc"`             // 职位描述
	MinSalary float64 `gorm:"column:min_salary;comment:最低薪资" json:"min_salary"` // 最低薪资
	MaxSalary float64 `gorm:"column:max_salary;comment:最高薪资" json:"max_salary"` // 最高薪资
	CompanyID int32   `gorm:"column:company_id" json:"company_id"`
	Require   string  `gorm:"column:require;comment:职位要求" json:"require"`         // 职位要求
	Publiser  int32   `gorm:"column:publiser;comment:发布职位的招聘者ID" json:"publiser"` // 发布职位的招聘者ID
}

// TableName Job's table name
func (*Job) TableName() string {
	return TableNameJob
}