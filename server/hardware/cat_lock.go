package hardware

import "time"

func SpinAndResetCatLock(angle int32) {
	for i := 0; i < 50; i++ {
		catLockControlPin.DutyCycle(uint32(angle)/270*100+25, 1000)
		time.Sleep(time.Millisecond * 20)
	}
}
