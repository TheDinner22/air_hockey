/* we will do web sockets later :P
var socket

function ws(){
    socket = new WebSocket('ws://localhost:8000/ws');

    socket.onopen = function(_) {
        // Handle connection open
        console.log("open")
    };

    socket.onmessage = function(event) {
        // Handle received message
        console.log(event.data)
    };

    socket.onclose = function(_) {
        // Handle connection close
        console.log("close")
    };

}

function sendMessage(message) {
    socket.send(message);
}
*/

var canvas: HTMLCanvasElement
var ctx: CanvasRenderingContext2D
var game_state: GameState

document.addEventListener("DOMContentLoaded", (_) => {
    setup_canvas()

    Player.radius = canvas.width / 10

    init_game()

    // ws()
});

function setup_canvas() {
    canvas = document.getElementById("my-canvas") as HTMLCanvasElement
    ctx = canvas.getContext('2d') as CanvasRenderingContext2D;
}

class Player {
    score: number
    name: string
    pos: [number, number]

    static radius: number

    constructor(name: string) {
        this.score = 0
        this.name = name
        this.pos = [-1, -1]
    }

    move(x: number, y: number) {
        this.pos[0] = x
        this.pos[1] = y
    }

    draw() {
        ctx.beginPath();
        ctx.fillStyle = "red";
        ctx.arc(this.pos[0], this.pos[1], Player.radius, 0, 2 * Math.PI);
        ctx.fill();
    }
}

class GameState {
    p1: Player
    p2: Player
    puck: [number, number]

    constructor(p1: string, p2: string) {
        this.p1 = new Player(p1)
        this.p2 = new Player(p2)
        this.puck = [-1, -1]
    }


    draw() {
        // background
        ctx.fillStyle = "black";
        ctx.fillRect(0, 0, canvas.width, canvas.height);

        // puck
        ctx.beginPath();
        ctx.fillStyle = "red";
        ctx.arc(this.puck[0], this.puck[1], canvas.width / 15, 0, 2 * Math.PI);
        ctx.fill();

        // p1
        this.p1.draw()

        // p2
        this.p2.draw()
    }
}

function init_game() {
    game_state = new GameState("foo", "bar")

    // put puck at middle
    game_state.puck = [canvas.width / 2, canvas.height / 2]

    // put p1 on bottom
    game_state.p1.move(canvas.width / 2, canvas.height - Player.radius)

    // put p2 on top
    game_state.p2.move(canvas.width / 2, Player.radius)

    game_state.draw()
}
