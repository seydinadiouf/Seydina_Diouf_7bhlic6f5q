package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
	"school-manager/config"
	"school-manager/dto/payload"
	model2 "school-manager/model"
	pb "school-manager/proto"
	"school-manager/service"
	"strings"
)

func (*Server) GetStudentsWithFilter(ctx context.Context, in *pb.GetStudentRequest) (*pb.StudentsPage, error) {
	log.Printf("Get students was invoked with")

	var students []*pb.Student
	SchoolClassName := in.SchoolClassName
	TeacherName := in.TeacherName
	Limit := in.Limit
	Offset := in.Offset

	response, err := service.FilterStudent(SchoolClassName, TeacherName, Limit, Offset)

	studentsDTO := response.Data

	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Internal server error: %v", err),
		)
	}

	for _, studentData := range studentsDTO {
		teacher := &pb.Teacher{
			TeacherName: studentData.SchoolClass.Teacher.TeacherName,
			TeacherId:   int64(studentData.SchoolClass.Teacher.TeacherID),
		}
		schoolClass := &pb.SchoolClass{
			SchoolClassName: studentData.SchoolClass.SchoolClassName,
			SchoolClassId:   int64(studentData.SchoolClass.SchoolClassID),
			Teacher:         teacher,
		}
		student := &pb.Student{
			StudentName: studentData.StudentName,
			StudentId:   int64(studentData.StudentID),
			SchoolClass: schoolClass,
		}
		students = append(students, student)
	}

	return &pb.StudentsPage{
		TotalCount: int64(response.TotalCount),
		PageCount:  int64(response.PageCount),
		PageNumber: int64(response.PageNumber),
		PageSize:   int64(response.PageSize),
		Students:   students,
	}, nil

}

func (*Server) AddStudentToClass(ctx context.Context, in *pb.CreateStudentRequest) (*pb.Student, error) {
	db := config.DB()
	// Parse the request body
	requestBody := &payload.StudentToClass{
		SchoolClassName: in.SchoolClassName,
		StudentName:     in.StudentName,
	}

	schoolClass, err := getSchoolClassByName(db, requestBody.SchoolClassName)

	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Unable to find class with this name: %v", err),
		)
	}

	student := model2.Student{StudentName: requestBody.StudentName, SchoolClassID: schoolClass.SchoolClassID}

	res := db.Create(&student)

	res.Preload("SchoolClass").Preload("SchoolClass.Teacher").First(&student)

	teacherPb := &pb.Teacher{
		TeacherName: student.SchoolClass.Teacher.TeacherName,
		TeacherId:   int64(student.SchoolClass.Teacher.TeacherID),
	}
	schoolClassPb := &pb.SchoolClass{
		SchoolClassName: student.SchoolClass.SchoolClassName,
		SchoolClassId:   int64(student.SchoolClass.SchoolClassID),
		Teacher:         teacherPb,
	}
	studentPb := &pb.Student{
		StudentName: student.StudentName,
		StudentId:   int64(student.StudentID),
		SchoolClass: schoolClassPb,
	}
	// Return a success response
	return studentPb, nil
}

func getSchoolClassByName(db *gorm.DB, schoolClassName string) (*model2.SchoolClass, error) {
	var schoolClass model2.SchoolClass
	if err := db.Where("LOWER(school_class_name) = ?", strings.ToLower(schoolClassName)).First(&schoolClass).Error; err != nil {
		return nil, err
	}
	return &schoolClass, nil
}
