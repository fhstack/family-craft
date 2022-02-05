package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fhstack/family-craft/server/consts"
	"github.com/fhstack/family-craft/server/hardware"
	"github.com/fhstack/family-craft/server/util"
	"github.com/gin-gonic/gin"
)

var catLock = util.NewLock()

func CatLockSpin(c *gin.Context) {
	angleStr := c.Param("angle")
	log.Printf("recv spin lock, angle: %s", angleStr)
	angle, err := strconv.ParseInt(angleStr, 10, 64)
	if err != nil {
		log.Printf("param illegal: %s", err)
		c.JSON(http.StatusOK, BuildRespByErr(err))
		return
	}

	ok := catLock.TryLock()
	if !ok {
		c.JSON(http.StatusOK, &Response{
			Code:    consts.ErrCode_Locking,
			Message: consts.ErrCode_Locking_Msg,
		})
		return
	}
	defer catLock.Unlock()
	hardware.SpinAndResetCatLock(int32(angle))
	c.JSON(http.StatusOK, &Response{})
}
