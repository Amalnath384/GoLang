package student

import (
	"awesomeProject/Practice28-Mockery-Rest-Mongo/pkg/repository"
	"awesomeProject/Practice28-Mockery-Rest-Mongo/test/collections"
	mockRepo "awesomeProject/Practice28-Mockery-Rest-Mongo/test/mocks"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)




func TestListStudent(t *testing.T) {
	req := httptest.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	//handler := http.HandlerFunc(NewStudentService().ListStudent)
	//handler.ServeHTTP(response, req)
	ctx := context.Background()
	mockRepoClient := new(mockRepo.Repository)
	mockRepoClient.On("ListStudent", ctx).Return(collections.SampleStudents, nil) //db
	repository.Repo = mockRepoClient //replacing with my database interface
	testService := NewStudentService()
	testService.ListStudent(response, req) //service
	//t.Log(err)
	assert.Equal(t, 200, response.Code, "OK response as expected")
	mockRepoClient.AssertExpectations(t)
}
func TestGetStudent(t *testing.T) {
	req := httptest.NewRequest("GET", "/students/id", nil)
	response := httptest.NewRecorder()
	fmt.Println(response.Body,response.Code)
	//handler := http.HandlerFunc(NewStudentService().ListStudent)
	//
	//handler.ServeHTTP(response, req)
	params := mux.Vars(req)
	id := params["id"]
	ctx := context.Background()
	mockRepoClient := new(mockRepo.Repository)
	mockRepoClient.On("GetStudent", ctx, id).Return(collections.SampleStudents[0], nil)
	repository.Repo = mockRepoClient
	testService := NewStudentService()
	testService.GetStudent(response, req)
	//t.Log(err)
	fmt.Println(response.Body,response.Code)
	assert.Equal(t, 200, response.Code, "OK response as expected")
	mockRepoClient.AssertExpectations(t)
}
func TestCreateStudent(t *testing.T) {
	//requestBody, _ := json.Marshal(&collections.SampleCreateStudent[0])
	req := httptest.NewRequest("POST", "/students", nil)
	//resp := model.StudentDetails{
	response := httptest.NewRecorder()
	json.NewDecoder(req.Body).Decode(&response)
	//response.Body
	//
	//data,_ := io.ReadAll(response.Body)
	//_ = json.Unmarshal(data,&resp)
	ctx := context.Background()
	mockRepoClient := new(mockRepo.Repository)
	mockRepoClient.On("CreateStudent", ctx, &collections.SampleCreateStudent[0]).Return(&collections.SampleCreateStudent[0], nil)
	repository.Repo = mockRepoClient
	testService := NewStudentService()
	//fmt.Println(response.Code, response.Body)
	testService.CreateStudent(response, req)

	fmt.Println(response.Body)
	assert.Equal(t, 200, response.Code, "OK response as expected")
	mockRepoClient.AssertExpectations(t)
}
func TestUpdateStudent(t *testing.T) {
	req := httptest.NewRequest("GET", "/students/id", nil)
	response := httptest.NewRecorder()
	//handler := http.HandlerFunc(NewStudentService().ListStudent)
	//
	//handler.ServeHTTP(response, req)
	params := mux.Vars(req)
	id := params["id"]
	ctx := context.Background()
	mockRepoClient := new(mockRepo.Repository)
	mockRepoClient.On("GetStudent", ctx, id).Return(collections.SampleStudents[0], nil)
	repository.Repo = mockRepoClient
	testService := NewStudentService()
	testService.GetStudent(response, req)
	//t.Log(err)
	assert.Equal(t, 200, response.Code, "OK response as expected")
	mockRepoClient.AssertExpectations(t)
	req1 := httptest.NewRequest("UPDATE", "/students/id", nil)
	response1 := httptest.NewRecorder()
	// //handler := http.HandlerFunc(NewStudentService().ListStudent)
	// //
	// //handler.ServeHTTP(response, req)
	//
	// params := mux.Vars(reqId)
	// id := params["id"]
	//fmt.Println(id)
	//
	ctx = context.Background()
	mockRepoClient1 := new(mockRepo.Repository)
	// mockRepoClient.On("GetStudent", ctx, id).Return(collections.SampleStudents[0], nil)
	mockRepoClient1.On("UpdateStudent", ctx, collections.SampleStudents[0]).Return(collections.SampleStudents[0], nil)
	repository.Repo = mockRepoClient1
	testService1 := NewStudentService()
	testService1.UpdateStudent(response1, req1)
	// //t.Log(err)
	//
	// //assert.Equal(t, 200, response.Code, "OK response as expected")
	mockRepoClient1.AssertExpectations(t)
}
func TestDeleteStudent(t *testing.T) {
	req := httptest.NewRequest("Delete", "/students/id", nil)
	response := httptest.NewRecorder()
	params := mux.Vars(req)
	id := params["id"]
	ctx := context.Background()
	mockRepoClient := new(mockRepo.Repository)
	mockRepoClient.On("DeleteStudent", ctx, id).Return( nil)
	repository.Repo = mockRepoClient
	testService := NewStudentService()
	testService.DeleteStudent(response, req)
	assert.Equal(t, 200, response.Code, "OK response as expected")
	mockRepoClient.AssertExpectations(t)
}


