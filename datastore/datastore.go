package datastore

import (
	"Go-Lang-Zopsmart/model"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	client *mongo.Client
}

func New() *Student {
	return &Student{}
}

func (s *Student) connectMongoDB(ctx *gofr.Context) *mongo.Collection {
	if s.client == nil {
		uri := "mongodb+srv://tavisshiChauhan:tavishi@cluster0.mz8vc3l.mongodb.net/?retryWrites=true&w=majority"
		if uri == "" {
			log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
		}
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		s.client = client
	}
	fmt.Println("Connected to MongoDB")
	return s.client.Database("sample_mflix").Collection("students")
}

func (s *Student) GetByID(ctx *gofr.Context, ID string) (*model.Student, error) {
	coll := s.connectMongoDB(ctx)
	var result model.Student
	i, errr := strconv.Atoi(ID)
	if errr != nil {
		return nil, errr
	}
	fmt.Println("ID:", i)

	err := coll.FindOne(context.Background(), bson.M{"id": i}).Decode(&result)
	if err != nil {
		fmt.Println("FindOne ERROR:", err)
		return nil, err
	}

	fmt.Println(coll.FindOne(context.Background(), bson.M{"id": i}))
	return &result, nil
}

func (s *Student) Create(ctx *gofr.Context, student *model.Student) (*model.Student, error) {

	coll := s.connectMongoDB(ctx)

	// MongoDB collection and data
	data := model.Student{ID: student.ID, Name: student.Name, Age: student.Age, Class: student.Class}
	result, insertErr := coll.InsertOne(ctx, data)
	if insertErr != nil {
		fmt.Println("InsertOne ERROR:", insertErr)
		os.Exit(1)
	} else {
		fmt.Println("Data inserted with objectID: ", result.InsertedID)
	}

	return &data, nil
}

func (s *Student) Update(ctx *gofr.Context, student *model.Student) (*model.Student, error) {
	coll := s.connectMongoDB(ctx)

	existingStudent, err := s.GetByID(ctx, strconv.Itoa(student.ID))
	if err != nil {
		fmt.Println("Error fetching existing student data:", err)
		return nil, err
	}

	updatedStudent := merge(existingStudent, student)

	filter := bson.M{"id": student.ID}
	update := bson.M{"$set": updatedStudent}

	_, err = coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Failed to update student")
		return nil, errors.DB{Err: err}
	}

	fmt.Println("Student updated successfully", student.ID)
	return updatedStudent, nil
}

func merge(existing *model.Student, update *model.Student) *model.Student {
	if update.Name != "" {
		existing.Name = update.Name
	}
	if update.Age != 0 {
		existing.Age = update.Age
	}
	if update.Class != "" {
		existing.Class = update.Class
	}
	return existing
}



func (s *Student) Delete(ctx *gofr.Context, id int) error {
	coll := s.connectMongoDB(ctx)
	
	filter := bson.M{"id": id}
	result, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of documents deleted: %d\n", result.DeletedCount)
	return nil
}
