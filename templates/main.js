var start_swarm;
var swarm_button;
var item_ids =[];

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
        method = ip_array[0]
        urlPath = ip_array[1]
        sumReq = ip_array[2]
        sumFail = ip_array[3]
        avgLatency = ip_array[4]
        console.log(e)
        var stats_table = document.querySelector("#stats-body");
        item_id = urlPath.split("/").join("").trim()
        if (item_id.length == 0) {
            item_id = "root"
        }
        console.log("Got item id: ", item_id)
        exists = stats_table.querySelector("#" + item_id)
        if (typeof(exists) != 'undefined' && exists != null) {
            var tr = exists  
            tds = tr.children
            tds[0].firstChild.nodeValue = method
            tds[1].firstChild.nodeValue = urlPath
            tds[3].firstChild.nodeValue = sumFail
            tds[4].firstChild.nodeValue = avgLatency
            tds[2].firstChild.nodeValue = sumReq
        }
        else {
            var tr = document.createElement("tr");
            tr.setAttribute("id", item_id)    
            console.log("new row, tr: ", tr)
            var td_type = document.createElement("td");
            var td_name = document.createElement("td");
            var td_req = document.createElement("td");
            var td_fail = document.createElement("td");
            var td_avg_latency = document.createElement("td");
            var typeCol = document.createTextNode(method);
            var nameCol = document.createTextNode(urlPath);
            var sumReqCol = document.createTextNode(sumReq);
            var sumFailCol = document.createTextNode(sumFail);
            var avgLatencyCol = document.createTextNode(avgLatency);
            td_type.appendChild(typeCol)
            td_name.appendChild(nameCol)
            td_req.appendChild(sumReqCol)
            td_fail.appendChild(sumFailCol)
            td_avg_latency.appendChild(avgLatencyCol)
            tr.appendChild(td_type)
            tr.appendChild(td_name)
            tr.appendChild(td_req)
            tr.appendChild(td_fail)
            tr.appendChild(td_avg_latency)
            stats_table.appendChild(tr)
        }
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


function stop() {
    window.alert("stop called")
    var payload = {};
    let data = JSON.stringify(payload);
    fetch('http://139.59.59.106:8000/api/v1/stop', {method: 'POST', body:data});
}

