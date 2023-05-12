package service

import (
	"math"
	"school-manager/school-manager-service/config"
	"school-manager/school-manager-service/dto"
	"school-manager/school-manager-service/model"
	"strconv"
	"strings"
)

type NamedArgument struct {
	IsSchoolClassNameNull bool
	IsTeacherNameNull     bool
	SchoolClassName       string
	TeacherName           string
	Limit                 int
	Offset                int
}
type PaginationResponse struct {
	TotalCount int              `json:"totalCount"`
	PageCount  int              `json:"pageCount"`
	PageNumber int              `json:"pageNumber"`
	PageSize   int              `json:"pageSize"`
	Data       []dto.StudentDTO `json:"students"`
}

func FilterStudent(SchoolClassName string, TeacherName string, Limit string, OffSet string) (*PaginationResponse, error) {
	db := config.DB()

	limit, err := strconv.Atoi(Limit)
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(OffSet)
	if err != nil {
		offset = 0
	}

	IsSchoolClassNameNull := len(SchoolClassName) == 0
	IsTeacherNameNull := len(TeacherName) == 0

	var students []model.Student

	var totalCount int64

	_ = db.Model(&model.Student{}).Count(&totalCount).Error

	namedArgument := NamedArgument{
		IsSchoolClassNameNull: IsSchoolClassNameNull,
		SchoolClassName:       strings.ToLower(SchoolClassName),
		IsTeacherNameNull:     IsTeacherNameNull,
		TeacherName:           strings.ToLower(TeacherName),
		Offset:                offset,
		Limit:                 limit,
	}

	if res := db.Preload("SchoolClass").Preload("SchoolClass.Teacher").Raw("SELECT * "+
		"FROM students s "+
		"LEFT JOIN classes c ON s.school_class_id = c.school_class_id "+
		"LEFT JOIN teachers t ON c.school_class_id = t.school_class_id "+
		"WHERE (@IsSchoolClassNameNull OR  LOWER(c.school_class_name) = @SchoolClassName) "+
		"AND (@IsTeacherNameNull OR  LOWER(t.teacher_name) = @TeacherName) OFFSET @Offset  LIMIT @Limit ", namedArgument).
		Model(&model.Student{}).Find(&students); res.Error != nil {
		return nil, res.Error
	} else {
		pageCount := int(math.Ceil(float64(totalCount) / float64(limit)))
		pageNumber := int(math.Floor(float64(offset)/float64(limit))) + 1
		pageSize := len(students)
		response := PaginationResponse{
			TotalCount: int(totalCount),
			PageCount:  pageCount,
			PageNumber: pageNumber,
			PageSize:   pageSize,
			Data:       (&dto.StudentDTOMapper{}).ToDTOs(students),
		}
		return &response, nil
	}

}
