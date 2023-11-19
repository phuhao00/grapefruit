package mongo

import (
	"context"
	"fmt"
	"testing"
)

func TestMongoDemo(t *testing.T) {
	c, err := Connect(context.Background(), "mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	collection := c.Database("hhh").Collection("abc")
	insertResult, err := collection.InsertOne(context.Background(), Person{"Alice", 30, "alice@example.com"})
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(insertResult)

}
