package configs

import (
	"context"
	model "data_logger/models"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Data").Collection(collectionName)
	return collection
}

// InsertData inserts data into the database
func InsertData(data *model.DataModel, collectionName string) {
	collection := GetCollection(DB, collectionName)

	var to_insert = bson.D{{"topic", data.Topic}, {"body", data.Payload}, {"uuid", data.Uuid}, {"datetime", data.Datetime}}

	_, err := collection.InsertOne(context.TODO(), to_insert)
	if err != nil {
		log.Fatal(err)
	}
}

// GetData gets data from the database
func GetData(collectionName string, indent bool) string {
	collection := GetCollection(DB, collectionName)
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	var output []byte

	if indent {
		output, err = json.MarshalIndent(results, "", "  ")
	} else {
		output, err = json.Marshal(results)
	}

	return string(output)
}

// DeleteData deletes data from the database
func DeleteData(collectionName string, topic string) {
	collection := GetCollection(DB, collectionName)
	_, err := collection.DeleteOne(context.Background(), bson.D{{"topic", topic}})
	if err != nil {
		log.Fatal(err)
	}
}

// Filter data from the database
func FilterData(collectionName string, Arg string, Value string, indent bool) string {
	collection := GetCollection(DB, collectionName)

	cursor, err := collection.Find(context.Background(), bson.D{{Arg, Value}})
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	var output []byte

	if indent {
		output, err = json.MarshalIndent(results, "", "  ")
	} else {
		output, err = json.Marshal(results)
	}

	return string(output)
}

func CountData(collectionName string) int64 {
	collection := GetCollection(DB, collectionName)
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	return count
}
