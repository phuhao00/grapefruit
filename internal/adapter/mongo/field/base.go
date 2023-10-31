package field

import "go.mongodb.org/mongo-driver/bson"

// ID field is constant for referencing the "_id" field name.
const ID = "_id"

// Empty is a predefined empty map.
var Empty = bson.M{}
