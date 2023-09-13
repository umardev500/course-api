package controller

import (
	"course-api/application/service"
	"course-api/domain/model"
	"fmt"
	"os"
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
	finish, fileLocation, err := uc.service.UploadChunk(formFile, dirID, chunkTotal, chunkIndex)
	if err != nil {
		return failed(c, fiber.StatusInternalServerError, err.Error())
	}
	if !finish {
		return ok(c, fiber.StatusOK, "chunk uploaded", nil)
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	fileURL := fmt.Sprintf("%s%s/%s", host, port, *fileLocation)
	return ok(c, fiber.StatusOK, "file upload finished", model.UploadResponse{
		FileURL: fileURL,
	})
}
