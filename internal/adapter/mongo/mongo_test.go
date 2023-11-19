package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

type Person struct {
	Name  string `bson:"name"`
	Age   int    `bson:"age"`
	Email string `bson:"email"`
}

func TestMongo(t *testing.T) {

	// 连接到 MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)

	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			fmt.Println(err)

		}
	}()

	// 选择要操作的数据库和集合
	collection := client.Database("mydatabase").Collection("people")

	// 插入数据
	insertResult, err := collection.InsertOne(ctx, Person{"Alice", 30, "alice@example.com"})
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("Inserted ID:", insertResult.InsertedID)

	// 查询数据
	var result Person
	filter := Person{Name: "Alice"}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("Found person:", result)
}
