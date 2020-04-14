package services

import (
	"context"
	"errors"
	"mqttcs/student"
)

type StudentManager struct {
}

func (s StudentManager) GetStudent(ctx context.Context, request *student.StudentRequest, response *student.Student) error {
	studentMap := map[string]student.Student{
		"davie":  student.Student{Name: "davie", Class: "软件工程", Grade: 80},
		"steven": student.Student{Name: "steven", Class: "计算机科学", Grade: 93},
		"tony":   student.Student{Name: "tony", Class: "电子网络", Grade: 55},
		"jack":   student.Student{Name: "jack", Class: "工商管理", Grade: 70},
	}

	if request.Name == "" {
		return errors.New("请求参数错误，请重新请求")
	}

	student := studentMap[request.Name]
	if student.Name != "" {
		// response = &student
		response.Name = student.Name
		response.Class = student.Class
		response.Grade = student.Grade
		return nil
	}
	return errors.New("未查到相关信息")

}
