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
	angleStr := c.Query("angle")
	log.Printf("recv spin lock, angle: %s", angleStr)
	angle, err := strconv.ParseInt(angleStr, 10, 64)
	if err != nil {
		log.Printf("param illegal: %s", err)
		c.JSON(http.StatusOK, BuildParamIllegalResp(err.Error()))
		return
	}

	ok := catLock.TryLock()
	if !ok {
		c.JSON(http.StatusOK, BuildRespByCode(consts.ErrCode_Locking))
		return
	}
	defer catLock.Unlock()
	hardware.SpinAndResetCatLock(int32(angle))
	c.JSON(http.StatusOK, BuildSuccResp(nil))
}
