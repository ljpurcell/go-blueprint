package template

// TwirpTemplates contains the methods used for building
// an app that uses [github.com/twitchtv/twirp]
type TwirpTemplates struct{}

func (t TwirpTemplates) Main() []byte {
	return MakeTwirpMain()
}
func (t TwirpTemplates) Server() []byte {
	return MakeTwirpServer()
}

func (t TwirpTemplates) Routes() []byte {
	return MakeHTTPRoutes()
}

func (t TwirpTemplates) ProtoFile() []byte {
	return MakeProtoFile()
}

// MakeTwirpServiceDefinition returns a byte slice that represents
// the internal/server/
func MakeProtoFile() []byte {
	return []byte(`syntax = "proto3";

    option go_package = "{{.ProjectName}}/twirp/hello";

    // HelloWorld service sends a personalized "hello world" message to clients
    service HelloWorld {
    rpc SayHelloTo(ClientName) returns (PersonalizedHelloWorld);
    }

    // The name of the client receiving the message
    message ClientName {
    string name = 1;
    }

    // The message returned to the client
    message PersonalizedHelloWorld {
    string msg = 1;
    }

    `)
}

// MakeTwirpMain returns a byte slice that represents
// the main.go file for the Twirp implementation.
func MakeTwirpMain() []byte {
	return []byte(`
    package main

    import (
        "net/http"
        "github.com/twitchtv/twirp"
		pb "{{.ProjectName}}/twirp/hello"
    )

    func main() {	
        hwServer := pb.NewHelloWorldServer(&handlers.HelloWorld{})
        mux := http.NewServeMux()
        mux.Handle(hwServer.PathPrefix(), hwServer)

        // Start the server
        http.ListenAndServe(":8080", hwServer)
    }
    `)
}

// MakeTwirpServiceDefinition returns a byte slice that represents
// the internal/server/
func MakeTwirpServer() []byte {
	return []byte(`
    package server

    import (
        "context"
        "github.com/twitchtv/twirp"
        "{{.ProjectName}}/twirp/hello"
    )

    type HelloWorld struct{}

    func (hw *HelloWorld) HelloWorld(ctx context.Context, req *protobuf.HelloRequest) (*protobuf.HelloResponse, error) {
        return &protobuf.HelloResponse{Message: "Hello, Twirp!"}, nil
    }
    `)
}
