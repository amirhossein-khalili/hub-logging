package handlers

import (
	"hub_logging/external/api/rest"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type LogHandler struct {
}

func SetupLogRoute(rh *rest.RestHandler) {

	app := rh.App

	//create instance user service & inject to handler

	handler := LogHandler{}

	/* ---------------------------------------------------------------
					PUBLIC ENDPOINTS
	*----------------------------------------------------------------/
	*/
	app.Get("", handler.Logs)

	/* ---------------------------------------------------------------
					PRIVATE ENDPOINTS
	*----------------------------------------------------------------/
	*/

}

func (h *LogHandler) Logs(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "register"})
}
