//	Anthem of Ukraine:
//
//	Ще не вмерла України і слава, і воля,
//	Ще нам, браття молодії, усміхнеться доля.
//	Згинуть наші воріженьки, як роса на сонці.
//	Запануєм i ми, браття, у своїй сторонці.
//
//	Душу й тіло ми положим за нашу свободу,
//	І покажем, що ми, браття, козацького роду!
//
//	Душу й тіло ми положим за нашу свободу,
//	І покажем, що ми, браття, козацького роду!

package main

import (
	"fmt"
	"time"
)

func main() {
	// Print "Слава Україні!" and "Героям слава!" with a second pause 
	for {
		fmt.Println("Слава Україні!")
		time.Sleep(time.Second)
		fmt.Println("Героям слава!")
		time.Sleep(time.Second)
	}
}
