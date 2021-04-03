package main

import (
	"autobot/config"
	"autobot/strategy"
)

func main() {
	//Anthena??
	appConfig := config.Config{
		APIKey:    "Inkem Kavali??",
		APISecret: "Velithe Nalugu Matalu... Kuderthe Oka API Key.",
	}

	client := appConfig.SpawnKiteConnectClient()

	strategy.StrategyOne(client)
}
