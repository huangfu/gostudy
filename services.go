package main

import (
	"log"

	"github.com/gogearbox/gearbox"
)

func main() {
	// Setup gearbox
	gb := gearbox.New()

	// create a logger middleware
	logMiddleware := func(ctx gearbox.Context) {
		log.Printf("log message!")

		// Next is what allows the request to continue to the next
		// middleware/handler
		ctx.Next()
	}

	// create an unauthorized middleware
	unAuthorizedMiddleware := func(ctx gearbox.Context) {
		ctx.Status(gearbox.StatusUnauthorized).SendString("You are unauthorized to access this page!")
	}

	// Register the log middleware for all requests
	gb.Use(logMiddleware)

	// Define your handlers
	gb.Get("/hello", func(ctx gearbox.Context) {
		ctx.SendString("Hello World!")
	})

	// Register the routes to be used when grouping routes
	routes := []*gearbox.Route{
		gb.Get("/id", func(ctx gearbox.Context) {
			ctx.SendString("User X")
		}),
		gb.Delete("/id", func(ctx gearbox.Context) {
			ctx.SendString("Deleted")
		}),
	}

	// Group account routes
	accountRoutes := gb.Group("/account", routes)

	// Group account routes to be under api
	gb.Group("/api", accountRoutes)

	// Define a route with unAuthorizedMiddleware as the middleware
	// you can define as many middlewares as you want and have
	// the handler as the last argument
	gb.Get("/protected", unAuthorizedMiddleware, func(ctx gearbox.Context) {
		ctx.SendString("You accessed a protected page")
	})

	// Start service
	gb.Start(":3000")
}