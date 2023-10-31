package mongo

import "go.mongodb.org/mongo-driver/mongo"

// Database contains mongo.Database
type Database struct {
	database *mongo.Database
}

// Collection returns database
func (d *Database) Collection(collection string) *Collection {
	return &Collection{Collection: d.database.Collection(collection)}
}
