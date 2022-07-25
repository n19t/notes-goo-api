package fbauth

import (
	"context"
	"notes-goo-api/internal/env"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var (
	firebaseConfigFile = env.GetEnvVariable("FIREBASE_CONFIG_FILE")
)

func InitAuth() (*auth.Client, error) {
	opt := option.WithCredentialsFile(firebaseConfigFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	client, errAuth := app.Auth(context.Background())
	if errAuth != nil {
		return nil, errAuth
	}
	return client, nil
}
