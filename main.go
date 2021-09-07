package main

import (
	"fmt"
	"io/ioutil"
	"log"

	complexpb "github.com/felixwqp/protobuf_go_play/src/complex"
	"github.com/felixwqp/protobuf_go_play/src/enum_example"
	simplepb "github.com/felixwqp/protobuf_go_play/src/simple"

	"github.com/golang/protobuf/proto"
	// proto "github.com/golang/protobuf/proto"
)

func main() {
	ReadWriteDemo()
	doEnum()
	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "FirstMsg",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   12,
				Name: "Second",
			},
			&complexpb.DummyMessage{
				Id:   144,
				Name: "aaSecond",
			},
		},
	}
	fmt.Println(cm)
}

func doEnum() {
	ep := enum_example.EnumMessage{
		Id:           42,
		DayOfTheWeek: enum_example.DayOfTheWeek_MONDAY,
	}
	ep.DayOfTheWeek = enum_example.DayOfTheWeek_SATURDAY
	fmt.Println(ep)

}

func ReadWriteDemo() {
	sm := doSimple()
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("Written", sm2)
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
