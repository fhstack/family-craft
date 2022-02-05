package hardware

import (
	"errors"
	"os/exec"
	"strconv"
)

func SpinCatLock(angle int32) error {
	if angle < 0 || angle > 180 {
		return errors.New("illegal angle")
	}

	args := []string{"cat_lock.py", strconv.FormatInt(int64(angle), 10)}
	err := exec.Command("python", args...).Run()
	if err != nil {
		return err
	}

	return nil
}
