package routes

import (
	"time"

	"github.com/Ciggzy1312/url-shortener/helpers"

	"github.com/gofiber/fiber"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"custom"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"custom"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func shortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "cannot parse JSON"})
	}

	// check domain error

	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"err": "Domain error"})
	}

}
