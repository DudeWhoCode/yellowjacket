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

function handle_sse() {
    // Create a new HTML5 EventSource
    var source = new EventSource('/events/');
    // Create a callback for when a new message is received.
    source.onmessage = function(e) {

        // Append the `data` attribute of the message to the DOM.
        // document.body.innerHTML += e.data + '<br>';
        var ip_array = e.data.split(',');
        sumReq = ip_array[0]
        sumFail = ip_array[1]
        avgLatency = ip_array[2]
        console.log(e)
        document.getElementById("sum-success").innerHTML = sumReq
        document.getElementById("sum-fail").innerHTML = sumFail
        document.getElementById("avg-latency").innerHTML = avgLatency
    };
}

window.onload = function () {
   start_swarm = document.getElementById("start-swarm");
   swarm_button = document.getElementById('new-test');
   handle_sse()
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

