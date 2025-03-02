package main

import (
	"fmt"
	"github.com/gordonklaus/portaudio"
	"log"
	"time"
)

func main() {
	//	reader := bufio.NewReader(os.Stdin)
	//	request, _ := reader.ReadString('\n')
	//	date, _ := check_date(request)
	//	fmt.Println(date)

	err := portaudio.Initialize()
	if err != nil {
		log.Fatalf("Erreur PortAudio : %v", err)
	}
	defer portaudio.Terminate()

	fmt.Println("✅ PortAudio installé avec succès sur Windows !")
	time.Sleep(2 * time.Second)

}
