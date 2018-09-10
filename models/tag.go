package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Tags struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name" validate:"required"`
	CreatedBy  string        `bson:"createdBy" json:"created_by"`
	ModifiedBy string        `bson:"modifiedBy" json:"modified_by"`
	State      int           `bson:"state" json:"state"`
}

const (
	db         = "mu_ti"
	collection = "mu_ti_tag"
)

func (t *Tags) InsertTag(tag Tags) error {
	return Insert(db, collection, tag)
}

func (t *Tags) FindAllTags() ([]Tags, error) {
	var result []Tags
	err := FindAll(db, collection, nil, nil, &result)
	return result, err
}

func (t *Tags) FindTagById(id string) (Tags, error) {
	var result Tags
	err := FindOne(db, collection, bson.M{"_id": bson.ObjectIdHex(id)}, nil, &result)
	return result, err
}

func (t *Tags) UpdateTag(tag Tags) error {
	return Update(db, collection, bson.M{"_id": tag.Id}, tag)
}

func (t *Tags) RemoveTag(id string) error {
	return Remove(db, collection, bson.M{"_id": bson.ObjectIdHex(id)})
}
