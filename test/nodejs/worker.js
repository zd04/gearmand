var gearmanode = require('gearmanode');
var worker = gearmanode.worker({servers: [{host: '192.168.100.211'}, {port: 4731}]});
 
worker.addFunction('reverse', function (job) {
    job.sendWorkData(job.payload); // mirror input as partial result
    job.workComplete(job.payload.toString().split("").reverse().join(""));
});