package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
type participants struct{
	ID					 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	Name				 string             `json:"name,omitempty" bson:"name,omitempty"`
	RSVP         string             `json:"rsvp,omitempty" bson:"rsvp,omitempty"`
}
type Meeting struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title		     string             `json:"title,omitempty" bson:"title,omitempty"`
	Start        time               `json:"start,omitempty" bson:"start,omitempty"`
	End          time               `json:"end,omitempty" bson:"end,omitempty"`
	tstamp       time               `json:"tstamp,omitempty" bson:"tstamp,omitempty"`
	Participants participants       `json:"participants,omitempty" bson:"participants,omitempty"`
}

func ScheduleMeeting(response http.ResponseWriter, request *http.Request){
	response.Header().Set("content-type","application/json")
	var meeting Meeting
	_= json.NewDecoder(request.body).Decode(&meeting)
	collection := client.Database("DligentCoder").collection("meeting")
	ctx,_ := context.WithTimeout(content.Background(), 10*time.Second)
	result,_ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)

}
func ListAllMeetings(response http.ResponseWriter, request *http.Request){
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var meeting Meeting
	collection := client.Database("DiligentCoder").Collection("meeting")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Meeting{ID: id}).Decode(&meeting)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(meeting)
}
func getMeeting(response http.ResponseWriter, request *http.Request){
	var id;
	id = Meeting.ID;
	if(request.ID == id)
	return request;
}
func ListAll(response http.ResponseWriter, request *http.Request){//list all of a participant
	response.Header().Set("content-type", "application/json")
	var participants []Participants
	collection := client.Database("DiligentCoder").Collection("participants")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var meeting Meeting
		cursor.Decode(&meeting)
		people = append(participants, meeting)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(meeting)
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/meeting", ScheduleMeeting).Methods("POST")
	router.HandleFunc("/participants", ListAllMeetings).Methods("GET")
	router.HandleFunc("/meeting/{id}", getMeeting).Methods("GET")
	router.HandleFunc("/meeting/{id}", ListAll).Methods("GET")
	http.ListenAndServe(":12345", router)
}
