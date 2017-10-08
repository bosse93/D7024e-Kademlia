package main

import (
	"net"
	"log"
	"D7024e-Kademlia/github.com/protobuf/proto"
	"fmt"
	"strconv"
	"time"
	"sync"
)

type Network struct {
	node *Node
	waitingAnswerList map[KademliaID](chan interface{})
	listenConnection *net.UDPConn
	mux *sync.Mutex
	rtMux *sync.Mutex
	timeoutTime int
	fileNetwork *FileNetwork
	republishSleepTimeSeconds int 
}

func NewNetwork(node *Node, fileNetwork *FileNetwork, ip string, port int) *Network {
	network := &Network{}
	network.node = node
	network.waitingAnswerList = make(map[KademliaID]chan interface{})
	network.mux = &sync.Mutex{}
	network.rtMux = &sync.Mutex{}
	network.timeoutTime = 5
	network.fileNetwork = fileNetwork
	network.republishSleepTimeSeconds = 10

	// ESTABLISH UDP CONNECTION
	serverAddr, err := net.ResolveUDPAddr("udp", ip + ":" + strconv.Itoa(port))
	CheckError(err)

	serverConn, err := net.ListenUDP("udp", serverAddr)
	CheckError(err)
	network.listenConnection = serverConn
	buf := make([]byte, 4096)
	fmt.Println("Listening on port " + strconv.Itoa(port))
	go network.Listen(buf)
	go network.RepublishData()

	return network
}

//Listening for new packets on ip, port combination
func (network *Network) Listen(buf []byte) {
	defer network.listenConnection.Close()

	//For each new packet do marshalling
	for {
		n, addr, _ := network.listenConnection.ReadFromUDP(buf)
		message := &WrapperMessage{}
		replyErr := proto.Unmarshal(buf[0:n], message)

		if (message.ID[0:5] == "Reply") {
			go network.HandleReply(message, replyErr, addr)
		} else {
			go network.HandleRequest(message, replyErr, addr)
		}
	}
}

func (network *Network) SendPingMessage(contact Contact, returnChannel chan interface{}) {
	messageID := NewRandomKademliaID()
	remoteAddr, err := net.ResolveUDPAddr("udp", contact.Address)
	CheckError(err)

	packet := &RequestPing{}
	wrapperMsg := &WrapperMessage_RequestPing{packet}
	wrapper := &WrapperMessage{"RequestPing", network.node.rt.me.ID.String(), messageID.String(), wrapperMsg}

	network.createChannel(messageID, returnChannel)
	network.sendPacket(network.marshalHelper(wrapper), remoteAddr)

	go network.TimeoutWaiter(network.timeoutTime, returnChannel, messageID)
}


func (network *Network) SendFindContactMessage(targetID *KademliaID, contact *Contact, returnChannel chan interface{}) {
	messageID := NewRandomKademliaID()
	remoteAddr, err := net.ResolveUDPAddr("udp", contact.Address)
	CheckError(err)

	packet := &RequestContact{targetID.String()}
	wrapperMsg := &WrapperMessage_RequestContact{packet}
	wrapper := &WrapperMessage{"RequestContact", network.node.rt.me.ID.String(), messageID.String(), wrapperMsg}

	network.createChannel(messageID, returnChannel)
	network.sendPacket(network.marshalHelper(wrapper), remoteAddr)
	
	go network.TimeoutWaiter(network.timeoutTime, returnChannel, messageID)

}


func (network *Network) SendFindDataMessage(hash string, contact *Contact, returnChannel chan interface{}) {
	messageID := NewRandomKademliaID()
	remoteAddr, err := net.ResolveUDPAddr("udp", contact.Address)
	CheckError(err)

	packet := &RequestData{hash}
	wrapperMsg := &WrapperMessage_RequestData{packet}
	wrapper := &WrapperMessage{"RequestData", network.node.rt.me.ID.String(), messageID.String(), wrapperMsg}

	network.createChannel(messageID, returnChannel)
	network.sendPacket(network.marshalHelper(wrapper), remoteAddr)

	go network.TimeoutWaiter(network.timeoutTime, returnChannel, messageID)
}

func (network *Network) SendStoreMessage(hash string, address string, returnChannel chan interface{}) {
	messageID := NewRandomKademliaID()
	remoteAddr, err := net.ResolveUDPAddr("udp", address)
	CheckError(err)
	data := "PLS_REMOVE_ME"

	packet := &RequestStore{hash, data}  //EDIT ME
	wrapperMsg := &WrapperMessage_RequestStore{packet}
	wrapper := &WrapperMessage{"RequestStore", network.node.rt.me.ID.String(), messageID.String(), wrapperMsg}

	network.createChannel(messageID, returnChannel)
	network.sendPacket(network.marshalHelper(wrapper), remoteAddr)

	go network.TimeoutWaiter(network.timeoutTime, returnChannel, messageID)
}

func (network *Network) HandleReply(message *WrapperMessage, replyErr error, sourceAddress *net.UDPAddr) {
	if replyErr != nil {
		fmt.Println(message.ID)
		log.Println("Something went wrong in Listen, err: ", replyErr)
		return
	}

	answerChannel := network.getAnswerChannel(NewKademliaID(message.RequestID))

	if(answerChannel == nil) {
		fmt.Println("Forged Reply or Timeout")
		return
	}


	switch message.ID {
		case "ReplyPing":
			contact := NewContact(NewKademliaID(message.GetReplyPing().GetID()), message.GetReplyPing().GetAddress())
			answerChannel<-contact
			return

		case "ReplyContactList":
			contactList := []Contact{}
			for i := range message.GetReplyContactList().GetContacts() {
				contactList = append(contactList, NewContact(NewKademliaID(message.GetReplyContactList().Contacts[i].GetID()), message.GetReplyContactList().Contacts[i].GetAddress()))
			}
			answerChannel<-contactList
			break

		case "ReplyData":
			answerChannel <- sourceAddress.String()
			break

		case "ReplyStore":
			answerChannel<-message.GetReplyStore().GetData()
			break

		default:
			fmt.Println("Not a valid Reply ID. ID: " + message.ID)
			return
	}
	go network.updateRoutingTable(message.SourceID, sourceAddress.String())

}


func (network *Network) HandleRequest(message *WrapperMessage, replyErr error, sourceAddress *net.UDPAddr) {
	if replyErr != nil {
		fmt.Println(message.ID)
		log.Println("Something went wrong in Listen, err: ", replyErr)
		return
	}

	var wrapper *WrapperMessage

	switch message.ID {
		case "RequestPing":
			fmt.Println("Ping Recieved")
			packet := &ReplyPing{network.node.rt.me.ID.String(), network.node.rt.me.Address}
			wrapperMsg := &WrapperMessage_ReplyPing{packet}
			wrapper = &WrapperMessage{"ReplyPing", network.node.rt.me.ID.String(), message.RequestID, wrapperMsg}
			break

		case "RequestContact":
			network.rtMux.Lock()
			contactListReply := network.getClosestContacts(message.GetRequestContact().GetTarget())
			network.rtMux.Unlock()

			packet := &ReplyContactList{contactListReply}
			wrapperMsg := &WrapperMessage_ReplyContactList{packet}
			wrapper = &WrapperMessage{"ReplyContactList", network.node.rt.me.ID.String(), message.RequestID, wrapperMsg}			
			break


		case "RequestData":
			haveData := network.node.gotData(*NewKademliaID(message.GetRequestData().GetKey()))
			if(haveData) {		
				packet := &ReplyData{"PLSREMOVEME"}
				wrapperMsg := &WrapperMessage_ReplyData{packet}
				wrapper = &WrapperMessage{"ReplyData", network.node.rt.me.ID.String(), message.RequestID, wrapperMsg}

			} else {
				network.rtMux.Lock()
				contactListReply := network.getClosestContacts(message.GetRequestData().GetKey())
				network.rtMux.Unlock()

				packet := &ReplyContactList{contactListReply}
				wrapperMsg := &WrapperMessage_ReplyContactList{packet}
				wrapper = &WrapperMessage{"ReplyContactList", network.node.rt.me.ID.String(), message.RequestID, wrapperMsg}
			}
			
			break


		case "RequestStore":
			fileID := *NewKademliaID(message.GetRequestStore().GetKey())
			haveFile := network.node.gotData(fileID)
			if(!haveFile) {
				go network.fileNetwork.downloadFile(&fileID, sourceAddress.String(), false)
			} else {
				network.node.Store(fileID, time.Now())
			}

			packet := &ReplyStore{"ok"}
			wrapperMsg := &WrapperMessage_ReplyStore{packet}
			wrapper = &WrapperMessage{"ReplyStore", network.node.rt.me.ID.String(), message.RequestID, wrapperMsg}
			break

		default:
			fmt.Println("Not a valid Request ID. ID: " + message.ID)
			return

	}
	go network.updateRoutingTable(message.SourceID, sourceAddress.String())
	network.sendPacket(network.marshalHelper(wrapper), sourceAddress)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func (network *Network) getClosestContacts(target string) (contactListReply []*ReplyContactList_Contact) {
	closestContacts := network.node.rt.FindClosestContacts(NewKademliaID(target), 20)
	contactListReply = []*ReplyContactList_Contact{}
	for i := range closestContacts {
		contactReply := &ReplyContactList_Contact{closestContacts[i].ID.String(), closestContacts[i].Address}
		contactListReply = append(contactListReply, contactReply)
	}
	return contactListReply
}

func (network *Network) createChannel(messageID *KademliaID, returnChannel chan interface{}) {
	network.mux.Lock()
	network.waitingAnswerList[*messageID] = returnChannel
	network.mux.Unlock()
}

func (network *Network) getAnswerChannel(requestID *KademliaID) (answerChannel chan interface{}) {
	network.mux.Lock()
	answerChannel = network.waitingAnswerList[*requestID]
	network.mux.Unlock()
	return answerChannel
}


func (network *Network) TimeoutWaiter(sleepTime int, returnChannel chan interface{}, messageID *KademliaID) {
	time.Sleep(time.Duration(sleepTime) * time.Second)
	network.mux.Lock()
	network.waitingAnswerList[*messageID] = nil
	network.mux.Unlock()
	returnChannel <-false
}

func (network *Network) marshalHelper(wrapper *WrapperMessage) []byte{
	data, err := proto.Marshal(wrapper)
	if err != nil {
		log.Fatal("Marshall Error: ", err)
	}
	return data
}

func (network *Network) sendPacket(data []byte, targetAddress *net.UDPAddr) {
	buf := []byte(data)
	_, err := network.listenConnection.WriteToUDP(buf, targetAddress)
	if err != nil {
		log.Println(err)
	}
}

func (network *Network) updateRoutingTable(contactID string, contactAddress string) {
	contact := NewContact(NewKademliaID(contactID), contactAddress)
	network.rtMux.Lock()
	network.node.rt.AddContactNetwork(contact, network)
	network.rtMux.Unlock()
}

func (network *Network) RepublishData() {
	time.Sleep(time.Duration(network.republishSleepTimeSeconds) * time.Second)
	//fmt.Println("Republish Check")
	dataMap := network.node.getDataMap()
	for dataEntryID, timestamp := range dataMap {
		if(time.Now().After(timestamp)) {
			kademlia := NewKademlia(network)
			contactList, _ := kademlia.LookupContact(&dataEntryID, false)
			i := 0

			for k := range contactList {
				i++
				if(contactList[k].ID.String() == network.node.rt.me.ID.String()) {
					network.node.Store(dataEntryID, time.Now())
				} else {
					go network.sendStoreAndWaitForAnswer(dataEntryID.String(), contactList[k].Address, i)
				}
			}
			fmt.Println("Sent Republish")
		}
		if(time.Now().After(timestamp.Add(time.Duration(2) * time.Second))) {
			go network.node.deleteEntry(dataEntryID, network.fileNetwork.mux2)
		}
	}
	network.RepublishData()
}

func (network *Network) sendStoreAndWaitForAnswer(dataEntryID string, address string, number int) {
	
	returnChannel := make(chan interface{})
	go network.SendStoreMessage(dataEntryID, address, returnChannel)
	returnValue:= <-returnChannel
	switch returnValue := returnValue.(type) {
		case string:
			fmt.Println("Store " + strconv.Itoa(number) + " Reply: " + returnValue)
		case bool:
			fmt.Println("Store request timeout")
		default:
			fmt.Println("Something went wrong")
	}
}
