package model

type SchoolClass struct {
	SchoolClassID   int     `gorm:"column:school_class_id" json:"schoolClassId"`
	SchoolClassName string  `gorm:"column:school_class_name" json:"schoolClassName"`
	TeacherID       int     `gorm:"column:teacher_id" json:"teacherId"`
	Teacher         Teacher `gorm:"references:TeacherID"`
}

func (SchoolClass) TableName() string {
	return "classes"
}
