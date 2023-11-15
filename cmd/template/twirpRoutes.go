
package template

// TwirpTemplates contains the methods used for building
// an app that uses [github.com/twitchtv/twirp]
type TwirpTemplates struct{}

func (e TwirpTemplates) Main() []byte {
	return MainTemplate()
}
func (e TwirpTemplates) Server() []byte {
	return MakeTwirpServer()
}

func (e TwirpTemplates) Routes() []byte {
	return MakeProtoFile()
}

// MakeTwirpServiceDefinition returns a byte slice that represents 
// the internal/server/
func MakeTwirpServer() []byte {
    return []byte(`

`)
}

// MakeTwirpServiceDefinition returns a byte slice that represents 
// the internal/server/
func MakeProtoFile() []byte {
    return []byte(`syntax = "proto3"

    package server

    // HelloWorld service sends a personalized "hello world" message to clients
    service HelloWorld {
    rpc SayHelloTo(ClientName) returns (PersonalizedHelloWorld);
    }

    // The name of the client receiving the message
    message ClientName {
    string name = 1;
    }

    // The message returned to the clinent
    message PersonalizedHelloWorld {
    string msg = 1;
    }

    `) // TODO
}
