package builder

import (
	"go.mongodb.org/mongo-driver/bson"
	"grapefruit/kit/utils"
)

// appendIfHasVal appends the provided key and value to the map if the value is not nil.
func appendIfHasVal(m bson.M, key string, val interface{}) {
	if !utils.IsNil(val) {
		m[key] = val
	}
}
