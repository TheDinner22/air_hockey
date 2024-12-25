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
    console.log("hi")
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

function setup_canvas (){
    const c = document.getElementById("my-canvas")

    // runtime code to make my lsp happy :(
    if (c instanceof HTMLCanvasElement) {
        canvas = c

        ctx = canvas.getContext('2d')
    }
}

function sendMessage(message) {
    socket.send(message);
}
