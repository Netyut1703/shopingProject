package store

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/netyut1703/shopingProject/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var SellerCollection *mongo.Collection

//add seller

func (store *Store) addIteams(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	SellerCollection = client.Database("shopingProject").Collection("sellerdata")

	var sell *models.Seller
	err := json.NewDecoder(r.Body).Decode(&sell)

	if err != nil {
		fmt.Fprintln(w, "Failed to add seller")
		return
	}

	_, err = SellerCollection.InsertOne(ctx, sellerdata)
	if err != nil {
		fmt.Fprintln(w, "Failed to add seller")
		return
	}

	fmt.Print("seller added")

}

//update seller info

func (store *Store) updateSellerInfo(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	SellerCollection = client.Database("shopingProject").Collection("sellerdata")
	hexId := r.URL.Query().Get("id")
	objID, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		panic(err)
	}

	var sell Seller

	err = json.NewDecoder(r.Body).Decode(&sell)
	if err != nil {
		fmt.Fprintln(w, "Failed to update info")
		return
	}
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	update := bson.M{"$set": bson.M{
		"fullname":    sell.Fullname,
		"address":     sell.Address,
		"gender":      sell.Gender,
		"dateofbirth": sell.DateOfBirth,
		"location":    sell.Location,
	}}
	_, err = SellerCollection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		fmt.Println(err)
	}
}
