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

    init_game()

    // ws()
});

function setup_canvas() {
    canvas = document.getElementById("my-canvas") as HTMLCanvasElement
    ctx = canvas.getContext('2d') as CanvasRenderingContext2D;

    ctx.fillStyle = "black";
    ctx.fillRect(0, 0, canvas.width, canvas.height);
}

class Player {
    score: number
    name: string
    pos: [number, number]

    constructor(name: string) {
        this.score = 0
        this.name = name
        this.pos = [-1, -1]
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
}

function init_game() {
    game_state = new GameState("joe", "ella")

    // put puck at middle
    ctx.beginPath();
    ctx.fillStyle = "red";
    ctx.arc(canvas.width / 2, canvas.height / 2, canvas.width / 15, 0, 2 * Math.PI);
    ctx.fill();

    // put p1 on bottom
    // put p2 on top
}

