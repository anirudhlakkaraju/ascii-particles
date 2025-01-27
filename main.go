package main

import (
	"flag"
	"fmt"
	"log"
	"particles/particles"
	"time"
)

var c = `               .........:--===++++======================++++===--:.........
              ...:-=*#%%%%%%%%%%###***+++++++++++++****##%%%%%%%%%%#*+=::...
              :::.:--=+**##%%%%%%%###***************####%%%%%%%##**+=-:..:::
              :::::::..........::::::::::::::::::::::::::::.........::::::::
              ::::........................     ...................:::::::::.
              :===:......................        .................:::::--==..
               :---==-:::-=-:::..........        ..........::--=====+*+=--:......
               ::::..::::-==+++******++++++++++++++++******++======-::::::.   ....
               .:::........................::::..................:::::::::     ....
                ::::............................................:::::::::.     ....
                 :::......::..................................:::::::::::      ....
                 .:::.....::::..............................::::::::::::.    .....
                  ::::.....:::::::......................:::::::::::::::::::......
                   ::::....::::::::::::::::....::::::::::::::::::::::::::::....
                  .:::::...:::::::::::::::::::::::::::::::::::::::::::..
             ..::---:::::...::::::::::::::::::::::::::::::::::::::::::---:::.
         .::---------::::::..::::::::::::::::::::::::::::::::::::::::---------::.
       ::-----------::::::::..:::::::::::::::::::::::::::::::::::::::------------::.
     .:-------::::::::::::::::.::::::::::::::::::::::::::::::::::::::::::::---------:
    .-----:::::::::::::----:::::::::::::::::::::::::::::::::::------:::::::::::------.
    .--:::::::::::::::-----===-::::::::::::::::::::::::::::-===-----::::::::::::::---.
    ..::::::::::::::::::---=+***+=-::::::::::::::::::::-=+**+==---::::::::::::::::::..
     ..::::::::::::::::::::--==++***++==---------===++**+++=--:::::::::::::::::::::..
     ..:-::::::::::::::::::::::::-----==============-----::::::::::::::::::::::::-:..
       ..:==-::.::::::::::::::::....:::::::::::::::::::::::::::::::::::::::.:-==:..
         ...:-=+==-::...:::::::::::::...............:::::::::::::::..::--=+==::..
            ...::-=++**++==---:::.........................:::--==++**++=-::...
                 ....::--==++***#####*****************####***++==--::.....
                       ......::::::----============----::::::......
                                 ........................


`

var effects = map[string]particles.ASCII{
	"steam": particles.AsciiSteam,
	"fire":  particles.AsciiFire,
}

func main() {

	argType := flag.String("effect", "steam", "Specify the particle effect: 'steam' or 'fire'")
	flag.Parse()

	effect, ok := effects[*argType]
	if !ok {
		log.Fatalf("Invalid effect type: %s", *argType)
	}

	// Pour and enjoy!
	coffee := particles.NewCoffee(61, 8, 9.0, effect)
	coffee.Start()

	timer := time.NewTicker(100 * time.Millisecond)
	for {
		<-timer.C
		fmt.Printf("\033[H\033[2J")
		coffee.Update()
		steam := coffee.Display()
		for _, row := range steam {
			fmt.Printf("              %s\n", row)
		}
		fmt.Println(c)
	}
}
