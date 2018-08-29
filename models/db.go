package models

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

var globalS *mgo.Session

func InitData() {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{viper.GetString("database.DBHost")},
		Source:   viper.GetString("database.Source"),
		Username: viper.GetString("database.User"),
		Password: viper.GetString("database.DBPass"),
	}

	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		logrus.Fatalf("Create Session: %s\n", err)
	} else {
		logrus.Infoln("Mongodb commit successful!")
	}
	globalS = s
}

//每一次操作都copy一份 Session,避免每次创建Session,导致连接数量超过设置的最大值
//获取文档对象 c := Session.DB(db).C(collection)
func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	s := globalS.Copy()
	c := s.DB(db).C(collection)
	return s, c
}

//插入操作
func Insert(db, collection string, docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs...)
}

//判断是否存在
func IsExist(db, collection string, query interface{}) bool {
	ms, c := connect(db, collection)
	defer ms.Close()
	count, _ := c.Find(query).Count()
	return count > 0
}

//查询单个对象
func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

//查询所有对象
func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

//更新操作
func Update(db, collection string, query, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Update(query, update)
}

//删除操作
func Remove(db, collection string, query interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Remove(query)
}
