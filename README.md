#### Dependencies
golang v1.19.3+

# run

```shell
cd path_to_grpc_server_example
(change crtFilePath & keyFilePath before run)
go run main.go
```

# run in docker

```shell
docker build -t grpc_server_example:latest .
docker run -p 50051:50051 -p 50052:50052 -p 50053:50053 grpc_server_example:latest
```

# build client docker image
```shell
cd client
docker build -t simple-http-server:latest .
```