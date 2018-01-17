package main

import(
	pb "github.com/popwalker/studygo/say-grpc/api"
	"google.golang.org/grpc"
	"github.com/sirupsen/logrus"
	"context"
	"fmt"
	"flag"
	"os"
	"io/ioutil"
)

func main() {
	backend:= flag.String("b", "localhost:8080", "address of the say backend")
	output := flag.String("o", "output.wav", "wav file where the output will be written")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Printf("usage:\n\t%s \"text to speak\"\n", os.Args[0])
		os.Exit(1)
	}

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not dial %s:%v", *backend, err)
	}
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)

	text := &pb.Text{Text:"hi i am client"}
	res, err := client.Say(context.Background(), text)
	if err != nil {
		logrus.Fatalf("could not say %s:%v", text, err)
	}

	if err = ioutil.WriteFile(*output, res.Audio, 0666);err != nil{
		logrus.Fatalf("could not write to %s:%v",*output, err)
	}
}
