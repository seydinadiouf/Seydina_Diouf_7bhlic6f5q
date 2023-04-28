package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"school-manager/config"
	"school-manager/model"
)

func GetStudents(c echo.Context) error {
	teacherName := c.QueryParam("teacherName")
	schoolClassName := c.QueryParam("schoolClassName")

	db := config.DB()

	var students []*model.Student

	if res := db.Raw("SELECT s.student_id, s.student_name, s.school_class_id "+
		"FROM students s "+
		"JOIN classes c ON s.school_class_id = c.school_class_id "+
		"LEFT JOIN teachers t ON c.teacher_id = t.teacher_id "+
		"WHERE (c.school_class_name = ? OR ? IS NULL) "+
		"AND (t.teacher_name = ? OR ? IS NULL)", schoolClassName, schoolClassName, teacherName, teacherName).
		Scan(&students); res.Error != nil {

		data := map[string]interface{}{
			"message": res.Error.Error(),
		}
		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": students,
	}

	return c.JSON(http.StatusOK, response)
}
