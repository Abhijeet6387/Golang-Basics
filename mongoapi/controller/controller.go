package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Abhijeet6387/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://abhijeetsmishra:abhi1234@cluster0.0ce5imf.mongodb.net/"
const dbName = "netflix"
const colName = "watchList"

// Most Important - This declares a global variable collection to hold the MongoDB collection instance, which will be used for database operations.
var collection *mongo.Collection

// Connect with mongoDB
	/* 
		The init function is automatically called when the package is initialized. It sets up the MongoDB client and connects to the database. 
		It applies the connection options defined in clientOption, connects to the MongoDB server specified in the connection string, and assigns the database collection to the collection variable for further use.
	*/
func init(){
	
	// client options
	clientOption := options.Client().ApplyURI(connectionString)
	
	// connect to db
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("MongoDB connnection sucessfull")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")

}

// MongoDB helpers - file

// Insert one record
func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie in DB with id: " ,inserted.InsertedID)
}

// Update one record
func updateOneMovie(movieId string){
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil{
		log.Fatal(err)
	}
	filter := bson.M{"_id":id}
	update := bson.M{"$set" : bson.M{"watched": true}}

	res, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Modified count: ", res.ModifiedCount)
}

// Delete one record
func deleteOneMovie(movieId string){
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil{
		log.Fatal(err)
	}
	filter := bson.M{"_id":id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DeleteCount : ", deleteCount)
}

// delete all records
func deleteAllMovie() int64 {
	
	// filter := bson.D{{}}
	delRes, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil{
		log.Fatal(err)
	}
	deletedmovies :=  delRes.DeletedCount
	fmt.Println("No. of movies deleted : ", deletedmovies)
	return deletedmovies
}

// get all movies from database - used to retrieve all movie documents from the MongoDB collection and return them as a slice of primitive.M (MongoDB BSON maps)
func getAllMovies() []primitive.M{
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil{
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()){
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil{
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())

	return movies
}

// Actual Controllers
func ServeHome(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	w.Write([]byte("<h1>Welcome to mongo API Home</h1>"))
}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allmovies := getAllMovies()
	json.NewEncoder(w).Encode(allmovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_  = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkasWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}