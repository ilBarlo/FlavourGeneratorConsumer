package flavourmanager

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://ilBarlo:FlavourGenerator@clusterbarlo.qnlqwmd.mongodb.net/?retryWrites=true&w=majority"
const dbName = "flavours"
const colName = "resources"

var collection *mongo.Collection

// init connect with mongoDB

func init() {
	// Client option
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to mongodb
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection istance is ready")
}

// UpsertFlavour insert or update a new Flavour characterized by its UID
func upsertFlavour(flavour *Flavour) error {
	// create a filter to find the document with the same UID
	filter := bson.M{"uid": flavour.UID}

	// create an update that replaces the whole document with the new flavour
	update := bson.M{"$set": flavour}

	// try to update the document with the same UID, or insert the new document if it doesn't exist
	result, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	if result.MatchedCount == 1 {
		fmt.Printf("Updated flavour with UID %s\n", flavour.UID)
	} else if result.UpsertedCount == 1 {
		fmt.Printf("Inserted new flavour with UID %s\n", flavour.UID)
	}

	return nil
}
