package handler

import (
	"time"

	"github.com/Bekzodbekk/paris2024_livestream_protos/genproto/livepb"
	"github.com/gin-gonic/gin"
)

func (h *HandlerST) CreateLiveStream(ctx *gin.Context) {

	eventId := ctx.Param("eventId")
	var requestBody struct {
		LeftSide  string            `json:"left_side"`
		RightSide string            `json:"right_side"`
		Actions   map[string]string `json:"actions"`
	}

	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	liveStream := livepb.LiveStream{
		EventId:   eventId,
		LeftSide:  requestBody.LeftSide,
		RightSide: requestBody.RightSide,
		Action:    requestBody.Actions,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	resp, err := h.Service.CreateLive(&liveStream)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, resp)
}

func (h *HandlerST) GetLiveStream(ctx *gin.Context) {

	eventId := ctx.Param("eventId")
	resp, err := h.Service.GetLive(&livepb.GetStreamRequest{
		Id: eventId,
	})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, resp)
}
