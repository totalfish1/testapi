package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/totalfish1/testapi/models"
)

// GetAllJobListings scraps and return all job listings
func GetAllJobListings(ctx *fiber.Ctx) {
	jobs := []models.Joblisting{}
	ctx.JSON(fiber.Map{
		"ok":  true,
		"msg": "Working",
	})
}

// GetJobListingID gets a saved job lisitng from the db
func GetJobListingID(ctx *fiber.Ctx) {
	ctx.JSON(fiber.Map{
		"ok":  true,
		"msg": "Working",
	})
}

// UpdateJobListing updates a job listing from the db
func UpdateJobListing(ctx *fiber.Ctx) {
	ctx.JSON(fiber.Map{
		"ok":  true,
		"msg": "Working",
	})
}

// DeleteJobListing removes job listing from the db
func DeleteJobListing(ctx *fiber.Ctx) {
	ctx.JSON(fiber.Map{
		"ok":  true,
		"msg": "Working",
	})
}
