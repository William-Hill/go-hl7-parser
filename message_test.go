package hl7_test

import (
	hl7 "github.com/synkwise/go-hl7-parser"
	"os"
	"testing"
)

func readFile(fname string) ([]byte, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make([]byte, 1024)
	if _, err = file.Read(data); err != nil {
		return nil, err
	}
	return data, nil
}

func TestMessage(t *testing.T) {
	data, err := readFile("./testdata/msg.hl7")
	if err != nil {
		t.Fatal(err)
	}

	msg := &hl7.Message{Value: data}
	msg.Parse()
	if err != nil {
		t.Error(err)
	}
	if len(msg.Segments) != 5 {
		t.Errorf("Expected 5 segments got %d\n", len(msg.Segments))
	}

	data, err = readFile("./testdata/msg2.hl7")
	if err != nil {
		t.Fatal(err)
	}
	msg = &hl7.Message{Value: data}
	msg.Parse()
	if err != nil {
		t.Error(err)
	}
	if len(msg.Segments) != 5 {
		t.Errorf("Expected 5 segments got %d\n", len(msg.Segments))
	}

	data, err = readFile("./testdata/msg3.hl7")
	if err != nil {
		t.Fatal(err)
	}
	msg = &hl7.Message{Value: data}
	msg.Parse()
	if err != nil {
		t.Error(err)
	}
	if len(msg.Segments) != 9 {
		t.Errorf("Expected 9 segments got %d\n", len(msg.Segments))
	}

	data, err = readFile("./testdata/msg4.hl7")
	if err != nil {
		t.Fatal(err)
	}
	msg = &hl7.Message{Value: data}
	msg.Parse()
	if err != nil {
		t.Error(err)
	}
	if len(msg.Segments) != 9 {
		t.Errorf("Expected 9 segments got %d\n", len(msg.Segments))
	}
}

func TestMsgUnmarshal(t *testing.T) {
	fname := "./testdata/msg.hl7"
	file, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	msgs, err := hl7.NewDecoder(file).Messages()
	if err != nil {
		t.Fatal(err)
	}
	st := my7{}
	msgs[0].Unmarshal(&st)

	if st.FirstName != "John" {
		t.Errorf("Expected John got %s\n", st.FirstName)
	}
	if st.LastName != "Jones" {
		t.Errorf("Expected Jones got %s\n", st.LastName)
	}
}
