package game

import (
    "testing"
)

func TestCirlceContains(t *testing.T) {
    c1 := Circle{ Point{0, 0}, 1}
    c2 := Circle{ Point{2, 0}, 1}

    if c1.contains(c2) {
        t.Fatalf("case 1")
    }

    c1 = Circle{ Point{0, 0}, 100}
    c2 = Circle{ Point{1, 0}, 1}

    if !c1.contains(c2) {
        t.Fatalf("case 2")
    }

    c1 = Circle{ Point{10, 10}, 10}
    c2 = Circle{ Point{5, 5}, 1}

    if !c1.contains(c2) {
        t.Fatalf("case 3")
    }

    c1 = Circle{ Point{0, 0}, 0}
    c2 = Circle{ Point{0, 0}, 0}

    if c1.contains(c2) {
        t.Fatalf("case 4")
    }

    c1 = Circle{ Point{1, 1}, 1}
    c2 = Circle{ Point{1, 1}, 1}

    if !c1.contains(c2) {
        t.Fatalf("case 5")
    }
}
