package main

import (
	"github.com/gofiber/fiber"
	"github.com/totalfish1/testapi/controllers"
)

func main() {
	app := fiber.New()

	app.Get("/api/jobs", controllers.GetAllJobListings)
	app.Get("/api/jobs/:id", controllers.GetJobListingID)
	app.Post("/api/jobs", controllers.CreateJobListing)
	app.Patch("/api/jobs/:id", controllers.UpdateJobListing)
	app.Delete("/api/jobs/:id", controllers.DeleteJobListing)

	err := app.Listen(3000)
	if err != nil {
		panic(err)
	}
}
