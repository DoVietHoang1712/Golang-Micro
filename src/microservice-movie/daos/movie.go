/*
 * @File: daos.movie.go
 * @Description: Implements Movie CRUD functions for MongoDB
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package daos

import (
	"Golang-Microservice/src/microservice-movie/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Movie manages Movie CRUD
type Movie struct {
	C *mongo.Collection
}

// COLLECTION of the database table
const (
	COLLECTION = "movies"
)

// GetAll gets the list of Movie
func (m *Movie) GetAll() ([]models.Movie, error) {
	// Get a collection to execute thbcegmorste query against.
	ctx := context.TODO()
	var err error
	var movies []models.Movie
	cursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &movies)
	if err != nil {
		return nil, err
	}
	return movies, err
}

// GetByID finds a Movie by its id
func (m *Movie) GetByID(id string) (models.Movie, error) {
	p, _ := primitive.ObjectIDFromHex(id)

	var movie models.Movie
	err := m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&movie)
	return movie, err
}

// Insert adds a new Movie into database'
func (m *Movie) Insert(movie models.Movie) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), &movie)
}

// Delete remove an existing Movie
func (m *Movie) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}

// Update modifies an existing Movie
func (m *Movie) Update(id string, movie models.Movie) *mongo.SingleResult {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	return m.C.FindOneAndUpdate(context.TODO(), bson.M{"_id": p}, movie)
}
