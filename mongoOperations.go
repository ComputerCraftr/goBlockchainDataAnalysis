package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MongoConfig stores the configuration of mongodb to connect
type MongoConfig struct {
	Ip       string `json:"ip"`
	Database string `json:"database"`
}

var mongoConfig MongoConfig

func readMongodbConfig(path string) {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &mongoConfig)
}

func getSession() (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://" + mongoConfig.Ip)
	if err != nil {
		panic(err)
	}
	//defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session, err
}
func getCollection(session *mgo.Session, collection string) *mgo.Collection {

	c := session.DB(mongoConfig.Database).C(collection)
	return c
}
func saveBlock(c *mgo.Collection, block BlockModel) {
	//first, check if the item already exists
	result := BlockModel{}
	err := c.Find(bson.M{"hash": block.Hash}).One(&result)
	if err != nil {
		//item not found, so let's add a new entry
		err = c.Insert(block)
		check(err)
	} else {
		err = c.Update(bson.M{"hash": block.Hash}, &block)
		if err != nil {
			log.Fatal(err)
		}
	}

}
func saveNode(c *mgo.Collection, block BlockModel) {
	var node NodeModel
	node.Id = block.Hash
	node.Label = block.Hash
	node.Title = block.Hash
	node.Value = 1
	node.Shape = "dot"
}
func saveEdge(c *mgo.Collection, block BlockModel) {
}
