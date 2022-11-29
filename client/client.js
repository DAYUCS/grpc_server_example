const { HelloRequest, Person } = require('./helloworld_pb.js');
const { GreeterClient } = require('./helloworld_grpc_web_pb.js');

XMLHttpRequest = require("xmlhttprequest").XMLHttpRequest;

//Apisix grpc-web
var client = new GreeterClient("http://10.39.101.186:9080/grpc/web");

//Make an unary RPC call
var request = new HelloRequest();
request.setName("Alex");
request.setItemsList(["1st item", "2nd item"]);
request.setGender(2);
var person = new Person({ name: "Alex", age: 20 });
request.setPerson(person);
client.sayHello(request, {}, (err, respose) => {
    if (err) {
        console.log(err.code);
        console.log(err.message);
    } else {
        console.log(response.getMessage());
    }
});