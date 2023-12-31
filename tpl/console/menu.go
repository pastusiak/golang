package console

import (
	"fmt"
)

func ShowMenu() {
	fmt.Println("")
	fmt.Println(" __________________________")
	fmt.Println("|                          |")
	fmt.Println("| APPLICATION MENU         |")
	fmt.Println("|__________________________|")
	fmt.Println("|")

	for _, item := range cmdList {
		trigger, description := item.CmdConfig()
		fmt.Printf("| [%s] %s\n\r", trigger, description)
	}

	fmt.Println("|__________________________")
	fmt.Println("|                          |")
	fmt.Println("| SYSTEM MENU              |")
	fmt.Println("|__________________________|")
	fmt.Println("|")
	fmt.Println("| [q] Shutdown application")
	fmt.Println("| [?] Command list")
	fmt.Println("|__________________________")
	fmt.Println("")
	fmt.Print("Select option: ")
}
