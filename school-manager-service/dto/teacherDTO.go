package dto

import (
	"school-manager/school-manager-service/model"
)

type TeacherDTO struct {
	TeacherID   int    `json:"teacherId"`
	TeacherName string `json:"teacherName"`
}

type TeacherDTOMapper struct{}

func (m *TeacherDTOMapper) ToDTO(e *model.Teacher) *TeacherDTO {
	return &TeacherDTO{
		TeacherID:   e.TeacherID,
		TeacherName: e.TeacherName,
	}
}

func (m *TeacherDTOMapper) ToEntity(d *TeacherDTO) *model.Teacher {
	return &model.Teacher{
		TeacherID:   d.TeacherID,
		TeacherName: d.TeacherName,
	}
}
