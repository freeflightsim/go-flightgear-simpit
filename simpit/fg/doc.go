

// Interface to a running FlightGear instance
package fgio


/*
The idea is to contain the FlightGear the interface
in this package, whether its Ws websocket or Udp
- websocket for simple listeners
- protocol udp for faster stuff (later and requires protocol file)

so far--
- Sends websocket messages ie Commands
- recieves websocket frames, and channel to process


Newbie Problems
- websocket restart, reconnect prolblems



Properties = nodes



 */




