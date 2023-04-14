package model

type Student struct {
	Id           uint   `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Gender       string `json:"gender"`
	Address      string `json:"address"`
	ProgramStudy string `json:"programstudy"`
	GPA          int    `json:"gpa"`
}
