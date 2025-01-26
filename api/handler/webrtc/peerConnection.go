package webrtc

import (
	"boilerplate/lib"
	"boilerplate/lib/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (h *ConnectHandler) StartCall(c *fiber.Ctx) error {
	request := dto.CallInfo{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	answer, err := h.startCall.Execute(h.peerConnection, request)
	if err != nil {
		log.Errorf("Error starting call: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not making webrtc call, create answer error.",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.Sdp{Sdp: lib.Encode(answer)})
}
