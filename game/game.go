package game

type point struct {
    x int32
    y int32
}

type Player struct {
    name string
    score int8
    pos point
}

type Puck struct {
    pos point
    velocity point // in this case point is more like a vector dx, dy not x, y
}

type GameState struct {
    p1 Player
    p2 Player
    puck Puck
}
