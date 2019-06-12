
var url = "ws://"+window.location.host+"/ws";

var ws = new WebSocket(url);

var containerData = document.getElementById("containerData");

ws.onmessage = function (msg) {
    containerData.innerHTML+=`<pre>${msg.data}</pre>`;
}