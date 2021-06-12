package config

import rl "github.com/gen2brain/raylib-go/raylib"

type Config struct {
	WindowTitle           string
	WindowWidth           int32
	WindowHeight          int32
	WindowFPS             int32
	WindowBackgroundColor rl.Color
}

func GetConfig() Config {
	return Config{
		WindowTitle:           "TicTac-Go",
		WindowWidth:           800,
		WindowHeight:          600,
		WindowFPS:             60,
		WindowBackgroundColor: rl.DarkGray,
	}
}
