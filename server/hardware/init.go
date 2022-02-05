package hardware

import (
	"github.com/stianeikeland/go-rpio/v4"
)

var catLockControlPin rpio.Pin

func Init() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}

	catLockControlPin = rpio.Pin(8)
	catLockControlPin.Pwm()
	catLockControlPin.Freq(1000)
}
