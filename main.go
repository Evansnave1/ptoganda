package main

import (
	"log"
	"net/http"
      // Framework for HTTP routing
	 "ptoganda/service" // Replace with your actual module name
)

// App struct holds references to services
type App struct {
	GoogleCalendarService *service.GoogleCalendarService
	PlandayService        *service.PlandayService
}

func main() {
	// Initialize services
	googleService := service.NewGoogleCalendarService(nil) // Pass Google API client here
	plandayService := service.NewPlandayService()

	// Create the App instance
	app := &App{
		GoogleCalendarService: googleService,
		PlandayService:        plandayService,
	}

	// Create a new Gin router
	router := gin.Default()

	// Define routes
	router.GET("/health", healthCheck)                      // Health check endpoint
	router.POST("/sync-google", app.syncGoogleCalendar)     // Sync schedules with Google Calendar
	router.GET("/schedule.ics", app.serveICSFile)           // Serve iCalendar file

	// Start the server
	log.Println("Server running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Health check endpoint
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// Sync Google Calendar endpoint
func (app *App) syncGoogleCalendar(c *gin.Context) {
	// Fetch schedules from Planday
	schedules, err := app.PlandayService.FetchSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedules from Planday"})
		return
	}

	// Sync schedules with Google Calendar
	err = app.GoogleCalendarService.SyncWithPlanday(schedules)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sync with Google Calendar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully synced schedules"})
}

// Serve iCalendar (.ics) file
func (app *App) serveICSFile(c *gin.Context) {
	icsData := app.PlandayService.GenerateICSFile()
	c.Header("Content-Type", "text/calendar")
	c.String(http.StatusOK, icsData)
}
