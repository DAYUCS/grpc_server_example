const { Input, Output } = require('./helloworld_pb.js');
const { HelloworldServiceClient } = require('./helloworld_grpc_web_pb.js');

//Apisix grpc-web
var client = new HelloworldServiceClient("http://10.39.101.186:9080/grpc/web");

//Make an unary RPC call
var request = new HelloRequest();
request.setName("Alex");
request.setItemsList(["1st item", "2nd item"]);
request.setGender(GENDER_MALE);
request.setPerson({ name: Alex, age: 20 });
client.sayHello(request, {}, (err, respose) => {
    console.log(response.getMessage());
});