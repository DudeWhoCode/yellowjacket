var start_swarm;
var swarm_button;

function swarm() {
    start_swarm.style.display = 'block';
}

window.onload = function () {
   start_swarm = document.getElementById("start-swarm");
   swarm_button = document.getElementById('new-test');
};

function attack() {
    var xhttp = new XMLHttpRequest();
    var waspsCount = document.getElementsByClassName('wasp');
    var hatchRate = document.getElementsByClassName('swarm');
    var payload = {};
    payload.wasps = waspsCount;
    payload.swarms = hatchRate;
    xhttp.open("POST", "http://localhost:8000/attack", true);
    xhttp.setRequestHeader("Content-type", "application/json");
    xhttp.send(payload);
    // var response = JSON.parse(xhttp.responseText);
}

