package models
import(
	
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	ID      int64  `json:"id"`
	Name   string `json:"name"`
	Email   string `json:"email"`
	Password int64  `json:"password"`
}
