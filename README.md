# go-0x002

## Unary

1. Generate proto file under `go-0x002/unary/proto`

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative schema.proto
```
