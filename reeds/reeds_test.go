package reeds

import (
    "testing"
    "time"
)

func TestStoreReed(t *testing.T) {

    repo, err := NewReedsRepo( "localhost", "test", "reedstest" )

    if err != nil {
        t.Errorf("Could not create reeds repo for tests %d", err)
    }

    reed := Reed{
        "Rarely say yes to feature requests",
        "Here’s a simple set of Yes/No questions that you can quickly answer before you add another item to your product roadmap.  Saying yes to a feature request – whether it’s a to an existing customer, a product enquiry, a teammate, or a manager – is immediately rewarding.",
        "https://blog.intercom.io/wp-content/uploads/2014/11/Rarely-Say-Yes-984.jpg",
        "https://blog.intercom.io/rarely-say-yes-to-feature-requests/?utm_medium=email&utm_source=email&utm_campaign=say-no-email",
        time.Now().Add(-time.Hour * 24 * 3),
        time.Now(),
    }

    err = repo.Store(&reed)

    if err != nil {
        t.Errorf("No error shuld be thrown got %d", err)
    }
}
