package main

import (
	"fmt"
	api_adapter_servidor "genexis/pos/autoservicios/presentation/api/gin/adapter"
	"genexis/pos/autoservicios/presentation/container"
	"log"
	"os"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(os.Stdout)
	fmt.Println("  __  __  _____           _    _ _______ ____     _____ ______ _______      _______ _____ ______   _____   ____   _____")
	fmt.Println(" |  \\/  |/ ____|     /\\  | |  | |__   __/ __ \\   / ____|  ____|  __ \\ \\    / /_   _/ ____|  ____| |  __ \\ / __ \\ / ____|")
	fmt.Println(" | \\  / | (___      /  \\ | |  | |  | | | |  | | | (___ | |__  | |__) \\ \\  / /  | || |    | |__    | |__) | |  | | (___ ")
	fmt.Println(" | |\\/| |\\___ \\    / /\\ \\| |  | |  | | | |  | |  \\___ \\|  __| |  _  / \\ \\/ /   | || |    |  __|   |  ___/| |  | |\\___ \\ ")
	fmt.Println(" | |  | |____) |  / ____ \\ |__| |  | | | |__| |  ____) | |____| | \\ \\  \\  /   _| || |____| |____  | |    | |__| |____) |")
	fmt.Println(" |_|  |_|_____/  /_/    \\_\\____/   |_|  \\____/  |_____/|______|_|  \\_\\  \\/   |_____\\_____|______| |_|     \\____/|_____/ ")

	if err := container.InitializeContainer(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	log.Println("--- INICIANDO AUTO SERVICE POS ---")
	if err := api_adapter_servidor.Start(); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
