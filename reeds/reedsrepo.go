package reeds

import (
	"log"

	"gopkg.in/mgo.v2"
)

type ReedsRepo struct {
	reeds *mgo.Collection
}

func NewReedsRepo(host string, db string, collection string) (ReedsRepo, error) {

	session, err := mgo.Dial(host)
	if err != nil {
		log.Println(err)
		return ReedsRepo{}, err
	}
	reeds := session.DB(db).C(collection)

	return ReedsRepo{reeds}, nil
}
