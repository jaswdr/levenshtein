package levenshtein

import (
    "testing"
    "gotest.tools/assert"
)

var toCompare = [][]interface{}{
    []interface{}{"HONDA", "HYUNDAI", 3},
    []interface{}{"HoNdA", "hYuNdAi", 3},
    []interface{}{"book", "back", 2},
    []interface{}{"back", "book", 2},
    []interface{}{"kitten", "sitting", 3},
    []interface{}{"sitting", "kitten", 3},
    []interface{}{"sitting", "sitting", 0},
    []interface{}{"kitten", "kitten", 0},
    []interface{}{"Saturday", "Sunday", 3},
    []interface{}{"Sunday", "Saturday", 3},
    []interface{}{"industry", "interest", 6},
    []interface{}{"horse", "ros", 3},
    []interface{}{"soylent green is people", "people soiled our green", 19},
}

func TestCompare(t *testing.T) {
    for _, values := range toCompare {
        from := values[0].(string)
        to := values[1].(string)
        diff := values[2].(int)
        assert.Equal(t, diff, Compare(from, to))
    }
}
