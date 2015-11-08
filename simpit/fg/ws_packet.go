
package fgio

import (
	"strconv"
	"fmt"
)



type WsPacket struct {
	Prop string ` json:"path" `
	Name string ` json:"name" `
	Type string ` json:"type" `
	Index int ` json:"index" `
	NChildren int ` json:"nChildren" `
	RawValue interface{} ` json:"value" `
}

// Returns RawValue as string regardless of Type
// todo decide is use 32 bit ints cos it makes life
// seasier with arduino and rpi(arm)
func (me *WsPacket) StrValue() string {

	switch me.Type {
	case "string":
		return me.RawValue.(string)

	case "double":
		return fmt.Sprintf("%.f", me.RawValue)

	case "float":
		return strconv.FormatFloat(me.RawValue.(float64), 'f', 20, 64)
	}

	return "#### OOOPS ##########"
}
