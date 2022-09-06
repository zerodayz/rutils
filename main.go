package main

import (
	"flag"
	"rutils/utils"
	"strings"
)

func init() {
	flag.StringVar(&utils.Counter, "countdown", "", "Creates an ASCII Art like count down clock. Accepts values like 1h15m30.918273645s")
	flag.StringVar(&utils.SpinningWheelNames, "spinning-wheel", "", "Spin the wheel like a boss. Accepts a comma separated list of names.")
	flag.Parse()
}

func main() {
	if utils.Counter != "" {
		utils.Count(utils.Counter)
	}
	if utils.SpinningWheelNames != "" {
		utils.SpinningWheelNamesList = strings.Split(utils.SpinningWheelNames, ",")
		utils.SpinningWheel()
	}
}
