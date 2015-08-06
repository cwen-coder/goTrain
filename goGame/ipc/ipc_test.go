package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(re, rq string) *Response {
	//return "ECHO:" + request
	var t Response
	t.Body = re
	t.Code = rq
	return &t
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("From Client1", "111")
	resp2, _ := client2.Call("From Client2", "222")
	if resp1.Body != "From Client1" || resp2.Body != "From Client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:", resp2)
	}
	client1.Close()
	client2.Close()

}
