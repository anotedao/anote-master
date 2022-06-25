package main

import (
	"log"

	"github.com/mo7zayed/reqip"
	macaron "gopkg.in/macaron.v1"
)

func pingView(ctx *macaron.Context) {
	pr := &PingResponse{
		Success: true,
	}

	log.Printf("your ip is %s", reqip.GetClientIP(ctx.Req.Request))

	ctx.JSON(200, pr)
}

type PingResponse struct {
	Success bool `json:"success"`
}
