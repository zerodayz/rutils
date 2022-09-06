package utils

import (
	"fmt"
	figlet "github.com/mbndr/figlet4go"
	"strings"
	"time"
)

var Counter string

func Figlet(msg string, color figlet.Color) []string {
	ascii := figlet.NewAsciiRender()

	options := figlet.NewRenderOptions()
	options.FontName = "standard"
	options.FontColor = []figlet.Color{color}

	renderStr, _ := ascii.RenderOpts(msg, options)
	out := strings.Split(fmt.Sprintf("%v", renderStr), "\n")
	return out
}

func Count(timeSleep string) {
	d, err := time.ParseDuration(timeSleep)
	if err != nil {
		fmt.Println("unable to parse duration.")
		return
	}
	for i := int(d.Seconds()); i > 0; i-- {
		// Clear the screen
		print("\033[H\033[2J")
		timeLeft := time.Duration(i) * time.Second

		h := int(timeLeft.Hours())
		m := int(timeLeft.Minutes()) - (h * 60)
		s := int(timeLeft.Seconds()) - (m * 60) - (h * 3600)

		asciiClock := fmt.Sprintf("%02d:%02d:%02d", h, m, s)

		out := Figlet(asciiClock, figlet.ColorGreen)
		for _, line := range out {
			fmt.Println(line)
		}

		time.Sleep(time.Second)

	}
	print("\033[H\033[2J")
	out := Figlet("Times up!", figlet.ColorRed)
	for _, line := range out {
		fmt.Println(line)
	}
}
