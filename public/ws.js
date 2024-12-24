var socket

document.addEventListener("DOMContentLoaded", (_) => {
    console.log("hi")
    ws()
});

function ws(){
    socket = new WebSocket('ws://localhost:8000/ws');

    socket.onopen = function(event) {
        // Handle connection open
        console.log("open")
    };

    socket.onmessage = function(event) {
        // Handle received message
        console.log(event.data)
    };

    socket.onclose = function(event) {
        // Handle connection close
        console.log("close")
    };

}

function sendMessage(message) {
    socket.send(message);
}
