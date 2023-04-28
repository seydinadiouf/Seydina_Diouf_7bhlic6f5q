package model

type Student struct {
	StudentID     int    `gorm:"column:student_id" json:"StudentId"`
	StudentName   string `gorm:"column:student_name" json:"StudentName"`
	SchoolClassID int
}
