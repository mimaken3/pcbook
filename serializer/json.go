package serializer

import (
	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// ProtobufToJSON converts protocol buffer message to JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	// There's a couple of things that we can config
	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false, // Such as write enums as integers or strings
		EmitDefaults: true,  // Write fields with default value or not
		Indent:       "  ",  // What's the indentation we want to use
		OrigName:     true,  // Do we want to use the original field name defined in the proto file
	}

	return marshaler.MarshalToString(message)
}
