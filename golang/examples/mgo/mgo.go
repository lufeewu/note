package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Operater struct {
	mgoSession *mgo.Session
	dbname     string
	document   string
}

type Person struct {
	Age  int
	Name string
	High int
}

//连接数据库
func (operater *Operater) Connect(url string) error {
	mgoSession, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	operater.mgoSession = mgoSession
	return nil
}

// 数据库测试
func UpdateWithSelector() {
	var op = Operater{
		dbname:   "test",
		document: "person",
	}
	err := op.Connect("127.0.0.1:27017")
	if err != nil {
		logrus.Warnln(err)
		return
	}

	collection := op.mgoSession.DB(op.dbname).C(op.document)
	update := Person{
		33,
		"詹姆斯",
		201,
	}

	err = collection.Insert(update)
	if err != nil {
		logrus.Warnln(err)
		return
	}

	err = collection.Update(bson.M{"name": "jay chou"}, update)
	if err != nil {
		logrus.Warnln(err)
		return
	}

	return

}

func main() {
}
