package model

type Student struct {
	StudentID     int    `gorm:"primary_key;auto_increment;column:student_id" json:"StudentId"`
	StudentName   string `gorm:"column:student_name" json:"StudentName"`
	SchoolClassID int
	SchoolClass   SchoolClass `gorm:"references:SchoolClassID"`
}
