package routers

import (
	"github.com/FedorLap2006/disgolf"
)

func Routers(dc *disgolf.Bot) {
	//Perolas Commands
	addPerolaRouter(dc)
	listPerolasRouter(dc)
	deletePerolaRouter(dc)

	//Ping Command
	pingRouter(dc)
}
