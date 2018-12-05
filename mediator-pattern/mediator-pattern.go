package mediator_pattern

import (
	"fmt"
	"time"
)

type User struct {
	Name string
}

func (u *User) GetUserName() string {
	return u.Name
}

func (u *User) SendMessage(chatRoom *ChatRoom, msg string) {
	chatRoom.ShowMessage(u, msg)
}

type ChatRoom struct {
}

func (c *ChatRoom) ShowMessage(user *User, msg string) {
	fmt.Printf("%v %v \n\t%v\n", time.Now().Format("2006-01-02 15:04:05"), user.GetUserName(), msg)
}
