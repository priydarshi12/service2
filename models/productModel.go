package models
import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name string `json:"name"`	
	Type string `json:"type"`
	Mfd  string `json:"mfd"`	
	Mrp int `json:"mrp"`
}

