package mediator_pattern

import "testing"

func Test_MediatorPattern(t *testing.T) {
	chatRoom := new(ChatRoom)
	user1 := &User{Name: "kevin"}
	user2 := &User{Name: "john"}
	user3 := &User{Name: "lucy"}
	user4 := &User{Name: "robert"}
	user1.SendMessage(chatRoom, "hello everyone!")
	user2.SendMessage(chatRoom, "nice to meet you!")
	user3.SendMessage(chatRoom, "nice to meet you too!")
	user4.SendMessage(chatRoom, "good night!")
}
