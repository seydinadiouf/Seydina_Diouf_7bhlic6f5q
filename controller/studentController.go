package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
	"net/http"
	"school-manager/config"
	"school-manager/model"
	"school-manager/model/payload"
	"strings"
)

func GetStudents(c echo.Context) error {
	SchoolClassName := c.QueryParam("schoolClassName")
	TeacherName := c.QueryParam("teacherName")

	type NamedArgument struct {
		IsSchoolClassNameNull bool
		IsTeacherNameNull     bool
		SchoolClassName       string
		TeacherName           string
	}

	IsSchoolClassNameNull := len(SchoolClassName) == 0
	IsTeacherNameNull := len(TeacherName) == 0

	db := config.DB()

	pg := paginate.New()

	if res := db.Raw("SELECT s.student_id, s.student_name, s.school_class_id "+
		"FROM students s "+
		"LEFT JOIN classes c ON s.school_class_id = c.school_class_id "+
		"LEFT JOIN teachers t ON c.school_class_id = t.school_class_id "+
		"WHERE (@IsSchoolClassNameNull OR  LOWER(c.school_class_name) = @SchoolClassName) "+
		"AND (@IsTeacherNameNull OR  LOWER(t.teacher_name) = @TeacherName)", NamedArgument{IsSchoolClassNameNull: IsSchoolClassNameNull, SchoolClassName: strings.ToLower(SchoolClassName), IsTeacherNameNull: IsTeacherNameNull, TeacherName: strings.ToLower(TeacherName)}).
		Model(&model.Student{}); res.Error != nil {

		data := map[string]interface{}{
			"message": res.Error.Error(),
		}
		return c.JSON(http.StatusOK, data)

	} else {

		return c.JSON(http.StatusOK, pg.With(res).Request(c.Request()).Response(&[]model.Student{}))
	}

}

func AddStudentToClass(c echo.Context) error {
	db := config.DB()

	// Parse the request body
	var requestBody payload.StudentToClass
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	student, err := getStudentByID(db, requestBody.StudentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "failed to find student"})
	}

	schoolClass, err := getSchoolClassByName(db, requestBody.SchoolClassName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "failed to find school class"})
	}

	// Add student to school class
	if err := db.Model(&student).Where("student_id = ?", requestBody.StudentID).Update("school_class_id", schoolClass.SchoolClassID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to add student to school class"})
	}

	// Return a success response
	return c.JSON(http.StatusOK, &student)
}

func getStudentByID(db *gorm.DB, studentID int) (*model.Student, error) {
	var student model.Student
	if err := db.Where("student_id = ?", studentID).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func getSchoolClassByName(db *gorm.DB, schoolClassName string) (*model.SchoolClass, error) {
	var schoolClass model.SchoolClass
	if err := db.Where("school_class_name = ?", schoolClassName).First(&schoolClass).Error; err != nil {
		return nil, err
	}
	return &schoolClass, nil
}
