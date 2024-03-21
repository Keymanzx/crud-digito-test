package pokemon

import (
	db "api-gin/src/db/mongo"
	"api-gin/src/models/pokemon"
	"context"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllPokemon retrieves all Pokémon from the Pokémon collection
func GetAllPokemon(pageIndex, pageSize string) ([]*pokemon.Pokemon, error) {
	var result []*pokemon.Pokemon

	// Call the ConnectToMongoDB function
	client, err := db.ConnectToMongoDB()
	if err != nil {
		// Handle error
		fmt.Println("Error connecting to MongoDB:", err)
		return nil, err
	}

	// Access Pokémon collection
	collection := client.Database("db_digito").Collection("tb_pokemon")

	// Define the pipeline with pagination stages
	pipeline := bson.A{}
	if pageIndex != "" && pageSize != "" {
		pageIdx, _ := strconv.Atoi(pageIndex)
		pageSz, _ := strconv.Atoi(pageSize)
		pipeline = append(pipeline, bson.D{{Key: "$skip", Value: (pageIdx - 1) * pageSz}})
		pipeline = append(pipeline, bson.D{{Key: "$limit", Value: pageSz}})
	}

	// Perform aggregation
	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor and decode each Pokémon
	for cursor.Next(context.Background()) {
		var p pokemon.Pokemon
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		result = append(result, &p)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetByPokemonID(pokemonID string) (*pokemon.Pokemon, error) {
	// Call the ConnectToMongoDB function
	client, err := db.ConnectToMongoDB()
	if err != nil {
		// Handle error
		fmt.Println("Error connecting to MongoDB:", err)
		return nil, err
	}
	defer client.Disconnect(context.Background()) // Close the connection when done

	// Convert the string ID to an ObjectID
	id, err := primitive.ObjectIDFromHex(pokemonID)
	if err != nil {
		return nil, err
	}

	// Access the Pokemon collection
	collection := client.Database("db_digito").Collection("tb_pokemon")

	// Define a filter to find the Pokemon by ID
	filter := bson.M{"_id": id}

	// Find the Pokemon
	var result pokemon.Pokemon
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Println("Error finding Pokemon:", err)
		return nil, err
	}

	return &result, nil
}

func CreatePokemon(pokemon *pokemon.Pokemon) error {
	// Call the ConnectToMongoDB function
	client, err := db.ConnectToMongoDB()
	if err != nil {
		// Handle error
		fmt.Println("Error connecting to MongoDB:", err)
		return err
	}

	collection := client.Database("db_digito").Collection("tb_pokemon")

	// Insert user document
	_, err = collection.InsertOne(context.Background(), pokemon)
	if err != nil {
		return err
	}

	fmt.Println("Pokemon inserted successfully!")

	return nil
}

// UpdateByID updates a user by their ID in the users collection
func UpdateByID(pkmID string, updatedPkm *pokemon.Pokemon) error {
	// Call the ConnectToMongoDB function
	client, err := db.ConnectToMongoDB()
	if err != nil {
		// Handle error
		fmt.Println("Error connecting to MongoDB:", err)
		return err
	}

	// Access users collection
	collection := client.Database("db_digito").Collection("tb_pokemon")

	// Convert the string ID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(pkmID)
	if err != nil {
		return err
	}

	// Define a filter to find the user by ID
	filter := bson.M{"_id": objectID}

	// Define an update to set the new user data
	update := bson.M{"$set": updatedPkm}

	// Update the user
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	fmt.Println("User updated successfully!")

	return nil
}

func DeleteByID(userID string) error {
	// Call the ConnectToMongoDB function
	client, err := db.ConnectToMongoDB()
	if err != nil {
		// Handle error
		fmt.Println("Error connecting to MongoDB:", err)
		return err
	}

	// Access users collection
	collection := client.Database("db_digito").Collection("tb_pokemon")

	// Convert the string ID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	// Define a filter to find the user by ID
	filter := bson.M{"_id": objectID}

	// Delete the user
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
