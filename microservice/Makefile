gokit:
	mkdir -p protogen
	protoc --go_out=plugins=grpc:protogen proto/*.proto
	protoc -I. --gotemplate_out=debug=true,destination_dir=hello/gen:protogen proto/*.proto

micro:
	