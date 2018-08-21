package models

import "gopkg.in/mgo.v2/bson"

type Tag struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name"`
	CreatedBy  string        `bson:"createdBy" json:"created_by"`
	ModifiedBy string        `bson:"modifiedBy" json:"modified_by"`
	State      int           `bson:"state" json:"state"`
}

const (
	db         = "Tags"
	collection = "TagModel"
)

func (t *Tag) InsertTag(tag Tag) error {
	return Insert(db, collection, tag)
}

func (t *Tag) FindAllTags(tag Tag) error {
	return FindAll(db, collection, nil, nil, tag)
}
