package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// Since the data is written in binary format, we cannot read it.
// So let's write another function to serialize it to JSON format instead.
// In this func, we must convert the protobuf message into a JSON string first.
func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	// After calling ProtobufToJSON, we got the JSON string
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}

	// All we need to do is to write that string to the file
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON data to file: %w", err)
	}

	return nil
}

// WriteProtobufToBinaryFile writes protocol buffer message to binary file
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	// Serialize the message to binary
	data, err := proto.Marshal(message)

	if err != nil {
		// if an error occurs, Just wrap it and return to the caller
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}

	// Save the data to the spesified filename
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary dta to file: %w", err)
	}

	return nil
}

// ReadProtobufFromBinaryFile reads protocol buffer message from binary file
func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}

	// Deserialize the binary data into a protobuf message
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message: %w", err)
	}

	return nil
}
