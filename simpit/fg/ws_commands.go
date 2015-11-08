
package fgio

import (
	//"encoding/json"
)


// Websocket Command
type WsCommand struct {
	Cmd string ` json:"command" `
	Node string ` json:"node" `
}


// A web socket command with nodes
type WsCommandVal struct {
	Cmd string ` json:"command" `
	Node string ` json:"node" `
	Value string ` json:"value" `
}


