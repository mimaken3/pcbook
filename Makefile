gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

clean:
	rm pb/*.go

run: 
	go run main.go

# cover flag: measure the code coverage of our test
# race flag:  detect any racing condition in our codes
test: 
	go test -cover -race ./...
