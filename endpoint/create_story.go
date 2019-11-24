package endpoint

import (
	"context"

	"github.com/graniticio/granitic/v2/ws"
)

//CSRequestBody defines a request body for create_story
type CSRequestBody struct {
	Story string
}

//CSResponseBody defines response body for create_story
type CSResponseBody struct {
	status string
	RB     *CSRequestBody
}

//ProcessPayload processes the request
func (rb *CSRequestBody) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, request *CSRequestBody) {
	res.Body = CSResponseBody{
		status: "done",
		RB: &CSRequestBody{
			Story: request.Story,
		},
	}
}
