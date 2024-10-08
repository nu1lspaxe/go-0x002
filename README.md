# go-0x002

## Service Definition

1. Unary
   The client sends a single requrest and gets a single response back.
   ```proto
   rpc SayHello(HelloRequest) returns (HelloResponse);
   ```

2. Server streaming
   The client sends a request and gets a stream to read a sequence of messages back.
   ```proto
   rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);
   ```

3. Client streaming
   The client writes a sequence of messages and waits for the server to read them and return its response.
   ```proto
   rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);
   ```

4. Bidirectional streaming
   Both sides send a sequence of messages using a read-write stream. Clients and servers can read and write in whatever order they like.
   ```proto
   rpc BidiHello(stream HelloRequest) returns (stream HelloResponse);
   ```

### Scripts
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative schema.proto
```

## Unary

Generate proto file under `go-0x002/unary/proto`, run the [script](#scripts).

## Streaming Server

Generate proto file under `go-0x002/streaming-server/proto`, run the [script](#scripts).

## Streaming Client

Generate proto file under `go-0x002/streaming-client/proto`, run the [script](#scripts).

## Streaming Bidirectional

Generate proto file under `go-0x002/streaming-bidirectional/proto`, run the [script](#scripts).
