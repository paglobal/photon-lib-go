package photon

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var onEvents IPCHubEventsMap = make(IPCHubEventsMap)
var onceEvents IPCHubEventsMap = make(IPCHubEventsMap)
var ipcs IPCMap = make(IPCMap)

var IPCHubInstance IPCHub = IPCHub{
	onEvents,
	onceEvents,
	ipcs,
}

func IPCInit() {
	router := gin.Default()

	router.GET(
		"/ipc", func(c *gin.Context) {
			socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
			if err != nil {
				log.Println(err)
				return
			}

			onEvents := make(EventsMap)
			onceEvents := make(EventsMap)

			ipc := IPC{
				onEvents,
				onceEvents,
				socket,
				"",
			}

			IPCHubInstance.AddIPC(&ipc)

			defer socket.Close()
			defer IPCHubInstance.RemoveIPC(&ipc)

			for {
				//Read message from browser
				var data Data
				err := socket.ReadJSON(&data)
				if err != nil {
					log.Println(err)
					return
				}

				onEvents := ipc.ReturnEventsMap("on")
				onceEvents := ipc.ReturnEventsMap("once")

				for _, v := range onEvents[data.Event] {
					v(data.Payload, &ipc)
				}

				for _, v := range onceEvents[data.Event] {
					v(data.Payload, &ipc)
				}

				delete(onceEvents, data.Event)
			}
		})

	router.Run(":53174")
}
