package handler

import (
	"net/http"

	"github.com/adharshmk96/stk/gsk"
)

/*
PingHandler returns ping 200 response
Response:
- 200: OK
- 500: Internal Server Error
*/
func (h *pingHandler) PingHandler(gc *gsk.Context) {

	ping := h.service.PingService()

	gc.Status(http.StatusOK).JSONResponse(gsk.Map{
		"message": ping,
	})
}
