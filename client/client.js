const { HelloRequest, Person } = require('./helloworld_pb.js');
const { GreeterClient } = require('./helloworld_grpc_web_pb.js');

//Enovy grpc-web
var client = new GreeterClient('http://10.39.107.102:8000', null, null);

//Make an unary RPC call
var request = new HelloRequest();
request.setName("Alex");
request.setItemsList(["1st item", "2nd item"]);
request.setGender(2);
person = new Person();
person.setName("Alex");
person.setAge(20);
request.setPerson(person);

client.sayHello(request, {}, (err, response) => {
    if (err) {
        console.log(`Unexpected error for sayHello: code = ${err.code}` +
            `, message = "${err.message}"`);
    } else {
        console.log("Response from an Unary gRPC call:");
        console.log(response.toObject());
    }
});

//Server stream RPC call
var stream = client.sayHelloServerStream(request, {});
stream.on('data', (response) => {
    console.log("Response from a Server Stream gRPC call:");
    console.log(response.toObject());
});
stream.on('error', (err) => {
    console.log(`Unexpected stream error: code = ${err.code}` +
        `, message = "${err.message}"`);
});
stream.on('end', function (end) {
    // stream end signal
});

// to close the stream
//stream.cancel();