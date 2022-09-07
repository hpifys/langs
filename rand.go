package langs

import (
    "errors"
    "math/rand"
)

func RandIntRange(min, max int) int {
    if min < 0 || max < 0 || max < min {
        panic(errors.New("params invalid"))
    }
    return rand.Intn(max-min+1) + min
}
