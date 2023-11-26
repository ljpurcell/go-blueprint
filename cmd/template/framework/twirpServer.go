package framework

import (
	_ "embed"
)

//go:embed files/server/twirp_server.go.tmpl
var twirpServerTemplate []byte

//go:embed files/main/twirp_main.go.tmpl
var twirpMainTemplate []byte

//go:embed files/protoFile/twirp_hello.proto.tmpl
var twirpProtoFileTemplate []byte

// TwirpTemplates contains the methods used for building
// an app that uses [github.com/twitchtv/twirp]
type TwirpTemplates struct{}

func (t TwirpTemplates) Routes() []byte {
	return nil
}

func (t TwirpTemplates) RoutesWithDB() []byte {
	return nil
}

func (t TwirpTemplates) ServerWithDB() []byte {
	return nil
}

func (t TwirpTemplates) TestHandler() []byte {
	return nil
}

func (t TwirpTemplates) Main() []byte {
	return twirpMainTemplate
}

func (t TwirpTemplates) Server() []byte {
	return twirpServerTemplate
}

func (t TwirpTemplates) ProtoFile() []byte {
	return twirpProtoFileTemplate
}
