package websocket

func (h *Hub) SendToClient(id string, group string, message []byte) {
	data := &MessageData{
		Id:      id,
		Group:   group,
		Message: message,
	}

	h.Message <- data
}

func (h *Hub) SendToGroup(group string, message []byte) {
	data := &GroupMessageData{
		Group:   group,
		Message: message,
	}

	h.GroupMessage <- data
}

func (h *Hub) SendToAll(message []byte) {
	data := &BroadcastMessageData{
		Message: message,
	}

	h.BroadcastMessage <- data
}

func (h *Hub) SendToClientService() {
	for {
		select {
		case data := <-h.Message:
			if groupMap, ok := h.Group[data.Group]; ok {
				if client, ok := groupMap[data.Id]; ok {
					client.Send <- data.Message
				}
			}
		}
	}
}

func (h *Hub) SendToGroupService() {
	for {
		select {
		case data := <-h.GroupMessage:
			if groupMap, ok := h.Group[data.Group]; ok {
				for _, client := range groupMap {
					client.Send <- data.Message
				}
			}
		}
	}
}

func (h *Hub) SendToAllService() {
	for {
		select {
		case data := <-h.BroadcastMessage:
			for _, v := range h.Group {
				for _, client := range v {
					client.Send <- data.Message
				}
			}
		}
	}
}
