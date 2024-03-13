package main

import (
	"log"

	"github.com/song940/feishu-go/feishu"
)

func main() {
	client := feishu.NewClient(&feishu.Config{
		AppID:     "cli_a37e09c80539d00c",
		AppSecret: "cfgdqmY6IgHDk9Tim2Dg7cX5irFGArXt",
	})
	auth, _ := client.GetTenantAccessTokenInternal()
	// client.GetTenantAccessToken()
	// client.GetAppAccessToken()
	// client.GetAppAccessTokenInternal()
	// client.GetUserAccessToken()

	client.SetAccessToken(auth.TenantAccessToken)
	users, _ := client.GetUsersBy("open_id", []string{
		"18510100102",
	})
	for _, user := range users {
		log.Println(user.UserID, user.Mobile, user.Email)
	}
	// message := feishu.NewTextMessage("hello world")
	// message.ReceiveId = users[0].UserID
	// msgresp, _ := client.SendMessage(&message)
	// log.Println(msgresp)
	msgresp, err := client.SendTextMessage(users[0].UserID, "hello world!")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(msgresp)

}
