package collections

import "awesomeProject/Practice28-Mockery-Rest-Mongo/pkg/model"

var (
	SampleStudents = []*model.StudentDetails{
		{
			Id:     "61083181cf3c682d2ce79867",
			Name:   "Jasmeet",
			Rollno: 2,
			Age:    27,
			Class:  10,
		},
		{
			Id:     "6108cfaccf3c6827ac3470c4",
			Name:   "Jasmeet",
			Rollno: 2,
			Age:    27,
			Class:  10,
		},
	}
	SampleCreateStudent = []model.StudentDetails{
		{
			Id:     "61083181cf3c682d2ce79867",
			Name:   "Jasmeet",
			Rollno: 2,
			Age:    27,
			Class:  10,
		},
	}

	SampleCreateStudentBuffer = []byte(`{
			Id: "61083181cf3c682d2ce79867",
			Name: "Jasmeet",
			Rollno: 2,
			Age: 27,
			Class: 10,
		}`)
)
