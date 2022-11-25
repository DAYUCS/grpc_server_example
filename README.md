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
docker run grpc_server_example:latest
```
