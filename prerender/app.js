const prerender = require('prerender');
var server = prerender({
    waitAfterLastRequest: 500
});
server.start();