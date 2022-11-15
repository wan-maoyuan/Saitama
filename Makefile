# errors protoc 文件夹
ERROR_PROTO_FILES=$(shell find errors -name *.proto)


.PHONY: errors
errors:
	protoc --proto_path=. 														\
 	       --go_out=paths=source_relative:. 									\
 	       --go-grpc_out=paths=source_relative:. 								\
		   --go-errors_out=paths=source_relative:. 								\
	       $(ERROR_PROTO_FILES)