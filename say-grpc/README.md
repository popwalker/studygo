### This is a simple rpc service sample with both server and client side.

### Description
This little project was come from a youtube video series named:[justforfunc](https://youtube.com/c/justforfunc) when i learning go programming.
Thought it's useful for learn gRPC, so I write the code for recording my learning journey
Thanks [francesc](https://twitter.com/francesc)'s sharing

### Run

1.start rpc server
```shell
cd $GOPATH/src/github/popwalker/studygo/say-grpc/backend
make build
docker run --rm --name say-service -p 8080:8080 superfat/say
```
2.start client call
```shell
cd $GOPATH/src/github/popwalker/studygo/say-grpc/say
go run main.go "hello, there"
```
and then, you will get a output.wav file at the directory

### Requirements:
- Install [protobuf](https://github.com/google/protobuf/releases)
- Install [Docker](https://docs.docker.com/get-started/)

### References
- flite [docs](http://www.speech.cs.cmu.edu/flite/)
- gRPC [docs](https://grpc.io/)

