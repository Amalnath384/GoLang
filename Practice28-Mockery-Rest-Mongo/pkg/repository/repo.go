package repository

import (
	"awesomeProject/Practice28-Mockery-Rest-Mongo/pkg/model"
	"awesomeProject/Practice28-Mockery-Rest-Mongo/pkg/repository/driver/mongo"
	"context"
	"fmt"
)

type Repository interface {
	CreateStudent(ctx context.Context, Student *model.StudentDetails) (*model.StudentDetails, error)
	DeleteStudent(ctx context.Context, id string) error
	UpdateStudent(ctx context.Context, Student *model.StudentDetails) (*model.StudentDetails, error)
	GetStudent(ctx context.Context, id string) (*model.StudentDetails, error)
	ListStudent(ctx context.Context) ([]*model.StudentDetails, error)
	Close()
}

var Repo Repository

func Init(db *model.Database) {
	switch db.Driver {
	case "etcd":
		// C = etcd.Init(db.Driver, db.Crendential)
		fmt.Printf("etcd is not implemented right now!")
		return
	case "mongodb":
		Repo = mongo.Init(db.Endpoint)
		return
	default:
		fmt.Printf("Can't find database driver %s!\n", db.Driver)
	}
}

func Exit() {
	Repo.Close()
}
