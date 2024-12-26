"use strict";
var socket;
function ws_session_create() {
    if (socket != undefined && socket.OPEN) {
        console.error("cannot create session");
    }
    // now get that uu-id from the html returned
    const uuid_elem = document.getElementById("uuid");
    if (uuid_elem == null) {
        console.error("failed to get uuid! Was null");
        return;
    }
    const uuid_str = uuid_elem.innerText;
    const url = 'ws://localhost:8000/session/create?uuid=' + uuid_str;
    socket = new WebSocket(url);
    socket.onopen = function (_) {
        // Handle connection open
        console.log("created!");
    };
    socket.onmessage = function (event) {
        // Handle received message
        console.log(event.data);
    };
    socket.onclose = function (_) {
        // Handle connection close
        console.log("close");
    };
}
function ws_session_join() {
    if (socket != undefined && socket.OPEN) {
        console.error("cannot join session");
    }
    // get uuid from text box
    const text_input = document.getElementById("session-uuid");
    const uuid_to_join = text_input.value;
    const url = 'ws://localhost:8000/session/join?uuid=' + uuid_to_join;
    socket = new WebSocket(url);
    socket.onopen = function (_) {
        // Handle connection open
        console.log("joined!");
    };
    socket.onmessage = function (event) {
        // Handle received message
        console.log(event.data);
    };
    socket.onclose = function (_) {
        // Handle connection close
        console.log("close");
    };
}
function send_message(msg) {
    if (socket != undefined && !socket.OPEN) {
        return;
    }
    socket.send(msg);
}
/*
function sendMessage(message) {
    socket.send(message);
}
*/
var canvas;
var ctx;
var game_state;
var mouse_pos = [-1, -1];
function get_mouse_pos(event) {
    console.log(event.x, event.y);
    mouse_pos = [event.x, event.y];
}
document.addEventListener("DOMContentLoaded", (_) => {
    setup_canvas();
    Player.radius = canvas.width / 10;
    init_game();
    // ws()
});
function setup_canvas() {
    canvas = document.getElementById("my-canvas");
    ctx = canvas.getContext('2d');
}
class Player {
    score;
    name;
    pos;
    static radius;
    constructor(name) {
        this.score = 0;
        this.name = name;
        this.pos = [-1, -1];
    }
    move(x, y) {
        this.pos[0] = x;
        this.pos[1] = y;
    }
    draw() {
        ctx.beginPath();
        ctx.fillStyle = "red";
        ctx.arc(this.pos[0], this.pos[1], Player.radius, 0, 2 * Math.PI);
        ctx.fill();
    }
}
class GameState {
    p1;
    p2;
    puck;
    constructor(p1, p2) {
        this.p1 = new Player(p1);
        this.p2 = new Player(p2);
        this.puck = [-1, -1];
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
        this.p1.draw();
        // p2
        this.p2.draw();
    }
}
function init_game() {
    game_state = new GameState("foo", "bar");
    // put puck at middle
    game_state.puck = [canvas.width / 2, canvas.height / 2];
    // put p1 on bottom
    game_state.p1.move(canvas.width / 2, canvas.height - Player.radius);
    // put p2 on top
    game_state.p2.move(canvas.width / 2, Player.radius);
    game_state.draw();
}
