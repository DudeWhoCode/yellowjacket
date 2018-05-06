var start_swarm;
var swarm_button;

function close_btn() {
    console.log('close call');
    start_swarm.style.display = 'none'
}

function toggle_box() {
    if(start_swarm.style.display === 'block') {
        console.log('Setting display to none');
        start_swarm.style.display = 'none'
    }
    else{
        start_swarm.style.display = 'block';
    }

}


window.onload = function () {
   start_swarm = document.getElementById("start-swarm");
   swarm_button = document.getElementById('new-test');
};

function attack() {
    var xhttp = new XMLHttpRequest();
    var waspsCount = document.getElementsByClassName('wasp')[0].value;
    var hatchRate = document.getElementsByClassName('swarm')[0].value;
    var payload = {};
    payload.wasps = Number(waspsCount);
    payload.swarms = Number(hatchRate);
    console.log(
        'Payload: ', payload
    );
    xhttp.open("POST", "http://localhost:8000/attack", true);
    xhttp.setRequestHeader("Content-type", "application/json");
    xhttp.send(JSON.stringify(payload));
    console.log('Before toggling');
    toggle_box();
    console.log('After toggling')
    // var response = JSON.parse(xhttp.responseText);
}

