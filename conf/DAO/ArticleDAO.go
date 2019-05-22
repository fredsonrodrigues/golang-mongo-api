package DAO

import (
	. "golang-api/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type ArticlesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

//Conectando ao Banco
func (a *ArticlesDAO) Connect() {
	session, err := mgo.Dial(a.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(a.Database)
}

//recuperando todas as matérias
func (a *ArticlesDAO) GetAll() ([]Article, error) {
	var articles []Article
	err := db.C("article").Find(bson.M{}).All(&articles)
	return articles, err
}

func (a *ArticlesDAO) GetByID(id string) (Article, error) {
	var article Article
	err := db.C("article").FindId(bson.ObjectIdHex(id)).One(&article)
	return article, err
}