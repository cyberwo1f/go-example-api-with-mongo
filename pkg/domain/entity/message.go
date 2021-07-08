package entity

//go:generate gomodifytags --file $GOFILE --struct Message -add-tags json,bson -w -transform camelcase
type Message struct {
	Id      int    `json:"id" bson:"id"`
	UserId  int    `json:"userId" bson:"userId"`
	Message string `json:"message" bson:"message"`
}
