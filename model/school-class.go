package model

type SchoolClass struct {
	SchoolClassID   int    `gorm:"column:school_class_id" json:"schoolClassId"`
	SchoolClassName string `gorm:"column:school_class_name" json:"schoolClassName"`
}

func (SchoolClass) TableName() string {
	return "classes"
}
