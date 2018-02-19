GRPC API for Mail service:
======================

Requirement:
----------

[GRPC](https://grpc.io/)

[Installation instruction](https://grpc.io/docs/quickstart/go.html#install-grpc)

Generate golang package:
```bash
protoc --proto_path=./ --go_out=./ ./api.proto
```
