// Code generated by protoc-gen-goclay, but your can (must) modify it.
// source: pb/strings.proto

package strings

import (
	desc "github.com/utrack/clay/integration/binding_with_body_and_response/pb"
	"github.com/utrack/clay/v2/transport"
)

type StringsImplementation struct{}

func NewStrings() *StringsImplementation {
	return &StringsImplementation{}
}

// GetDescription is a simple alias to the ServiceDesc constructor.
// It makes it possible to register the service implementation @ the server.
func (i *StringsImplementation) GetDescription() transport.ServiceDesc {
	return desc.NewStringsServiceDesc(i)
}
