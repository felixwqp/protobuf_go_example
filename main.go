package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"simple/simplepb"

	"github.com/golang/protobuf/proto"
	// proto "github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("Ritten", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Cannot serialise to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Cannot write to file", err)
		return err
	}
	fmt.Println("Data Written")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("something wrong file reading file", err)
		return err
	}
	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Cannot put the bytes into protocol buffers struct", err2)
		return err2
	}
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Msg",
		SampleList: []int32{1, 4, 5, 7},
	}
	fmt.Println(sm)
	sm.Name = "Renamed"
	fmt.Println(sm)
	fmt.Println(sm.GetId())
	return &sm
}
