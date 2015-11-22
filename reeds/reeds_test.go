package reeds

import (
    "testing"
)

func TestStoreReed(t *testing.T) {

    reed := Reed{ "Mostly harmless" }

    err := Store(&reed)

    if err != nil {
        t.Errorf("No error shuld be thrown got %d", err)
    }
}
