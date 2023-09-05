package services

import (
	"context"
	"fmt"
	"log"

	"project1/interfaces"
	"project1/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionService struct {
	client                *mongo.Client
	CustomerCollection    *mongo.Collection
	TransactionCollection *mongo.Collection
	ctx                   context.Context
}

func NewTransactionServiceInit(client *mongo.Client, Customercollection *mongo.Collection, TransactionCollection *mongo.Collection, ctx context.Context) interfaces.Ipayment {
	return &TransactionService{
		client:                client,
		CustomerCollection:    Customercollection,
		TransactionCollection: TransactionCollection,
		ctx:                   ctx,
	}

}

func (a *TransactionService) CreatePayment(cardno float64, cvvverified int, amount float64) (string, error) {

	session, err := a.client.StartSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.EndSession(context.Background())

	filter := bson.M{"cardno": cardno}

	var account *models.Paymentscard

	err1 := a.CustomerCollection.FindOne(context.Background(), filter).Decode(&account)
	if err1 != nil {
		return "error", err1
	}
	fmt.Println("going to in")

	if account.Balance > amount {
		fmt.Println("its in")

		_, err = session.WithTransaction(context.Background(), func(ctx mongo.SessionContext) (interface{}, error) {
			_, err := a.CustomerCollection.UpdateOne(context.Background(),
				bson.M{"cardno": cardno},
				bson.M{"$inc": bson.M{"balance": -amount}})
			if err != nil {
				fmt.Println("Transaction Failed")
				return nil, err
			}

			transactionToInsert := models.Payments{
				Id:          "T011",
				Customer_id: "",
				Status:      "success",
				Gateway:     "",
				Type:        "",
				Amount:      amount,
				Card:        models.Paymentscard{},
				Token:       "",
			}
			res, _ := a.TransactionCollection.InsertOne(context.Background(), &transactionToInsert)
			var newUser *models.Payments
			query := bson.M{"_id": res.InsertedID}
			fmt.Println(res.InsertedID)

			err3 := a.TransactionCollection.FindOne(a.ctx, query).Decode(&newUser)
			if err3 != nil {
				return nil, err3
			}
			return "success", nil
		})
		if err != nil {
			return "failed", err
		}
	} else {
		return "Insufficent balance", nil
	}
	return "success", nil
}
