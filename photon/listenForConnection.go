package photon

import (
	"log"
	"net/http"
	"time"
)

func ListenForConnection(fileServerPort string) {
	for {
		time.Sleep(time.Second)

		fileServerCheck := listen(fileServerPort)

		if fileServerCheck == "continue" {
			continue
		}

		// Reached this point: server is up and running!
		break
	}
	log.Println("Server up and running!")
}

func listen(port string) string {
	log.Println("Checking if started...")

	resp, err := http.Get("http://localhost" + port + "/index.html")

	if err != nil {
		log.Println("Failed:", err)
		return "continue"
	}

	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Not OK:", resp.StatusCode)
		return "continue"
	}

	return "break"
}
