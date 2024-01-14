package helpers

import "fmt"

func GenerateRoomID(senderId, receiverId int) string {
	if senderId < receiverId {
		return fmt.Sprintf("%d_%d", senderId, receiverId)
	}

	return fmt.Sprintf("%d_%d", receiverId, senderId)
}
