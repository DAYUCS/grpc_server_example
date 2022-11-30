const { HelloRequest, Person } = require('./helloworld_pb.js');
const { GreeterClient } = require('./helloworld_grpc_web_pb.js');

//Enovy grpc-web
var client = new GreeterClient('http://10.39.107.102:8080', null, null);

//Make an unary RPC call
var request = new HelloRequest();
request.setName("Alex");
request.setItemsList(["1st item", "2nd item"]);
request.setGender(2);
var person = new Person({ name: "Alex", age: 20 });
request.setPerson(person);

client.sayHello(request, {}, (err, response) => {
    if (err) {
        console.log(`Unexpected error for sayHello: code = ${err.code}` +
            `, message = "${err.message}"`);
    } else {
        console.log(response.getMessage());
    }
});

//Server stream RPC call
var stream = client.sayHelloServerStream(request, {});
stream.on('data', (response) => {
    console.log(response.getMessage());
});
stream.on('error', (err) => {
    console.log(`Unexpected stream error: code = ${err.code}` +
        `, message = "${err.message}"`);
});