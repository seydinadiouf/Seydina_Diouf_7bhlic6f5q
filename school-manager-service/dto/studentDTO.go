package dto

import (
	"school-manager/school-manager-service/model"
)

type StudentDTO struct {
	StudentID   int             `json:"studentId"`
	StudentName string          `json:"studentName"`
	SchoolClass *SchoolClassDTO `json:"schoolClass"`
}

type StudentDTOMapper struct{}

func (m *StudentDTOMapper) ToDTO(e *model.Student) *StudentDTO {
	return &StudentDTO{
		StudentID:   e.StudentID,
		StudentName: e.StudentName,
		SchoolClass: (&SchoolClassDTOMapper{}).ToDTO(&e.SchoolClass),
	}
}

func (m *StudentDTOMapper) ToEntity(d *StudentDTO) *model.Student {
	return &model.Student{
		StudentID:   d.StudentID,
		StudentName: d.StudentName,
		SchoolClass: *(&SchoolClassDTOMapper{}).ToEntity(d.SchoolClass),
	}
}

func (m *StudentDTOMapper) ToDTOs(students []model.Student) []StudentDTO {
	var studentDTOs []StudentDTO
	for _, student := range students {
		studentDTO := m.ToDTO(&student)
		studentDTOs = append(studentDTOs, *studentDTO)
	}
	return studentDTOs
}
