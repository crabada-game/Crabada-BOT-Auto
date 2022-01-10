package main

import (
	"Crabada/model"
	"Crabada/service"
	"fmt"
	"time"
)

func main() {
	s := service.NewCrabClient()
	address := s.PrivateToPublicKey(service.SecretKey)

	for {
		teams := &model.TeamList{}
		s.RestyClient.R().
			SetResult(&teams).
			Get("https://idle-api.crabada.com/public/idle/teams?user_address=" + address + "&page=1&limit=100")
		for _, team := range teams.Result.Data {
			if team.TeamID == service.TeamID {
				// unix to date
				fmt.Println("Game Finish at:", time.Unix(team.MineEndTime, 0), "Current Time:", time.Now())
				if team.MineEndTime < time.Now().Unix() {
					s.EndGame(team.GameID)
					s.StartGame(service.TeamID)
				}
			}
		}
		time.Sleep(10 * time.Second)
	}
}
