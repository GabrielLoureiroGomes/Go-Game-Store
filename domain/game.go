package domain

import (
	"context"

	db "gostore/database"

	bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

type Game struct {
	Id             primitive.ObjectID `bson:"_id"`
	Name           string             `bson:"name"`
	Producer       string             `bson:"producer"`
	Platform       string             `bson:"platform"`
	ParentalRating int                `bson:"parentalRating"`
	Cooperative    bool               `bson:"cooperative"`
	Rating         int                `bson:"rating"`
	Comment        string             `bson:"comment"`
}

func GetAllGames() ([]Game, error) {

	filter := bson.D{{}}
	games := []Game{}

	client, err := db.GetMongoClient()

	if err != nil {
		return games, err
	}

	collection := client.Database(db.DB).Collection(db.GAMES)

	cur, findError := collection.Find(context.TODO(), filter)

	if findError != nil {
		return games, findError
	}

	for cur.Next(context.TODO()) {

		t := Game{}
		err := cur.Decode(&t)

		if err != nil {
			return games, err
		}

		games = append(games, t)
	}

	cur.Close(context.TODO())

	if len(games) == 0 {
		return games, mongo.ErrNoDocuments
	}

	return games, nil
}

func GetGameByName(name string) (Game, error) {

	result := Game{}

	filter := bson.D{primitive.E{Key: "name", Value: name}}

	client, err := db.GetMongoClient()

	if err != nil {
		return result, err
	}

	collection := client.Database(db.DB).Collection(db.GAMES)

	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func CreateGame(game Game) error {

	client, err := db.GetMongoClient()

	if err != nil {
		return err
	}

	collection := client.Database(db.DB).Collection(db.GAMES)

	_, err = collection.InsertOne(context.TODO(), game)

	if err != nil {
		return err
	}

	return nil
}

func DeleteGame(gameName string) error {

	filter := bson.D{primitive.E{Key: "name", Value: gameName}}

	client, err := db.GetMongoClient()

	if err != nil {
		return err
	}

	collection := client.Database(db.DB).Collection(db.GAMES)

	_, err = collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}

func UpdateGame(game Game) error {

	filter := bson.D{primitive.E{Key: "name", Value: game.Name}}

	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "name", Value: game.Name},
		{Key: "producer", Value: game.Producer},
		{Key: "platform", Value: game.Platform},
		{Key: "parentalRating", Value: game.ParentalRating},
		{Key: "cooperative", Value: game.Cooperative},
		{Key: "rating", Value: game.Rating},
	}}}

	client, err := db.GetMongoClient()

	if err != nil {
		return err
	}

	collection := client.Database(db.DB).Collection(db.GAMES)

	_, err = collection.UpdateOne(context.TODO(), filter, updater)

	if err != nil {
		return err
	}

	return nil
}
