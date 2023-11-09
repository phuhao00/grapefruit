package do

//Resume 简历
type Resume struct {
	Project []*Project `json:"project"`
	Vlog    []*Vlog    `json:"vlog"`
	Images  []*Photos  `json:"images"`
}
