package reeds

import (
    "log"

    "gopkg.in/mgo.v2"
)

func Store(reed *Reed) error {
    session, err := mgo.Dial("localhost")
    if err != nil {
        log.Fatal(err)
        return err
    }
    defer session.Close()

    c := session.DB("reed").C("reeds")
    err = c.Insert(reed)
    if err != nil {
        log.Fatal(err)
        return err
    }

    return nil
}
