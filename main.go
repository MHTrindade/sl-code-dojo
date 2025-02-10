package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Person represents an individual with a name, age, and a list of contacts.
// The embedded gorm.Model provides fields such as ID, CreatedAt, UpdatedAt, and DeletedAt.
type Person struct {
	gorm.Model
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Contacts []Contact
}

// Contact represents a contact method for a Person.
type Contact struct {
	gorm.Model
	Type     string `json:"type"`
	Value    string `json:"value"`
	PersonID uint
}

func main() {
	// Data Source Name (DSN) for connecting to the MySQL database.
	// Format: username:password@tcp(host:port)/database?options
	dsn := "root:root@tcp(127.0.0.1:3306)/sl_dojo?charset=utf8mb4&parseTime=True&loc=Local"

	// Open a connection to the MySQL database using GORM.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Automatically migrate the schema, creating or updating the tables.
	db.AutoMigrate(
		&Person{},
		&Contact{},
	)

	// Create a new Fiber app instance to handle HTTP requests.
	app := fiber.New()

	// Define a GET route for the root path ("/") that returns a welcome message.
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// Define a POST route at "/people" to create a new Person along with their contacts.
	// Expects a JSON payload matching the structure of the Person model.
	app.Post("people", func(c *fiber.Ctx) error {
		var p Person

		// Parse the incoming JSON request body into the Person struct.
		if err := c.BodyParser(&p); err != nil {
			// If parsing fails, return a 400 Bad Request status with an error message.
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
		}

		// Save the new Person record to the database.
		// This will also persist any associated contacts provided in the JSON payload.
		db.Create(&p)
		// TODO: Implement proper error handling for the db.Create(&p) call.

		// Return the created Person record as JSON with an HTTP status code 201 (Created).
		return c.Status(fiber.StatusCreated).JSON(p)
	})

	// Start the Fiber web server, listening for incoming HTTP requests on port 3000.
	// If an error occurs while starting the server, log.Fatal will log the error and terminate the application.
	log.Fatal(app.Listen(":3000"))
}
