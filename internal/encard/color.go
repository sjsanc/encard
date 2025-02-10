package encard

import (
	"fmt"
	"strconv"
)

func darkenHex(hex string, factor float64) string {
	// Parse the hex string as individual R, G, B components
	r, _ := strconv.ParseInt(hex[0:2], 16, 0)
	g, _ := strconv.ParseInt(hex[2:4], 16, 0)
	b, _ := strconv.ParseInt(hex[4:6], 16, 0)

	// Apply the darkening factor
	newR := int(float64(r) * factor)
	newG := int(float64(g) * factor)
	newB := int(float64(b) * factor)

	// Ensure values stay within [0, 255]
	newR = clamp(newR, 0, 255)
	newG = clamp(newG, 0, 255)
	newB = clamp(newB, 0, 255)

	return fmt.Sprintf("%02X%02X%02X", newR, newG, newB)
}

func getShades(base string, count int) []string {
	shades := []string{}
	for i := 0; i < count; i++ {
		factor := 1 - float64(i)/float64(count)
		shades = append(shades, darkenHex(base, factor))
	}
	return shades
}
