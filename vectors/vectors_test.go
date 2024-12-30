package vectors

import (
    "testing"
)

func TestCirlceContains(t *testing.T) {
    c1 := Circle{ Vec2{0, 0}, 1}
    c2 := Circle{ Vec2{2, 0}, 1}

    if c1.Contains(c2) {
        t.Fatalf("case 1")
    }

    c1 = Circle{ Vec2{0, 0}, 100}
    c2 = Circle{ Vec2{1, 0}, 1}

    if !c1.Contains(c2) {
        t.Fatalf("case 2")
    }

    c1 = Circle{ Vec2{10, 10}, 10}
    c2 = Circle{ Vec2{5, 5}, 1}

    if !c1.Contains(c2) {
        t.Fatalf("case 3")
    }

    c1 = Circle{ Vec2{0, 0}, 0}
    c2 = Circle{ Vec2{0, 0}, 0}

    if c1.Contains(c2) {
        t.Fatalf("case 4")
    }

    c1 = Circle{ Vec2{1, 1}, 1}
    c2 = Circle{ Vec2{1, 1}, 1}

    if !c1.Contains(c2) {
        t.Fatalf("case 5")
    }
}
