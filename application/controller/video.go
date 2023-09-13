package controller

import "github.com/gofiber/fiber/v2"

type VideoController struct{}

// NewVideoController - create new video handle instance
// Return:
//   - *VideoController
func NewVideoController() *VideoController {
	return &VideoController{}
}

// Upload - method to upload the video
// Return:
//   - error
func (vc *VideoController) Upload(c *fiber.Ctx) error {

	return nil
}
