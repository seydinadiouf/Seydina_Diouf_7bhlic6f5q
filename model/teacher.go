package model

type Teacher struct {
	TeacherId     int    `gorm:"column:teacher_id" json:"TeacherId"`
	TeacherName   string `gorm:"column:teacher_name" json:"TeacherName"`
	SchoolClassID int
	SchoolClass   SchoolClass `gorm:"references:SchoolClassID"`
}
