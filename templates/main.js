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
    var waspsCount = document.getElementsByClassName('wasp')[0].value;
    var hatchRate = document.getElementsByClassName('swarm')[0].value;
    var payload = {};
    payload.wasps = Number(waspsCount);
    payload.swarms = Number(hatchRate);
    let data = JSON.stringify(payload);
    fetch('http://localhost:8001/api/v1/attack', {method: 'POST', body:data});
    toggle_box()
}

