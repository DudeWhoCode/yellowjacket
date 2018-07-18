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

function show_charts() {
    charts.style.display = 'block';
    stats.style.display = 'none'
    failures.style.display = 'none';
    exceptions.style.display = 'none';
    downloads.style.display = 'none';
}

function show_stats() {
    stats.style.display = 'block'
    charts.style.display = 'none';
    failures.style.display = 'none';
    exceptions.style.display = 'none';
    downloads.style.display = 'none';
    
}

function show_failures() {
    failures.style.display = 'block';
    stats.style.display = 'none';
    charts.style.display = 'none';
    exceptions.style.display = 'none';
    downloads.style.display = 'none';
}

function show_exceptions() {
    exceptions.style.display = 'block';
    failures.style.display = 'none';
    stats.style.display = 'none';
    charts.style.display = 'none';
    downloads.style.display = 'none';
}

function show_downloads() {
    downloads.style.display = 'block';
    exceptions.style.display = 'none';
    failures.style.display = 'none';
    stats.style.display = 'none';
    charts.style.display = 'none';
}


function handle_sse() {
    // Create a new HTML5 EventSource
    var source = new EventSource('/events/');
    // Create a callback for when a new message is received.
    source.onmessage = function(e) {
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
   charts = document.getElementsByClassName('charts-console')[0]
   stats = document.getElementsByClassName('stats-console')[0]
   failures = document.getElementsByClassName('failures-console')[0]
   exceptions = document.getElementsByClassName('exceptions-console')[0]
   downloads = document.getElementsByClassName('downloads-console')[0]
   handle_sse()
};

function attack() {
    var waspsCount = document.getElementsByClassName('wasp')[0].value;
    var hatchRate = document.getElementsByClassName('swarm')[0].value;
    var payload = {};
    payload.wasps = Number(waspsCount);
    payload.hatch_rate = Number(hatchRate);
    let data = JSON.stringify(payload);
    fetch('http://139.59.59.106:8000/api/v1/attack', {method: 'POST', body:data});
    toggle_box()
}

