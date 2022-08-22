package models

type Filter struct {
	FilterType	string			`json:"filterType" bson:"filterType"`
	Value	string			`json:"value" bson:"value"`
}
