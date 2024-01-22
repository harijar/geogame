const WS_API = 'ws://localhost:8080/v1/ws'
const socket = new WebSocket(WS_API)
const ping_data = JSON.stringify({"type": "ping", "payload": "ws ping"})

socket.onopen = function() {
    setInterval(ping, 1000)
}

function ping() {
    socket.send(ping_data)
}