package elon

import (
	"fmt"
	"math"
)

// Drive updates the number of meters driven based on the car's speed, and
// reduces the battery according to the battery drainage.
func (c *Car) Drive() {
	if c.battery-c.batteryDrain >= 0 {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}
}

// DisplayDistance returns the distance as displayed on the LED display as a
// string.
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery returns the battery percentage as displayed on the LED
// display as a string.
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish takes a trackDistance int as its parameter and returns whether the
// car can finish the race before draining its battery.
func (c *Car) CanFinish(trackDistance int) bool {
	numberOfDrives := int(math.Ceil(float64(trackDistance) / float64(c.speed)))
	batteryNeeded := numberOfDrives * c.batteryDrain
	return c.battery >= batteryNeeded
}
