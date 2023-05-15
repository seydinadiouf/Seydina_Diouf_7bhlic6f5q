package dto

import (
	"school-manager/model"
)

type SchoolClassDTO struct {
	SchoolClassID   int         `json:"schoolClassId"`
	SchoolClassName string      `json:"schoolClassName"`
	Teacher         *TeacherDTO `json:"teacher"`
}

type SchoolClassDTOMapper struct{}

func (m *SchoolClassDTOMapper) ToDTO(e *model.SchoolClass) *SchoolClassDTO {
	return &SchoolClassDTO{
		SchoolClassID:   e.SchoolClassID,
		SchoolClassName: e.SchoolClassName,
		Teacher:         (&TeacherDTOMapper{}).ToDTO(&e.Teacher),
	}
}

func (m *SchoolClassDTOMapper) ToEntity(d *SchoolClassDTO) *model.SchoolClass {
	return &model.SchoolClass{
		SchoolClassID:   d.SchoolClassID,
		SchoolClassName: d.SchoolClassName,
		Teacher:         *(&TeacherDTOMapper{}).ToEntity(d.Teacher),
	}
}
