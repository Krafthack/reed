package reeds

import (
    "log"
)

func (repo ReedsRepo) Store(reed *Reed) error {

    err := repo.reeds.Insert(reed)
    if err != nil {
        log.Fatal(err)
        return err
    }

    return nil
}
