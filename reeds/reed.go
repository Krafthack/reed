package reeds

import (
    "time"
)

type Reed struct {
    Title string
    Excerpt string
    ImageUrl string
    Url string
    AddedAt time.Time
    CompletedAt time.Time
}
