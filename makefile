proto_build:
	# 根据proto文件编译生成对应的go文件
	protoc -I ./api -I ./third_party/googleapis \
     -I ./third_party/github.com \
	--openapiv2_out ./api --openapiv2_opt logtostderr=true \
	--openapiv2_opt json_names_for_fields=true  \
	--openapiv2_opt allow_delete_body=true \
	--gofast_out=plugins=grpc,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:. \
	--go-gin_out ./api --go-gin_opt=paths=source_relative \
	./api/v1/*.proto

mockgen:
	cd ./internal && for file in `egrep -rnl "type.*?interface" ./domain | grep -v "_test" `; do \
		echo $$file ; \
		cd .. && mockgen -destination="./internal/adapters/mock/$$file" -source="./internal/$$file" && cd ./internal ; \
	done