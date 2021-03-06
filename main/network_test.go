package main

import (
	"D7024e-Kademlia/github.com/protobuf/proto"
	"net"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestNetwork_HandleReplyPing(t *testing.T) {
	id := NewKademliaID("ffffffff00000000000000000000000000000000")
	channel := make(chan interface{})
	network.CreateChannel(id, channel)
	contact := NewContact(id, "localhost:8000")

	packet := &ReplyPing{contact.ID.String(), contact.Address}
	wrapperMsg := &WrapperMessage_ReplyPing{packet}
	wrapper := &WrapperMessage{"ReplyPing", id.String(), id.String(), wrapperMsg}
	addr, _ := net.ResolveUDPAddr("udp", contact.Address)

	go network.HandleReply(wrapper, nil, addr)

	x := <-network.GetAnswerChannel(id)
	returnedContact, ok := x.(Contact)
	if ok {
		if returnedContact.ID.String() != "ffffffff00000000000000000000000000000000" {
			t.Error("Expected ffffffff00000000000000000000000000000000, got ", returnedContact.ID.String())
		}
		if returnedContact.Address != "localhost:8000" {
			t.Error("Expected localhost:8000, got " + returnedContact.Address)
		}
	} else {
		t.Error("Expected return to be of type 'Contact'")
	}
}

func TestNetwork_HandleReplyContactList(t *testing.T) {
	id := NewKademliaID("ffffffff00000000000000000000000000000000")
	channel := make(chan interface{})
	network.CreateChannel(id, channel)
	contact := NewContact(id, "localhost:8000")

	contactReply := &ReplyContactList_Contact{contact.ID.String(), contact.Address}
	contactListReply := []*ReplyContactList_Contact{}
	contactListReply = append(contactListReply, contactReply)

	packet := &ReplyContactList{contactListReply}
	wrapperMsg := &WrapperMessage_ReplyContactList{packet}
	wrapper := &WrapperMessage{"ReplyContactList", id.String(), id.String(), wrapperMsg}

	addr, _ := net.ResolveUDPAddr("udp", contact.Address)

	go network.HandleReply(wrapper, nil, addr)

	x := <-network.GetAnswerChannel(id)
	returnedContacts, ok := x.([]Contact)
	if ok {
		if returnedContacts[0].ID.String() != "ffffffff00000000000000000000000000000000" {
			t.Error("Expected ffffffff00000000000000000000000000000000, got ", returnedContacts[0].ID.String())
		}
		if returnedContacts[0].Address != "localhost:8000" {
			t.Error("Expected localhost:8000, got " + returnedContacts[0].Address)
		}
	} else {
		t.Error("Expected return to be of type '[]Contact'")
	}
}

func TestNetwork_HandleReplyData(t *testing.T) {
	id := NewKademliaID("ffffffff00000000000000000000000000000000")
	channel := make(chan interface{})
	network.CreateChannel(id, channel)
	contact := NewContact(id, "localhost:8000")

	packet := &ReplyData{""}
	wrapperMsg := &WrapperMessage_ReplyData{packet}
	wrapper := &WrapperMessage{"ReplyData", id.String(), id.String(), wrapperMsg}

	addr, _ := net.ResolveUDPAddr("udp", contact.Address)

	go network.HandleReply(wrapper, nil, addr)

	x := <-network.GetAnswerChannel(id)
	returnedAdress, ok := x.(string)
	if ok {
		if returnedAdress != "127.0.0.1:8000" {
			t.Error("Expected 127.0.0.1:8000, got ", returnedAdress)
		}
	} else {
		t.Error("Expected return to be of type 'String'")
	}
}

func TestNetwork_HandleReplyStore(t *testing.T) {
	id := NewKademliaID("ffffffff00000000000000000000000000000000")
	channel := make(chan interface{})
	network.CreateChannel(id, channel)
	contact := NewContact(id, "localhost:8000")

	packet := &ReplyStore{"ok"}
	wrapperMsg := &WrapperMessage_ReplyStore{packet}
	wrapper := &WrapperMessage{"ReplyStore", id.String(), id.String(), wrapperMsg}

	addr, _ := net.ResolveUDPAddr("udp", contact.Address)

	go network.HandleReply(wrapper, nil, addr)

	x := <-network.GetAnswerChannel(id)
	reply, ok := x.(string)
	if ok {
		if reply != "ok" {
			t.Error("Expected ok, got ", reply)
		}
	} else {
		t.Error("Expected return to be of type 'String'")
	}
}

func TestNetwork_RepublishData(t *testing.T) {
	storeKademlia := NewKademlia(network)
	go storeKademlia.Store("testStore.txt")
	time.Sleep(time.Duration(1) * time.Second)
	port := 9100
	a := "localhost:" + strconv.Itoa(port)

	ID := HashKademliaID("testStore.txt")
	rt := NewRoutingTable(NewContact(ID, a))
	//nodeList = append(nodeList, rt)
	rt.AddContact(network.node.rt.me)
	node := NewNode(rt)
	tcpNetwork := NewFileNetwork(node, "localhost", port)
	nw := NewNetwork(node, tcpNetwork, "localhost", port)
	//fmt.Println("Ny Nod varv " + strconv.Itoa(i+1) + ": " + rt.me.String())
	//go nw.Listen("localhost", port)
	kademlia := NewKademlia(nw)

	contactResult, _ := kademlia.LookupContact(ID, false)
	if len(contactResult) > 0 {
		for q := range contactResult {
			rt.AddContact(contactResult[q])
		}
	}

	if _, err := os.Stat("kademliastorage/" + ID.String()); os.IsNotExist(err) {
		os.Mkdir("kademliastorage/"+ID.String(), 0777)
	}

	if _, err := os.Stat("upload/" + ID.String()); os.IsNotExist(err) {
		os.Mkdir("upload/"+ID.String(), 0777)
	}

	if _, err := os.Stat("downloads/" + ID.String()); os.IsNotExist(err) {
		os.Mkdir("downloads/"+ID.String(), 0777)
	}
	go network.RepublishData()
	time.Sleep(20000 * time.Millisecond)

	gotData := node.GotData(*HashKademliaID("testStore.txt"))

	if !gotData {
		t.Error("Expected node to have data")
	}
}

func TestNetwork_HandleRequestPing(t *testing.T) {
	id := NewKademliaID("ffffffff00000000000000000000000000000000")
	packet := &RequestPing{}
	wrapperMsg := &WrapperMessage_RequestPing{packet}
	wrapper := &WrapperMessage{"RequestPing", id.String(), id.String(), wrapperMsg}

	contact := NewContact(id, "localhost:9000")
	serverAddr, err := net.ResolveUDPAddr("udp", contact.Address)
	CheckError(err)
	serverConn, err := net.ListenUDP("udp", serverAddr)
	CheckError(err)
	defer serverConn.Close()
	buf := make([]byte, 4096)

	go network.HandleRequest(wrapper, nil, serverAddr)

	for {
		n, _, _ := serverConn.ReadFromUDP(buf)
		message := &WrapperMessage{}
		_ = proto.Unmarshal(buf[0:n], message)
		if message.ID[0:9] != "ReplyPing" {
			t.Error("Expected message id 'ReplyPing', got " + message.ID)
		}
		if message.SourceID != network.node.rt.me.ID.String() {
			t.Error("Expected message source to be " + network.node.rt.me.ID.String() + ", got " + message.SourceID)
		}
		return
	}
}

func TestNetwork_HandleRequestContact(t *testing.T) {
	id := NewKademliaID("ffffffff00000000000000000000000000000000")
	packet := &RequestContact{NewRandomKademliaID().String()}
	wrapperMsg := &WrapperMessage_RequestContact{packet}
	wrapper := &WrapperMessage{"RequestContact", id.String(), id.String(), wrapperMsg}

	contact := NewContact(id, "localhost:9000")
	serverAddr, err := net.ResolveUDPAddr("udp", contact.Address)
	CheckError(err)
	serverConn, err := net.ListenUDP("udp", serverAddr)
	CheckError(err)
	defer serverConn.Close()
	buf := make([]byte, 4096)

	go network.HandleRequest(wrapper, nil, serverAddr)

	for {
		n, _, _ := serverConn.ReadFromUDP(buf)
		message := &WrapperMessage{}
		_ = proto.Unmarshal(buf[0:n], message)
		if message.ID[0:16] != "ReplyContactList" {
			t.Error("Expected message id 'ReplyContactList', got " + message.ID)
		}
		if message.SourceID != network.node.rt.me.ID.String() {
			t.Error("Expected message source to be " + network.node.rt.me.ID.String() + ", got " + message.SourceID)
		}
		return
	}
}

func TestNetwork_HandleRequestData(t *testing.T) {
	id := NewKademliaID("ffffffff00000000000000000000000000000000")
	packet := &RequestData{"ffffffff00000000000000000000000000000000"}
	wrapperMsg := &WrapperMessage_RequestData{packet}
	wrapper := &WrapperMessage{"RequestData", id.String(), id.String(), wrapperMsg}

	contact := NewContact(id, "localhost:9000")
	serverAddr, err := net.ResolveUDPAddr("udp", contact.Address)
	CheckError(err)
	serverConn, err := net.ListenUDP("udp", serverAddr)
	CheckError(err)
	defer serverConn.Close()
	buf := make([]byte, 4096)

	go network.HandleRequest(wrapper, nil, serverAddr)

	for {
		n, _, _ := serverConn.ReadFromUDP(buf)
		message := &WrapperMessage{}
		_ = proto.Unmarshal(buf[0:n], message)
		if message.ID[0:16] != "ReplyContactList" {
			t.Error("Expected message id 'ReplyContactList', got " + message.ID)
		}
		if message.SourceID != network.node.rt.me.ID.String() {
			t.Error("Expected message source to be " + network.node.rt.me.ID.String() + ", got " + message.SourceID)
		}
		return
	}
}

/*func TestNetwork_HandleRequestStore(t *testing.T) {
	id := NewKademliaID("ffffffff00000000000000000000000000000000")
	packet := &RequestStore{NewRandomKademliaID().String(), "data"}
	wrapperMsg := &WrapperMessage_RequestStore{packet}
	wrapper := &WrapperMessage{"RequestStore", id.String(), id.String(), wrapperMsg}

	contact := NewContact(id, "localhost:9000")
	serverAddr, err := net.ResolveUDPAddr("udp", contact.Address)
	CheckError(err)
	serverConn, err := net.ListenUDP("udp", serverAddr)
	CheckError(err)
	defer serverConn.Close()
	buf := make([]byte, 4096)

	go network.HandleRequest(wrapper, nil, serverAddr)

	for {
		n, _, _ := serverConn.ReadFromUDP(buf)
		message := &WrapperMessage{}
		_ = proto.Unmarshal(buf[0:n], message)
		if message.ID[0:10] != "ReplyStore" {
			t.Error("Expected message id 'ReplyStore', got " + message.ID)
		}
		if message.SourceID != network.node.rt.me.ID.String() {
			t.Error("Expected message source to be " + network.node.rt.me.ID.String() + ", got " + message.SourceID)
		}
		return
	}
}*/

func TestNetwork_SendPingMessage(t *testing.T) {
	id := NewKademliaID("ffffffff00000000000000000000000000000000")
	channel := make(chan interface{})
	contact := NewContact(id, "localhost:9000")
	go network.SendPingMessage(contact, channel)

	x := <-channel
	reply, ok := x.(bool)
	if ok {
		if reply != false {
			t.Error("Expected false, got ", reply)
		}
	} else {
		t.Error("Expected reply to be of type 'bool'")
	}
}

func TestNetwork_SendPingMessage2(t *testing.T) {
	id := NewKademliaID("ffffffff00000000000000000000000000000000")
	channel := make(chan interface{})
	contact := NewContact(id, "localhost:9000")
	serverAddr, err := net.ResolveUDPAddr("udp", contact.Address)
	CheckError(err)
	serverConn, err := net.ListenUDP("udp", serverAddr)
	CheckError(err)
	defer serverConn.Close()
	buf := make([]byte, 4096)

	go network.SendPingMessage(contact, channel)
	for {
		n, _, _ := serverConn.ReadFromUDP(buf)
		message := &WrapperMessage{}
		_ = proto.Unmarshal(buf[0:n], message)
		if message.ID[0:11] != "RequestPing" {
			t.Error("Expected message id 'RequestPing', got " + message.ID)
		}
		if message.SourceID != network.node.rt.me.ID.String() {
			t.Error("Expected message source to be " + network.node.rt.me.ID.String() + ", got " + message.SourceID)
		}
		return
	}

}

func TestNetwork_SendFindContactMessage(t *testing.T) {
	id := NewKademliaID("ffffffff00000000000000000000000000000000")
	channel := make(chan interface{})
	contact := NewContact(id, "localhost:9000")
	serverAddr, err := net.ResolveUDPAddr("udp", contact.Address)
	CheckError(err)
	serverConn, err := net.ListenUDP("udp", serverAddr)
	CheckError(err)
	defer serverConn.Close()
	buf := make([]byte, 4096)

	go network.SendFindContactMessage(id, &contact, channel)
	for {
		n, _, _ := serverConn.ReadFromUDP(buf)
		message := &WrapperMessage{}
		_ = proto.Unmarshal(buf[0:n], message)
		if message.ID[0:14] != "RequestContact" {
			t.Error("Expected message id 'RequestContact', got " + message.ID)
		}
		if message.SourceID != network.node.rt.me.ID.String() {
			t.Error("Expected message source to be " + network.node.rt.me.ID.String() + ", got " + message.SourceID)
		}
		return
	}
}

func TestNetwork_SendFindDataMessage(t *testing.T) {
	id := NewKademliaID("ffffffff00000000000000000000000000000000")
	channel := make(chan interface{})
	contact := NewContact(id, "localhost:9000")
	serverAddr, err := net.ResolveUDPAddr("udp", contact.Address)
	CheckError(err)
	serverConn, err := net.ListenUDP("udp", serverAddr)
	CheckError(err)
	defer serverConn.Close()
	buf := make([]byte, 4096)

	go network.SendFindDataMessage(id.String(), &contact, channel)
	for {
		n, _, _ := serverConn.ReadFromUDP(buf)
		message := &WrapperMessage{}
		_ = proto.Unmarshal(buf[0:n], message)
		if message.ID[0:11] != "RequestData" {
			t.Error("Expected message id 'RequestData', got " + message.ID)
		}
		if message.SourceID != network.node.rt.me.ID.String() {
			t.Error("Expected message source to be " + network.node.rt.me.ID.String() + ", got " + message.SourceID)
		}
		return
	}
}

func TestNetwork_SendStoreMessage(t *testing.T) {
	id := NewKademliaID("ffffffff00000000000000000000000000000000")
	channel := make(chan interface{})
	contact := NewContact(id, "localhost:9000")
	serverAddr, err := net.ResolveUDPAddr("udp", contact.Address)
	CheckError(err)
	serverConn, err := net.ListenUDP("udp", serverAddr)
	CheckError(err)
	defer serverConn.Close()
	buf := make([]byte, 4096)

	go network.SendStoreMessage(id.String(), contact.Address, channel)
	for {
		n, _, _ := serverConn.ReadFromUDP(buf)
		message := &WrapperMessage{}
		_ = proto.Unmarshal(buf[0:n], message)
		if message.ID[0:12] != "RequestStore" {
			t.Error("Expected message id 'RequestStore', got " + message.ID)
		}
		if message.SourceID != network.node.rt.me.ID.String() {
			t.Error("Expected message source to be " + network.node.rt.me.ID.String() + ", got " + message.SourceID)
		}
		return
	}
}

func TestNetwork_TimeoutWaiter(t *testing.T) {
	id := NewKademliaID("ffffffff00000000000000000000000000000000")
	channel := make(chan interface{})
	network.CreateChannel(id, channel)
	returnChannel := make(chan interface{})

	go network.TimeoutWaiter(0, returnChannel, id)

	x := <-returnChannel
	reply, ok := x.(bool)
	if ok {
		if reply != false {
			t.Error("Expected false, got ", reply)
		}
		if network.waitingAnswerList[*id] != nil {
			t.Error("Expected network.waitingAnswerList[*id] to be nil")
		}
	} else {
		t.Error("Expected reply to be of type 'bool'")
	}
}
