package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp string  `json:"timestamp"`
}

type Vehicle struct {
	ID               string       `json:"id"`
	LicensePlate     string       `json:"licensePlate"`
	VehicleType      string       `json:"vehicleType"`
	DriverName       string       `json:"driverName"`
	Latitude         float64      `json:"latitude"`
	Longitude        float64      `json:"longitude"`
	Speed            float64      `json:"speed"`
	Status           string       `json:"status"`
	IsOnline         bool         `json:"isOnline"`
	FuelLevel        float64      `json:"fuelLevel"`
	BatteryStatus    float64      `json:"batteryStatus"`
	TripCountToday   int          `json:"tripCountToday"`
	LastUpdate       string       `json:"lastUpdate"`
	Odometer         float64      `json:"odometer"`
	MaintenanceDue   bool         `json:"maintenanceDue"`
	LastServiceDate  string       `json:"lastServiceDate"`
	EngineHours      float64      `json:"engineHours"`
	RouteCoordinates []Coordinate `json:"routeCoordinates"`
}

var vehicles = []Vehicle{
	{ID: "VH-001", LicensePlate: "ABC123", VehicleType: "Truck", DriverName: "John Doe", Latitude: 37.7749, Longitude: -122.4194,
		Speed: 55.5, Status: "Moving", IsOnline: true, FuelLevel: 78, BatteryStatus: 0, TripCountToday: 5, LastUpdate: time.Now().Format(time.RFC3339),
		Odometer: 120000, MaintenanceDue: false, LastServiceDate: time.Now().AddDate(0, -3, 0).Format(time.RFC3339), EngineHours: 4500},
	{ID: "VH-002", LicensePlate: "XYZ789", VehicleType: "Van", DriverName: "Jane Smith", Latitude: 34.0522, Longitude: -118.2437,
		Speed: 0, Status: "Idle", IsOnline: true, FuelLevel: 50, BatteryStatus: 0, TripCountToday: 2, LastUpdate: time.Now().Add(-5 * time.Minute).Format(time.RFC3339),
		Odometer: 80000, MaintenanceDue: false, LastServiceDate: time.Now().AddDate(0, -6, 0).Format(time.RFC3339), EngineHours: 3200},
	{ID: "VH-003", LicensePlate: "LMN456", VehicleType: "Electric Car", DriverName: "Alice Johnson", Latitude: 40.7128, Longitude: -74.0060,
		Speed: 0, Status: "Stopped", IsOnline: false, FuelLevel: 0, BatteryStatus: 75, TripCountToday: 0, LastUpdate: time.Now().Add(-10 * time.Minute).Format(time.RFC3339),
		Odometer: 40000, MaintenanceDue: true, LastServiceDate: time.Now().AddDate(0, -12, 0).Format(time.RFC3339), EngineHours: 1500},
}

type Alert struct {
	VehicleID string `json:"vehicleId"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func init() {
	for i := range vehicles {
		vehicles[i].RouteCoordinates = []Coordinate{{
			Latitude:  vehicles[i].Latitude,
			Longitude: vehicles[i].Longitude,
			Timestamp: time.Now().Format(time.RFC3339),
		}}
	}
}

func main() {
	router := gin.Default()
	allowedOrigins := []string{"http://localhost:5173"}
	if os.Getenv("ENV") == "production" {
		allowedOrigins = []string{"https://gps-tracker-nine.vercel.app"}
	}
	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Go backend running 🚀"}) })
	router.GET("/api/vehicles", getVehicles)
	router.GET("/api/vehicles/:id", getVehicleByID)
	router.GET("/api/alerts", getAlerts)
	router.GET("/api/reports/daily-summary", getDailyReport)
	router.POST("/api/vehicles/random-update", randomUpdate)

	router.Run(":8080")
}

func getVehicles(c *gin.Context) {
	c.JSON(http.StatusOK, vehicles)
}

func getVehicleByID(c *gin.Context) {
	id := c.Param("id")
	for _, v := range vehicles {
		if v.ID == id {
			c.JSON(http.StatusOK, v)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Vehicle not found"})
}

func getAlerts(c *gin.Context) {
	alerts := []Alert{
		{"VH-001", "Low Fuel", time.Now().Add(-10 * time.Minute).Format(time.RFC3339)},
		{"VH-003", "Disconnected", time.Now().Add(-5 * time.Minute).Format(time.RFC3339)},
	}
	c.JSON(http.StatusOK, alerts)
}

func getDailyReport(c *gin.Context) {
	var moving, idle, stopped, maintenance, trips int
	var fuelSum, batterySum, speedSum, odometerSum float64
	var fuelCount, batteryCount, speedCount int

	for _, v := range vehicles {
		if v.Status == "Moving" {
			moving++
			speedSum += v.Speed
			speedCount++
		} else if v.Status == "Idle" {
			idle++
		} else if v.Status == "Stopped" {
			stopped++
		}

		if v.MaintenanceDue {
			maintenance++
		}

		if v.FuelLevel > 0 {
			fuelSum += v.FuelLevel
			fuelCount++
		}
		if v.VehicleType == "Electric Car" {
			batterySum += v.BatteryStatus
			batteryCount++
		}

		odometerSum += v.Odometer
		trips += v.TripCountToday
	}

	avgSpeed := 0.0
	if speedCount > 0 {
		avgSpeed = speedSum / float64(speedCount)
	}
	avgFuel := 0.0
	if fuelCount > 0 {
		avgFuel = fuelSum / float64(fuelCount)
	}
	avgBattery := 0.0
	if batteryCount > 0 {
		avgBattery = batterySum / float64(batteryCount)
	}

	summary := gin.H{
		"totalVehicles":       len(vehicles),
		"moving":              moving,
		"idle":                idle,
		"stopped":             stopped,
		"averageSpeed":        avgSpeed,
		"averageFuelLevel":    avgFuel,
		"averageBatteryLevel": avgBattery,
		"maintenanceVehicles": maintenance,
		"totalDistanceToday":  odometerSum,
		"totalTripsToday":     trips,
		"alertsCount":         2,
	}
	c.JSON(http.StatusOK, summary)
}

func randomUpdate(c *gin.Context) {
	for i := range vehicles {
		if vehicles[i].Status == "Moving" {
			vehicles[i].Latitude += (rand.Float64() - 0.5) * 0.01
			vehicles[i].Longitude += (rand.Float64() - 0.5) * 0.01
			vehicles[i].RouteCoordinates = append(vehicles[i].RouteCoordinates, Coordinate{
				Latitude:  vehicles[i].Latitude,
				Longitude: vehicles[i].Longitude,
				Timestamp: time.Now().Format(time.RFC3339),
			})
			vehicles[i].Speed = float64(rand.Intn(120))
			vehicles[i].LastUpdate = time.Now().Format(time.RFC3339)
			vehicles[i].Odometer += float64(rand.Intn(10))
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Vehicles updated"})
}
