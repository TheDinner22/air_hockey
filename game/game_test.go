package game

import (
    "testing"
)

func TestCirlceContains(t *testing.T) {
    c1 := Circle{ 0, 0, 1}
    c2 := Circle{ 2, 0, 1}

    if c1.contains(c2) {
        t.Fatalf("case 1")
    }

    c1 = Circle{ 0, 0, 100}
    c2 = Circle{ 1, 0, 1}

    if !c1.contains(c2) {
        t.Fatalf("case 2")
    }

    c1 = Circle{ 10, 10, 10}
    c2 = Circle{ 5, 5, 1}

    if !c1.contains(c2) {
        t.Fatalf("case 3")
    }

    c1 = Circle{ 0, 0, 0}
    c2 = Circle{ 0, 0, 0}

    if c1.contains(c2) {
        t.Fatalf("case 4")
    }

    c1 = Circle{ 1, 1, 1}
    c2 = Circle{ 1, 1, 1}

    if !c1.contains(c2) {
        t.Fatalf("case 5")
    }
}
