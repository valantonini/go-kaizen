console.log("Connecting to server ...");

let ws:WebSocket;

try {
    ws=new WebSocket('ws://localhost:5024/ws');
} catch(err) {
    console.log('Failed to connect to server ... exiting', err);
    // @ts-ignore
    Deno.exit(1);
}
ws.onopen=connected;
ws.onmessage=m=>processMessage(ws, m);
ws.onclose=disconnected;

const delay = (ms: number) => {
    // @ts-ignore
    return new Promise(resolve => setTimeout(resolve, ms));
}
// @ts-ignore
await delay(10000) /// waiting 1 second.

function connected() {
    console.log('Connected to server ...');
}
function disconnected() {
    console.log('Disconnected from server ...');
}

let count = 10;
function processMessage(ws:WebSocket, m:MessageEvent) {
    console.log('SERVER >> '+m.data);
}