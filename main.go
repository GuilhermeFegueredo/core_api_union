package main

import (
	"core_APIUnion/src/config"
	"core_APIUnion/src/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	/*
		// Gera uma string aleatória base 64 para ser usada como secret de token
		secret := func() {
			key := make([]byte, 64)
			_, err := rand.Read(key)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(base64.StdEncoding.EncodeToString(key))
		}
		secret()
	*/
	config.Load()
	r := routes.Generate()

	fmt.Printf("Escutando na porta %d: ", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
