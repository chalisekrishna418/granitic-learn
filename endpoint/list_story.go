package endpoint

import (
	"context"

	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

//LSResponseBody struct defining the response body for list_story
type LSResponseBody struct {
	Story   string
	StoryID string
}

//LSRequestBody defines what to expect on request for list_story
type LSRequestBody struct {
	Log logging.Logger
}

//Process process the request
func (rb *LSRequestBody) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	res.Body = LSResponseBody{
		Story:   "some example story",
		StoryID: "CF3020-3131",
	}
	log := rb.Log

	log.LogTracef("Story Displayed")
}
