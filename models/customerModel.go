package models
import "go.mongodb.org/mongo-driver/bson/primitive"
// Item represents a data model
type Customer struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	FName string `json:"fname"`	
	LName string `json:"lname"`	
	Email string `json:"email"`
}

