package roger

import (
	"errors"
	"testing"
)

func TestErrorPacketIsError(t *testing.T) {
	pkt := newErrorPacket(errors.New("test error"))
	if pkt.IsError() == false {
		t.Error("Test packet should return true when IsError is called")
	}
}

func TestErrorPacketResultObject(t *testing.T) {
	testError := errors.New("test error")
	pkt := newErrorPacket(testError)
	obj, err := pkt.GetResultObject()
	if err != testError {
		t.Error("An error packet should return the error when GetResultObject is called")
	}
	if obj != nil {
		t.Error("An error packet should return a nil object")
	}
}

func TestCommandFailurePacketIsError(t *testing.T) {
	failedCmdPkt := newPacket(0x01000002, []byte{})
	if failedCmdPkt.IsError() == false {
		t.Error("A command with an error flag should return true when IsError is called")
	}
}

func TestCommandFailurePacketResultObject(t *testing.T) {
	failedCmdPkt := newPacket(0x01000002, []byte{})
	obj, err := failedCmdPkt.GetResultObject()
	if err.Error() != "Command error with status: 1" {
		t.Error("A failed command packet's error message should contain the status code element of the command response")
	}
	if obj != nil {
		t.Error("A failed command packet should return a nil object")
	}
}

func TestCommandSuccessPacketIsError(t *testing.T) {
	successfulCmdPkt := newPacket(0x01000003, []byte{})
	if successfulCmdPkt.IsError() == true {
		t.Error("A command without an error flag should return false when IsError is called")
	}
}

func TestEmptyResponsePacketResultObject(t *testing.T) {
	emptyPkt := newPacket(0x01000003, []byte{})
	obj, err := emptyPkt.GetResultObject()
	if err == nil {
		t.Error("An empty packet should return an error")
	}
	if obj != nil {
		t.Error("An empty packet should return a nil object")
	}
}

func TestSuccessfulResponseResultObject(t *testing.T) {
	client, _ := NewRClient("localhost", 6311)
	pkt := client.EvaluateSync("2")
	obj, err := pkt.GetResultObject()
	if err != nil {
		t.Error("A successful query should not result in an error")
	}
	if obj == nil {
		t.Error("A successful query should return a response object")
	}
}

func TestNonSEXPResponse(t *testing.T) {
	stringResp := newPacket(0x01000003, []byte{byte(dtString)})
	obj, err := stringResp.GetResultObject()
	if err == nil {
		t.Error("Packets containing non SEXP content should return an error")
	}
	if obj != nil {
		t.Error("Packets containing non SEXP content should not return an object")
	}
}
