var ws = new WebSocket('ws://localhost:8080/echo');
ws.onmessage = function (msg) {
    var li = createLIFromWSMessage(msg);
    var logOutlet = document.querySelector('.log-outlet')
    logOutlet.insertBefore(li, logOutlet.firstChild)
}

function sendQuery(query) {
    if (_isWSOpenAndReady()) {
        var query = document.querySelector('.query-input').value;
        ws.send(query)
    }
    else {
        setTimeout(function () {
            sendQuery();
        }, 200);
    }
}

function createLIFromWSMessage(msg) {
    var li = document.createElement('li');
    li.style.listStyle = "none";
    li.style.fontSize = 'larger';
    li.innerHTML = (msg.data.replace(/ERROR/, "<b style='color: red'>ERROR</b>") || "");
    return li;
}

function _isWSOpenAndReady() {
    return 1 === ws.readyState;
}
