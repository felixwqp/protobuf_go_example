protoc -I src --go_out=src/simple/. src/simple/simple.proto 
protoc -I src --go_out=src/enum_example/. src/enum_example/enum_example.proto 
protoc -I src --go_out=src/complex/. src/complex/complex.proto 


# update the generate.sh for new proto file
# sh generate.sh
# git add src 
# git commit -m "new proto "
# git push
# go get -u github.com/felixwqp/protobuf_go_play