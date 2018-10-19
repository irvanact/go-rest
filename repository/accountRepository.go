package repository

import (
	. "go-rest-sample/model"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AccountRepository struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "accounts"
)

func (m *AccountRepository) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *AccountRepository) FindAll() ([]Account, error) {
	var accounts []Account
	err := db.C(COLLECTION).Find(bson.M{}).All(&accounts)
	return accounts, err
}

func (m *AccountRepository) FindById(id string) (Account, error) {
	var account Account
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&account)
	return account, err
}

func (m *AccountRepository) Insert(account Account) error {
	err := db.C(COLLECTION).Insert(&account)
	return err
}

func (m *AccountRepository) Delete(account Account) error {
	err := db.C(COLLECTION).Remove(&account)
	return err
}

func (m *AccountRepository) Update(account Account) error {
	err := db.C(COLLECTION).UpdateId(account.ID, &account)
	return err
}
