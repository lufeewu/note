package main

import (
	"fmt"

	"github.com/casbin/casbin/persist"
	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v2"
	"github.com/globalsign/mgo"
	"github.com/sirupsen/logrus"
)

func AddPolicy() error {
	// a := mongodbadapter.NewAdapter("127.0.0.1:27017") // Your MongoDB URL.
	a := mongodbadapter.NewAdapter(fmt.Sprintf("mongodb+srv://%s:%s@%s:%d/%s?ssl=%v&&authSource=%s",
		"", "", "127.0.0.1", 27017, false, ""))

	e, err := casbin.NewEnforcer("authz_model.conf", a)
	if err != nil {
		logrus.Warnln(err)
		return err
	}
	e.EnableLog(true)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	ok, err := e.Enforce("alice", "data1", "read")
	if err != nil {
		return err
	}
	logrus.Infoln("check:", ok)

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)
	type object struct {
		ID    string
		Table string
	}
	ok, err = e.AddPolicy("alice", object{ID: "aaaa", Table: "proj"}, "read")
	if err != nil {
		return err
	}
	logrus.Infoln("add:", ok)
	return nil
}

func NewMongoAdapter() (adapter persist.Adapter) {
	return adapter
}
func MongoTest() error {

	c, err := mgo.ParseURL("127.0.0.1:27317")
	if err != nil {
		logrus.Infoln(err)
		return err
	}
	logrus.Infoln(c)

	b, err := mgo.ParseURL(fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?ssl=%t&&authSource=%s",
		"a", "a", "127.0.0.1", 27017, "terra", false, "admin"))
	if err != nil {
		logrus.Infoln(err)
		return err
	}
	logrus.Infoln(b)

	// a := mongodbadapter.NewAdapter(fmt.Sprintf("mongodb://%s:%d/%s?ssl=%t&&authSource=%s",
	// "127.0.0.1", 27017, "terra", false, "admin"))
	d := NewMongoAdapter()
	e, err := casbin.NewEnforcer("rbac_multi_model.conf", d)
	if err != nil {
		logrus.Warnln(err)
		return err
	}
	// e.EnableLog(true)

	// Load the policy from DB.
	// e.LoadPolicy()

	// Check the permission.
	ok, err := e.Enforce("alice", "pom", "data1", "read")
	if err != nil {
		return err
	}
	logrus.Infoln("check:", ok)

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)
	ok, err = e.AddPolicy("alice", "pom", "data1", "read")
	if err != nil {
		return err
	}
	logrus.Infoln("add:", ok)

	// Save the policy back to DB.
	e.SavePolicy()

	ok, err = e.Enforce("alice", "bm", "data1", "read")
	if err != nil {
		return err
	}
	logrus.Infoln("check:", ok)

	return nil
}

func main() {
	// load the casbin model and policy from files, database is also supported.
	err := MongoTest()
	logrus.Warnln(err)
}
