package main

import (
	"math/rand"
	"net/http"
	"time"
	"strings" 
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
	{ID: "VH-001", LicensePlate: "PLT8474", VehicleType: "Truck", DriverName: "Michael Brown", Latitude: 40.747167, Longitude: -115.811004,
		Speed: 0, Status: "Stopped", IsOnline: true, FuelLevel: 51, BatteryStatus: 0, TripCountToday: 10,
		LastUpdate: "2025-05-12T23:26:30.752098", Odometer: 222886, MaintenanceDue: true, LastServiceDate: "2025-02-21T23:26:30.752110", EngineHours: 1617},
	{ID: "VH-002", LicensePlate: "PLT1600", VehicleType: "SUV", DriverName: "Alice Johnson", Latitude: 37.951317, Longitude: -121.565568,
		Speed: 0, Status: "Idle", IsOnline: false, FuelLevel: 7, BatteryStatus: 0, TripCountToday: 4,
		LastUpdate: "2025-05-12T23:26:30.752137", Odometer: 163400, MaintenanceDue: false, LastServiceDate: "2024-11-14T23:26:30.752143", EngineHours: 3401},
	{ID: "VH-003", LicensePlate: "PLT5019", VehicleType: "Van", DriverName: "Sarah Johnson", Latitude: 32.870367, Longitude: -124.161446,
		Speed: 0, Status: "Stopped", IsOnline: false, FuelLevel: 17, BatteryStatus: 0, TripCountToday: 0,
		LastUpdate: "2025-05-12T23:26:30.752164", Odometer: 38757, MaintenanceDue: true, LastServiceDate: "2024-10-14T23:26:30.752169", EngineHours: 3032},
	{ID: "VH-004", LicensePlate: "PLT2154", VehicleType: "Van", DriverName: "John Doe", Latitude: 35.894681, Longitude: -118.115521,
		Speed: 0, Status: "Idle", IsOnline: true, FuelLevel: 57, BatteryStatus: 0, TripCountToday: 7,
		LastUpdate: "2025-05-12T23:26:30.752183", Odometer: 249909, MaintenanceDue: true, LastServiceDate: "2024-10-01T23:26:30.752186", EngineHours: 5340},
	{ID: "VH-005", LicensePlate: "PLT2322", VehicleType: "Electric Car", DriverName: "Patricia Miller", Latitude: 32.743872, Longitude: -115.531382,
		Speed: 0, Status: "Stopped", IsOnline: false, FuelLevel: 0, BatteryStatus: 99, TripCountToday: 15,
		LastUpdate: "2025-05-12T23:26:30.752199", Odometer: 244591, MaintenanceDue: true, LastServiceDate: "2024-10-07T23:26:30.752202", EngineHours: 1630},
	{ID: "VH-006", LicensePlate: "PLT9027", VehicleType: "Van", DriverName: "Michael Brown", Latitude: 40.608112, Longitude: -117.083221,
		Speed: 44.9, Status: "Moving", IsOnline: false, FuelLevel: 94, BatteryStatus: 0, TripCountToday: 11,
		LastUpdate: "2025-05-12T23:26:30.752215", Odometer: 54674, MaintenanceDue: false, LastServiceDate: "2025-03-16T23:26:30.752219", EngineHours: 2196},
	{ID: "VH-007", LicensePlate: "PLT8177", VehicleType: "SUV", DriverName: "Chris Lee", Latitude: 34.952059, Longitude: -115.06603,
		Speed: 50.8, Status: "Moving", IsOnline: false, FuelLevel: 73, BatteryStatus: 0, TripCountToday: 14,
		LastUpdate: "2025-05-12T23:26:30.752240", Odometer: 153320, MaintenanceDue: true, LastServiceDate: "2024-11-05T23:26:30.752245", EngineHours: 1534},
	{ID: "VH-008", LicensePlate: "PLT3664", VehicleType: "Van", DriverName: "Michael Brown", Latitude: 41.222547, Longitude: -118.256997,
		Speed: 9.1, Status: "Moving", IsOnline: false, FuelLevel: 58, BatteryStatus: 0, TripCountToday: 5,
		LastUpdate: "2025-05-12T23:26:30.752271", Odometer: 140274, MaintenanceDue: false, LastServiceDate: "2024-06-04T23:26:30.752276", EngineHours: 2023},
	{ID: "VH-009", LicensePlate: "PLT7334", VehicleType: "Electric Car", DriverName: "Michael Brown", Latitude: 36.85547, Longitude: -122.196916,
		Speed: 0, Status: "Stopped", IsOnline: true, FuelLevel: 0, BatteryStatus: 87, TripCountToday: 7,
		LastUpdate: "2025-05-12T23:26:30.752289", Odometer: 241445, MaintenanceDue: false, LastServiceDate: "2024-12-03T23:26:30.752293", EngineHours: 3649},
	{ID: "VH-010", LicensePlate: "PLT8274", VehicleType: "Bus", DriverName: "John Doe", Latitude: 33.866489, Longitude: -116.930723,
		Speed: 0, Status: "Stopped", IsOnline: false, FuelLevel: 3, BatteryStatus: 0, TripCountToday: 20,
		LastUpdate: "2025-05-12T23:26:30.752317", Odometer: 116809, MaintenanceDue: false, LastServiceDate: "2024-09-04T23:26:30.752322", EngineHours: 2423},
	{ID: "VH-011", LicensePlate: "PLT3964", VehicleType: "Truck", DriverName: "Sarah Johnson", Latitude: 32.610076, Longitude: -124.227673,
		Speed: 0, Status: "Idle", IsOnline: false, FuelLevel: 44, BatteryStatus: 0, TripCountToday: 15,
		LastUpdate: "2025-05-12T23:26:30.752342", Odometer: 56636, MaintenanceDue: false, LastServiceDate: "2025-03-14T23:26:30.752345", EngineHours: 3028},
	{ID: "VH-012", LicensePlate: "PLT7886", VehicleType: "Sedan", DriverName: "David Wilson", Latitude: 35.902259, Longitude: -122.31561,
		Speed: 0, Status: "Stopped", IsOnline: true, FuelLevel: 22, BatteryStatus: 0, TripCountToday: 17,
		LastUpdate: "2025-05-12T23:26:30.752357", Odometer: 123674, MaintenanceDue: false, LastServiceDate: "2025-03-25T23:26:30.752360", EngineHours: 4673},
	{ID: "VH-013", LicensePlate: "PLT8644", VehicleType: "Van", DriverName: "Alice Johnson", Latitude: 33.109103, Longitude: -120.354669,
		Speed: 92.7, Status: "Moving", IsOnline: true, FuelLevel: 71, BatteryStatus: 0, TripCountToday: 17,
		LastUpdate: "2025-05-12T23:26:30.752372", Odometer: 33274, MaintenanceDue: true, LastServiceDate: "2024-10-04T23:26:30.752377", EngineHours: 5326},
	{ID: "VH-014", LicensePlate: "PLT4521", VehicleType: "Sedan", DriverName: "Daniel Garcia", Latitude: 36.487841, Longitude: -115.968395,
		Speed: 0, Status: "Idle", IsOnline: false, FuelLevel: 55, BatteryStatus: 0, TripCountToday: 10,
		LastUpdate: "2025-05-12T23:26:30.752396", Odometer: 179589, MaintenanceDue: false, LastServiceDate: "2025-02-19T23:26:30.752401", EngineHours: 3585},
	{ID: "VH-015", LicensePlate: "PLT6083", VehicleType: "Sedan", DriverName: "Patricia Miller", Latitude: 41.684312, Longitude: -123.093613,
		Speed: 0, Status: "Idle", IsOnline: true, FuelLevel: 81, BatteryStatus: 0, TripCountToday: 18,
		LastUpdate: "2025-05-12T23:26:30.752415", Odometer: 121341, MaintenanceDue: true, LastServiceDate: "2024-09-12T23:26:30.752419", EngineHours: 4301},
	{ID: "VH-016", LicensePlate: "PLT6371", VehicleType: "Sedan", DriverName: "Michael Brown", Latitude: 33.732895, Longitude: -115.574941,
		Speed: 17.3, Status: "Moving", IsOnline: false, FuelLevel: 83, BatteryStatus: 0, TripCountToday: 7,
		LastUpdate: "2025-05-12T23:26:30.752437", Odometer: 87515, MaintenanceDue: true, LastServiceDate: "2024-09-24T23:26:30.752440", EngineHours: 5371},
	{ID: "VH-017", LicensePlate: "PLT2552", VehicleType: "Electric Car", DriverName: "Sophia Hernandez", Latitude: 39.00641, Longitude: -118.404233,
		Speed: 55.8, Status: "Moving", IsOnline: true, FuelLevel: 0, BatteryStatus: 5, TripCountToday: 3,
		LastUpdate: "2025-05-12T23:26:30.752452", Odometer: 159601, MaintenanceDue: false, LastServiceDate: "2024-08-11T23:26:30.752455", EngineHours: 4774},
	{ID: "VH-018", LicensePlate: "PLT8756", VehicleType: "Electric Car", DriverName: "Chris Lee", Latitude: 34.588226, Longitude: -121.742025,
		Speed: 0, Status: "Idle", IsOnline: false, FuelLevel: 0, BatteryStatus: 16, TripCountToday: 10,
		LastUpdate: "2025-05-12T23:26:30.752467", Odometer: 213317, MaintenanceDue: true, LastServiceDate: "2024-06-16T23:26:30.752470", EngineHours: 2232},
	{ID: "VH-019", LicensePlate: "PLT9657", VehicleType: "Bus", DriverName: "Olivia Lewis", Latitude: 33.127552, Longitude: -114.44253,
		Speed: 0, Status: "Idle", IsOnline: true, FuelLevel: 29, BatteryStatus: 0, TripCountToday: 16,
		LastUpdate: "2025-05-12T23:26:30.752481", Odometer: 101619, MaintenanceDue: false, LastServiceDate: "2024-05-27T23:26:30.752484", EngineHours: 3542},
	{ID: "VH-020", LicensePlate: "PLT2433", VehicleType: "Truck", DriverName: "Patricia Miller", Latitude: 40.823735, Longitude: -120.448798,
		Speed: 0, Status: "Idle", IsOnline: false, FuelLevel: 78, BatteryStatus: 0, TripCountToday: 20,
		LastUpdate: "2025-05-12T23:26:30.752496", Odometer: 200470, MaintenanceDue: false, LastServiceDate: "2024-11-19T23:26:30.752501", EngineHours: 5961},
	{ID: "VH-021", LicensePlate: "PLT7876", VehicleType: "Truck", DriverName: "Emily Davis", Latitude: 33.75938, Longitude: -124.340372,
		Speed: 45.2, Status: "Moving", IsOnline: false, FuelLevel: 53, BatteryStatus: 0, TripCountToday: 6,
		LastUpdate: "2025-05-12T23:26:30.752520", Odometer: 61657, MaintenanceDue: false, LastServiceDate: "2025-01-24T23:26:30.752525", EngineHours: 1734},
	{ID: "VH-022", LicensePlate: "PLT9906", VehicleType: "Truck", DriverName: "David Wilson", Latitude: 40.824287, Longitude: -121.680886,
		Speed: 0, Status: "Idle", IsOnline: true, FuelLevel: 18, BatteryStatus: 0, TripCountToday: 5,
		LastUpdate: "2025-05-12T23:26:30.752543", Odometer: 218573, MaintenanceDue: true, LastServiceDate: "2025-01-08T23:26:30.752546", EngineHours: 1840},
	{ID: "VH-023", LicensePlate: "PLT8939", VehicleType: "Electric Car", DriverName: "Sarah Johnson", Latitude: 41.948249, Longitude: -119.438348,
		Speed: 0, Status: "Stopped", IsOnline: false, FuelLevel: 0, BatteryStatus: 80, TripCountToday: 12,
		LastUpdate: "2025-05-12T23:26:30.752556", Odometer: 219244, MaintenanceDue: true, LastServiceDate: "2024-05-31T23:26:30.752559", EngineHours: 3705},
	{ID: "VH-024", LicensePlate: "PLT7107", VehicleType: "Bus", DriverName: "Anthony Clark", Latitude: 37.643223, Longitude: -119.874505,
		Speed: 6.0, Status: "Moving", IsOnline: true, FuelLevel: 53, BatteryStatus: 0, TripCountToday: 6,
		LastUpdate: "2025-05-12T23:26:30.752577", Odometer: 185767, MaintenanceDue: false, LastServiceDate: "2024-12-12T23:26:30.752581", EngineHours: 5458},
	{ID: "VH-025", LicensePlate: "PLT5570", VehicleType: "Electric Car", DriverName: "James Rodriguez", Latitude: 41.104422, Longitude: -120.05865,
		Speed: 0, Status: "Idle", IsOnline: true, FuelLevel: 0, BatteryStatus: 78, TripCountToday: 17,
		LastUpdate: "2025-05-12T23:26:30.752599", Odometer: 213603, MaintenanceDue: false, LastServiceDate: "2024-10-29T23:26:30.752604", EngineHours: 3733},
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
	router.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5173" ||
				   origin == "https://gps-tracker-nine.vercel.app" ||  // your prod url
				   strings.HasSuffix(origin, ".vercel.app")            // allow all preview urls
		},
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
