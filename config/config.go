package config

import (
	"context"
	kiteconnect "github.com/zerodhatech/gokiteconnect"
	"log"
	"net/http"
)

type Config struct {
	APIKey    string
	APISecret string
}

func (c *Config) SpawnKiteConnectClient() *kiteconnect.Client {
	kc := kiteconnect.New(c.APIKey)

	var (
		requestToken string
	)

	log.Println(kc.GetLoginURL())

	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/api/user/callback/kite/", func(w http.ResponseWriter, r *http.Request) {
		requestToken = r.URL.Query()["request_token"][0]
		go server.Shutdown(context.TODO())
		w.Write([]byte("login successful!"))
		return
	})
	server.ListenAndServe()

	userSession, err := kc.GenerateSession(requestToken, c.APISecret)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil
	}

	kc.SetAccessToken(userSession.AccessToken)

	return kc
}
