package controller

import (
	"course-api/application/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UploadController struct {
	service *service.UploadService
}

// NewUploadController - create instance for uploader
//
// Return:
//   - *UploadController
func NewUploadController(service *service.UploadService) *UploadController {
	return &UploadController{
		service: service,
	}
}

// Upload - create upload handler
//
// Params:
//   - c *fiber.Ctx
//
// Return:
//   - error
func (uc *UploadController) UploadChunk(c *fiber.Ctx) error {
	formFile, err := c.FormFile("file")
	if err != nil {
		return failed(c, fiber.StatusInternalServerError, err.Error())
	}

	dirID := c.Params("file_id")
	chunkIndex, _ := strconv.Atoi(c.Params("index"))
	chunkTotal, _ := strconv.Atoi(c.Params("total"))
	uc.service.UploadChunk(formFile, dirID, chunkTotal, chunkIndex)

	return nil
}
