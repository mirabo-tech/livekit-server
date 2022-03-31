package main

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v11"
	"strconv"
)

func initDbConnection() {

}

func connectClient() {
	//ctx := context.Background()
	//provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	//if err != nil {
	//	// handle error
	//}

	// Configure an OpenID Connect aware OAuth2 client.
	//oauth2Config := oauth2.Config{
	//	ClientID:     clientID,
	//	ClientSecret: clientSecret,
	//	RedirectURL:  redirectURL,
	//
	//	// Discovery returns the OAuth2 endpoints.
	//	Endpoint: provider.Endpoint(),
	//
	//	// "openid" is a required scope for OpenID Connect flows.
	//	Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	//}
}

func main() {
	var client = gocloak.NewClient("http://localhost:8080")
	ctx := context.Background()
	adminUsername := "phuclm"
	adminPassword := "password"
	clientId := "webrtc-service"
	grantType := "password"
	tokenOpt := gocloak.TokenOptions{
		Username:  &adminUsername,
		Password:  &adminPassword,
		ClientID:  &clientId,
		GrantType: &grantType}
	jwt, err := client.GetToken(ctx, "mirabo", tokenOpt)
	if err != nil {
		panic(err.Error())
	}
	token := jwt.AccessToken
	//result, err := client.RetrospectToken(ctx, token, "webrtc-service", "rSbAYgk1kp3fRbQA5RSiM6oJESHn0CGb", "mirabo")
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Print(result)

	for i := 0; i < 1000; i++ {
		username := "testUser" + strconv.Itoa(i)
		enableUser := true
		user := gocloak.User{Username: &username, Enabled: &enableUser}
		userId, err := client.CreateUser(ctx, token, "mirabo", user)
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Println(userId)
		err = client.SetPassword(ctx, token, userId, "mirabo", "password", false)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
}
