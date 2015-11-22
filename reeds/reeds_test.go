package reeds

import (
    "testing"
)

func TestStoreReed(t *testing.T) {

    repo, err := NewReedsRepo( "localhost", "test", "reedstest" )

    if err != nil {
        t.Errorf("Could not create reeds repo for tests %d", err)
    }

    reed := Reed{ "Mostly harmless" }

    err = repo.Store(&reed)

    if err != nil {
        t.Errorf("No error shuld be thrown got %d", err)
    }
}
