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

function init_game() { }

