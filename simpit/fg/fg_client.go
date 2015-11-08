
package fgio



import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"


	"golang.org/x/net/websocket"
)

type FgClient struct {

	//Enabled bool
	Host string
	HttpPort string

	Props map[string]bool

	WsConn *websocket.Conn
	WsConnected bool
	WsChan chan WsPacket


	UdpConn *net.UDPConn
	UdpChan chan interface{}

}

// Creates a new FlightGear client
func NewFgClient(host string, http_port string) *FgClient{

	c := new(FgClient)
	c.Host = host
	c.HttpPort = http_port

	c.Props = make(map[string]bool)
	c.WsChan = make(chan WsPacket)
	//c.WsConnected = false
	//c.UdpChan = make(chan UdpPacket)

	return c
}

func (me *FgClient) Connect() error {

	//me.WsConnect()

	//me.UdpConnect()


	return nil
}

func (me *FgClient) WsStart() error {

	origin := "http://" + me.Host + ":" + me.HttpPort
	url := "ws://" + me.Host + ":" +  me.HttpPort + "/PropertyListener"

	var bits = make([]byte, 512)
	var n int
	var err error
	var packet WsPacket

	for {

		me.WsConn, err = websocket.Dial(url, "", origin)
		fmt.Println("Connecting", n, err)
		if err != nil {
			log.Println(err)
			time.Sleep(1 * time.Second)

		} else {

			for {
				n, err = me.WsConn.Read(bits)
				if err != nil {
					fmt.Println("WS Read err", n, err, me.WsConn)
				} else {
					//fmt.Println("rcv", string(bits[:n]))
					err := json.Unmarshal(bits[:n], &packet)
					if err != nil {
						fmt.Println("WS json decode error", err)
					} else {
						me.WsChan <- packet
					}
				}
			}
		}

	}
	return nil
}

// Start up, connect, start listener  props
func (me *FgClient) Start() error {

	//err := me.Connect()
	//if err != nil {
	//	fmt.Println("Fatal, cannot start", err)
	//}

	go me.WsStart()

	// add props to listen to
	for node, _ := range me.Props {
		me.WsAddListener(node)
	}
	// fire off even to get current value of props
	// TODO investegate whether we can get fgfs to send back val with addListener
	for node, _ := range me.Props {
		me.WsGet(node)
	}
	return nil
}


// Websocket listener started in go routine
/*
func (me *FgClient) WsListen(){

	var bits = make([]byte, 512)
	var n int
	var err error
	var packet WsPacket
	for {
		n, err = me.WsConn.Read(bits)
		if err != nil {
			fmt.Println("WS Read err", n, err)
		} else {
			//fmt.Println("rcv", string(bits[:n]))
			err := json.Unmarshal(bits[:n], &packet)
			if err != nil {
				fmt.Println("WS json decode error", err)
			} else {
				me.WsChan <- packet
			}
		}
	}
}
*/

// Add a list of props to listen on
func (me *FgClient) AddProps(props []string){

	// next we add the nodes
	for _, n := range props {
		me.Props[n] = true
	}

}



// Send ws
func (me *FgClient) WsAddListener(prop string){
	//fmt.Println(" + AddListener", node)
	me.WsSendCommand( WsCommand{"addListener", prop} )
}

// Sends a websocket Get to fgfs
func (me *FgClient) WsGet(prop string){
	me.WsSendCommand( WsCommand{"get", prop} )
}

// Set a property with a websocket..
func (me *FgClient) WsSet(prop string, value string) {
	// TODO maybe here we can check for type/validate to ensure not crash fgfs
	me.WsSendCommand( WsCommandVal{"set", prop, value} )
}



// Send a websocket command to fgfs
func (me *FgClient) WsSendCommand(comm interface{}) error {
	//fmt.Println("SendCommand", comm)
	if me.WsConn == nil {
		return nil
	}
	bits, err := json.Marshal(comm)
	if err != nil {
		fmt.Println("SendCommand.jsonerror", err)
		return err
	}
	//fmt.Println("bits", string(bits))
	// TODO catch
	if _, err := me.WsConn.Write(bits); err != nil {
		//log.Fatal(err)
		fmt.Println("written", err)
	}
	return nil
}


