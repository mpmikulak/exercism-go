package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var carsPerHour float64 = float64(productionRate) * successRate / 100

	return int(carsPerHour) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var tens int = carsCount / 10
	var ones int = carsCount % 10

	var cost int = (tens * 95_000) + (ones * 10_000)
	return uint(cost)
}
