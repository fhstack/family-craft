package hardware

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func SpinCatLock(angle int32) error {
	if angle < 0 || angle > 180 {
		return errors.New("illegal angle")
	}

	resp, err := http.Get("localhost:8888/" + strconv.FormatInt(int64(angle), 10))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("SpinCatLock failed, api status code: %d", resp.StatusCode)
	}

	return nil
}
