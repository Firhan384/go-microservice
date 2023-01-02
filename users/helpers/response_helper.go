package helpers

import "github.com/gofiber/fiber/v2"

type response struct {
	Data       any
	Valid      bool
	StatusCode uint16
	Messages   any
}

func ResponseJSON(c *fiber.Ctx, data any, messages any, statusCode uint16) response {
	var Response response

	if statusCode >= 400 && statusCode <= 499 {
		Response.Valid = false
	} else if statusCode >= 500 {
		Response.Valid = false
	} else {
		Response.Valid = true
	}

	return Response
}
