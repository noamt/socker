package main

type chatHub struct {
    registeredConnections map[*chatConnection]bool
    inboundMessages chan []byte
    registerRequests chan *chatConnection
    unregisterRequests chan *chatConnection
}

func newChatHub() *chatHub {
    return &chatHub {
        registeredConnections: make(map[*chatConnection]bool),
	inboundMessages: make(chan []byte),
	registerRequests: make(chan *chatConnection),
	unregisterRequests: make(chan *chatConnection),
    }
}

func (hub *chatHub) run() {
    for {
        select {
	    case c := <-hub.registerRequests:
	        hub.registeredConnections[c] = true
	    case c := <-hub.unregisterRequests:
	        if _, ok := hub.registeredConnections[c]; ok {
		    delete(hub.registeredConnections, c)
		    close(c.out)
		}
	    case m := <-hub.inboundMessages:
	        for c := range hub.registeredConnections {
		    select {
		        case c.out <- m:
			default:
			    delete(hub.registeredConnections, c)
			    close(c.out)
		    }
		}
	}
    }
}
