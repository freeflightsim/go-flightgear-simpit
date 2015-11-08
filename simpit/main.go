
package main

import (
	"os"
	"os/signal"
	"fmt"

	"github.com/freeflightsim/go-flightgear-simpit/simpit/config"
	"github.com/freeflightsim/go-flightgear-simpit/simpit/fg"
	//"github.com/freeflightsim/go-flightgear-simpit/simpit/vbus"
)

func main() {

	//TODO, flags for host, port, not sure where its gonna go yet, even config


	//= Handle ctrl+c to kill on terminal (required as were `multithread`)
	killChan := make(chan os.Signal, 1)
	signal.Notify(killChan, os.Interrupt)

	//= Load Config (this also loads def in protocol/
	conf, err := config.Load("../config/787.yaml")
	if err != nil {
		fmt.Println(" Config error= ", err)
		return
	}
	fmt.Println(" config = ", conf)


	//= Initialise the local State store thingi
	//bus := vbus.NewVBus()
	//bus.AddProps(  conf.GetOutputNodes() )

	// Some custom nodes in dev
	//ias_node := "/instrumentation/airspeed-indicator/indicated-speed-kt"
	//bus.AddProp(ias_node)

	//hdg_bug := "/autopilot/settings/heading-bug-deg"
	//bus.AddProp(hdg_bug)

	//eng_node := "/controls/engines/engine[1]/throttle"
	//bus.AddProp(eng_node)

	//= Initialise the flightgear client(s) later udp
	fg_client := fgio.NewFgClient("192.168.50.153", "7777")
	//fg_client.AddProps( bus.GetProps() )


	//=  Piface Digital IO board initialisaction (rpi only)
	//pdf_board := piio.NewPifaceBoard()
	//pdf_board.Init()
	//if pdf_board.Enabled == false {
	// On a pc with no piface, we fake inputs with timers
	//board.PretendInputs( conf.DInPins )
	//}

	//= Arduino Board (current dev is Duemilanove .. olde)
	//arduino_1 := ardio.NewArduinoBoard("ard_1")
	//arduino_2 := ardio.NewArduinoBoard("ard_mega_1")
	//ard_board.LoadConfig(conf.FgNodes)


	go fg_client.Start()
	//go arduino_1.Run()
	//go arduino_2.Run()


	//var last_v int64

	// Route all messages
	for {
		select {

		//= ctrl+c to kill
		case  <- killChan:
		// TODO gracefully shutdown things
			fmt.Println( "killed" )
			os.Exit(0)




		// Messages from Flightgear
		case msg := <- fg_client.WsChan:
			fmt.Println("", msg.Prop, msg.StrValue())
			//bus.Update( msg.Prop, msg.StrValue() )

			//if msg.Prop == eng_node {
			//	fmt.Println("eng", msg.StrValue() )
			//}

			/*
			for _, out_p := range conf.DOutPins {

				if out_p.Prop == msg.Prop {
					fmt.Println("        YES = ", msg.Prop)
					//on := out_p.IsOn( msg.StrValue() )
					//pdf_board.SetOutput(out_p.Pin, on)

				}
			}
			*/




		}
	}


}
