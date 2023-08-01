package fkdevice

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/robinpowered/go-proto/message"
	"github.com/robinpowered/go-proto/stream"

	pbapp "github.com/fieldkit/app-protocol"
	pbdata "github.com/fieldkit/data-protocol"
)

type DecodeFunction func(ctx context.Context, b []byte) error

func DecodeApp(ctx context.Context, b []byte) error {
	var reply pbapp.HttpReply
	err := proto.Unmarshal(b, &reply)
	if err != nil {
		return err
	}

	replyJson, err := json.MarshalIndent(reply, "", "  ")
	if err == nil {
		log.Printf("%s", replyJson)
	}

	return nil
}

func DecodeModuleConfig(ctx context.Context, b []byte) error {
	var reply pbdata.ModuleConfiguration
	err := proto.Unmarshal(b, &reply)
	if err != nil {
		return err
	}

	replyJson, err := json.MarshalIndent(reply, "", "  ")
	if err == nil {
		log.Printf("%s", replyJson)
	}

	return nil
}

func DecodeData(ctx context.Context, b []byte) error {
	var reply pbdata.DataRecord
	err := proto.Unmarshal(b, &reply)
	if err != nil {
		return err
	}

	replyJson, err := json.MarshalIndent(reply, "", "  ")
	if err == nil {
		log.Printf("%s", replyJson)
	}

	return nil
}

func Decode(ctx context.Context, decode DecodeFunction) error {
	unmarshalFunc := message.UnmarshalFunc(func(b []byte) (proto.Message, error) {
		return nil, decode(ctx, b)
	})

	_, err := stream.ReadLengthPrefixedCollection(os.Stdin, unmarshalFunc)
	if err != nil {
		return err
	}

	return nil
}
