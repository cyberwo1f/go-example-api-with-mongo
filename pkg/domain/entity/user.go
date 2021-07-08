package entity

//go:generate gomodifytags --file $GOFILE --struct User -add-tags json,bson -w -transform camelcase
type User struct {
	Id   int    `json:"id" bson:"id"`
	Name string `json:"userName" bson:"name"`
}
