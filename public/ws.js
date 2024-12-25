var socket

/**
* @type {HTMLCanvasElement}
*/
var canvas
/**
* @type {CanvasRenderingContext2D}
*/
var ctx

document.addEventListener("DOMContentLoaded", (_) => {
    setup_canvas()
    // ws()
});

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

function setup_canvas (){
    const c = document.getElementById("my-canvas")

    // runtime code to make my lsp happy :(
    // TODO remove this in prod??? idk why not
    if (c instanceof HTMLCanvasElement) {
        canvas = c
        ctx = canvas.getContext('2d')
    }

    ctx.fillStyle = "black";
    ctx.fillRect(0, 0, canvas.width, canvas.height);
}

