package utils

import (
	"fmt"
	figlet "github.com/mbndr/figlet4go"
	"os"
	"time"
)

var SpinningWheelNames string
var SpinningWheelNamesList []string
var b bool
var SelectedName string

func RemoveOldName(items []string, item string) []string {
	var newNames []string
	for _, i := range items {
		if i != item {
			newNames = append(newNames, i)
		}
	}
	return newNames
}

func SpinningWheel() {
	go func() {
		for {
			fmt.Scanln()
			b = !b
		}
	}()
	for {
		select {
		default:
			for _, r := range SpinningWheelNamesList {
				// Clear the screen
				print("\033[H\033[2J")
				if b {
					SpinningWheelNamesList = RemoveOldName(SpinningWheelNamesList, SelectedName)
					fmt.Println("Press the Enter Key to continue spinning the wheel.")
					out := Figlet(SelectedName, figlet.ColorGreen)
					for _, line := range out {
						fmt.Println(line)
					}
					fmt.Println("Still in the running:", SpinningWheelNamesList)
				} else {
					if len(SpinningWheelNamesList) == 1 {
						fmt.Println("This was the last name in the list.")
						out := Figlet(SpinningWheelNamesList[0], figlet.ColorGreen)
						for _, line := range out {
							fmt.Println(line)
						}
						fmt.Println("Thank you for playing.")
						os.Exit(0)
					}
					fmt.Println("Press the Enter Key to stop spinning the wheel.")
					out := Figlet(r, figlet.ColorRed)
					for _, line := range out {
						fmt.Println(line)
					}
					SelectedName = r
				}
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
