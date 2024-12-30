package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoapi/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoLocalHostUri = "mongodb://localhost:27017/"

const dbName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(mongoLocalHostUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodb is connected ")
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("collection refrence is ready ")
}

func addOneMovie(movie models.Netflix) {
	result, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		fmt.Println("sorry couldn't insert movie to  the  add one movie function")
		log.Fatal(err)
	}
	fmt.Println("document insterted successfully", result)
}

func updateOneMovie(movieId string) {
	parsedId, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		fmt.Println("sorry couldn't parse the movie id to mongoId inside add one movie function")
		log.Fatal(err)
	}
	filer := bson.M{"_id": parsedId}
	update := bson.M{"$set": bson.M{
		"watched": true,
	}}
	res, err := collection.UpdateOne(context.Background(), filer, update)
	if err != nil {
		fmt.Println("sorry couldn't update this  inside the updated one movie")
		log.Fatal(err)
	}
	fmt.Println("so result is successfully saved ", res)
}

func deleteOneMovie(movieId string) {
	parsedId, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		fmt.Println("sorry couldn't parse the movie id to mongoId inside add one movie function")
		log.Fatal(err)
	}
	filer := bson.M{"_id": parsedId}

	res, err := collection.DeleteOne(context.Background(), filer)

	if err != nil {
		fmt.Println("sorry couldn't delete  the movie  inside the dlete one movie function")
		log.Fatal(err)
	}

	fmt.Println("the delttation is completed", res)
}

func deleteAll() {
	filer := bson.M{}
	res, err := collection.DeleteMany(context.Background(), filer)
	if err != nil {
		fmt.Println("sorry couldn't delete  the whole movies ")
		log.Fatal(err)
	}
	fmt.Println("the deltation of whole movie is completed", res)
}

func getAllMovies() []primitive.M {

	fmt.Println("get all movies 1 ")
	cursor, err := collection.Find(context.Background(), bson.M{})
	defer cursor.Close(context.Background())

	fmt.Println("get all movies 2 ")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("get all movies 3 ")

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(movies)
		}
		movies = append(movies, movie)
	}
	fmt.Println("get all movies 4 ")

	return movies
}

func AddOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie models.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}
	addOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsRead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	updateOneMovie(vars["id"])
}

func GetAllMyMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	deleteAll()
}

func DeleteOneMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	deleteOneMovie(vars["id"])
}
