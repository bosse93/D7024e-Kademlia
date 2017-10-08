package main

import (
	"fmt"
	"strconv"
	"time"
	"net"
	"encoding/hex"
	//"io/ioutil"
	"strings"
	"os"
	"net/http"
	"io"
	"log"
)

func main() {
	StartNetwork()
}

func decodeHash(hash string) string{
	byteArray := []byte(hash)

	for i := 0; i < 19; i++ {
		if((string(byteArray[i*2]) == "0") && (string(byteArray[(i*2)+1]) == "3")) {
			fileName, _ := hex.DecodeString(string(byteArray[:(i)*2]))
			return string(fileName)
		}
	}
	return "Error when decoding dataID"
	/*
	fmt.Println("DECODER")
	fmt.Println(hash)

	index := strings.IndexByte(hash, byte("03"))
	fmt.Println(index)
	fileName, _ := hex.DecodeString(hash[:index-1])
	*/
	//return string(byteArray)
}

func HashKademliaID(fileName string) *KademliaID{
	fmt.Println("Fil Namn: " + fileName)
	f := hex.EncodeToString([]byte(fileName))
	if(len(f) > 38) {
		fmt.Println(f)
		fmt.Println("Name of file can be maximum 19 characters, including file extension.")
	}
	f = f + "03"
	for len(f) < 40 {
		f = f + "01"
	}
	return NewKademliaID(f)
}

func HandleRequest(conn *net.UDPConn, addr *net.UDPAddr, args []string, network *Network){
	//_,err := conn.WriteToUDP([]byte("From server: Hello I got your mesage " + p), addr)

	if args[0]=="store" {
		fmt.Println("this was a store message with arg "+ args[0])
		kademlia := NewKademlia(network)
		//FFFFFFFF0FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF

		newKad := HashKademliaID(args[1])
		kademlia.Store(newKad.String())
		_,storeErr := conn.WriteToUDP([]byte("stored: "+newKad.String()), addr)
		if storeErr != nil {
			fmt.Println("something went shit in store: %v", storeErr)
		}

	} else if args[0]=="cat" {
		/*
		kademlia := NewKademlia(network)
		//FFFFFFFF0FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF

		newKad := HashKademliaID(args[1])
		data, success := kademlia.LookupData(newKad.String())
		if success {
			_,err := conn.WriteToUDP([]byte(data), addr)
			if err != nil {
				fmt.Println("something went shit in lookup: %v", err)
			}
		} else {
			_,err := conn.WriteToUDP([]byte("no data found"), addr)
			if err != nil {
				fmt.Println("something went shit in lookup: %v", err)
			}
		}
		*/

	}

}

func CreateNodes(amount int) (firstNetwork *Network){
	

	
	firstNode := NewContact(NewRandomKademliaID(), "localhost:8000")
	firstNodeRT := NewRoutingTable(firstNode)
	node := NewNode(firstNodeRT)
	lastTCPNetwork := NewFileNetwork(node, "localhost", 8000)
	firstNetwork = NewNetwork(node, lastTCPNetwork, "localhost", 8000)
	//nodeList := []*RoutingTable{firstNodeRT}
	//lastNode := firstNode
	//create 100 nodes
	if _, err := os.Stat("kademliastorage/" + firstNode.ID.String()); os.IsNotExist(err) {
			os.Mkdir("kademliastorage/" + firstNode.ID.String(), 0755)
	}

	if _, err := os.Stat("upload/" + firstNode.ID.String()); os.IsNotExist(err) {
			os.Mkdir("upload/" + firstNode.ID.String(), 0755)
	}

	fileDst, _ := os.Create("upload/" + firstNode.ID.String() + "/workshop.jpeg")
	fileSrc, _ := os.Open("workshop.jpeg")

	if _, err := io.Copy(fileDst, fileSrc); err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat("downloads/" + firstNode.ID.String()); os.IsNotExist(err) {
		os.Mkdir("downloads/" + firstNode.ID.String(), 0755)
	}
	

	for i := 0; i < amount; i++ {
		port := 8001 + i
		a := "localhost:" + strconv.Itoa(port)

		ID := NewRandomKademliaID()
		rt := NewRoutingTable(NewContact(ID, a))
		//nodeList = append(nodeList, rt)
		rt.AddContact(firstNodeRT.me)
		node := NewNode(rt)
		tcpNetwork := NewFileNetwork(node, "localhost", port)
		nw := NewNetwork(node, tcpNetwork, "localhost", port)
		fmt.Println("Ny Nod varv " + strconv.Itoa(i+1) + ": " + rt.me.String())
		//go nw.Listen("localhost", port)
		time.Sleep(500 * time.Millisecond)
		kademlia := NewKademlia(nw)

		contactResult, _  := kademlia.LookupContact(ID, false)
		if(len(contactResult) > 0) {
			for q := range contactResult {
				rt.AddContact(contactResult[q])
			}
		}
		
		if _, err := os.Stat("kademliastorage/" + ID.String()); os.IsNotExist(err) {
			os.Mkdir("kademliastorage/" + ID.String(), 0755)
		}

		if _, err := os.Stat("upload/" + ID.String()); os.IsNotExist(err) {
			os.Mkdir("upload/" + ID.String(), 0755)
		}

		if _, err := os.Stat("downloads/" + ID.String()); os.IsNotExist(err) {
			os.Mkdir("downloads/" + ID.String(), 0755)
		}
		
	}
	return
}

func StartFrontend(lastNetwork *Network){

	addr := net.UDPAddr{
		Port: 1234,
		IP: net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		p := make([]byte, 2048)
		_,remoteaddr,err := ser.ReadFromUDP(p)
		split := strings.Split(string(p), " ")
		fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
		fmt.Println(split)
		if split[0] == "cat"{
			fmt.Println("I got cat back do cat stuff")
		} else if split[0] == "store" {
			fmt.Println("I got Store back")
		}
		if err !=  nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		//go sendResponse(ser, remoteaddr)
		go HandleRequest(ser, remoteaddr, split, lastNetwork)

	}
}

func StartNetwork() {
	if _, err := os.Stat("kademliastorage/"); os.IsNotExist(err) {
		os.Mkdir("kademliastorage", 0755)
	} else {
		os.RemoveAll("kademliastorage")
		time.Sleep(500 * time.Millisecond)
		os.Mkdir("kademliastorage", 0755)
	}

	if _, err := os.Stat("upload/"); os.IsNotExist(err) {
		os.Mkdir("upload", 0755)
	} else {
		os.RemoveAll("upload")
		time.Sleep(500 * time.Millisecond)
		os.Mkdir("upload", 0755)
	}
	
	if _, err := os.Stat("downloads/"); os.IsNotExist(err) {
		os.Mkdir("downloads", 0755)
	} 


	//Creates x amount of nodes in a network
	firstNetwork := CreateNodes(10)
	//defer removeDirectories(directories)
	//printFirstNodeRT(firstNode, firstNodeRT)
	//printLastNodeRT(nodeList)

	//testStore := HashKademliaID("workshop.jpeg")

	//pwd, _ := os.Getwd()
	//dat, err := ioutil.ReadFile(pwd+"/../src/D7024e-Kademlia/main/testStore.txt")
	//dat, err := ioutil.ReadFile("main/testStore.txt")
	//check(err)

	kademlia := NewKademlia(firstNetwork)
	//store link to workshop jpg
	fileName := "workshop.jpeg"

	kademlia.Store(fileName)
	time.Sleep(3*time.Second)
	kademlia = NewKademlia(firstNetwork)
	//lookup workshop jpg
	
	success := kademlia.LookupData(fileName)
	if(success) {
		fmt.Println("Data found and downloaded")
	} else {
		fmt.Println("Data not found")
	}

	//Setup Frontend

	//downloadFile("workshop.jpeg", "https://www.dropbox.com/s/b0a98iiuu1o9m5y/Workshopmockup-1.jpg?dl=1")
	StartFrontend(firstNetwork)


	/*for k1, v := range IDRTList {
		for k2, v2 := range v.node.data {
			fmt.Println("Node " + k1.String() + " has " + v2 + " stored for key " + k2.String())
		}
	}*/

}
func removeDirectories(directories []string) {
	fmt.Println("in remove")
	for i := range directories{
		os.Remove(directories[i])
	}
}


func printLastNodeRT(nodeList []*RoutingTable) {
	lastNode := nodeList[len(nodeList)-1]
	fmt.Println("Node: " + lastNode.me.String())
	for z := range lastNode.buckets {
		contactList := lastNode.buckets[z]
		fmt.Println("Bucket: " + strconv.Itoa(z))
		for elt := contactList.list.Front(); elt != nil; elt = elt.Next() {
			contact := elt.Value.(Contact)
			fmt.Println(contact.String())
		}
	}
}



func check(e error) {
	if e != nil {
		panic(e)
	}
}

func downloadFile(filepath string, url string) (err error) {
	fmt.Println("filepath: " + filepath + " url: " + url)
	// Create the file
	out, err := os.Create(filepath)
	if err != nil  {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil  {
		return err
	}

	return nil
}