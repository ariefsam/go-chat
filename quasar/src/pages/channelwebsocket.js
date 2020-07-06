import ReconnectingWebSocket from "reconnecting-websocket";

var connection;
function connect(address) {
    if (typeof connection != "undefined") {
        connection.close()
    }
    connection = new ReconnectingWebSocket(
        address
    );
    return connection
}

export { connect }