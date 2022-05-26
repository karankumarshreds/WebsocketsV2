package app

import "fmt"

type App struct{}

func  (a *App) InitializeApp() {
	fmt.Println("App initializing")
}