package model

type Teacher struct {
	TeacherID   int    `gorm:"column:teacher_id" json:"TeacherId"`
	TeacherName string `gorm:"column:teacher_name" json:"TeacherName"`
}
