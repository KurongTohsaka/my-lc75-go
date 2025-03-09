package main

func main() {

}

func largestAltitude(gain []int) int {
	maxAltitude := 0
	altitude := 0

	for _, v := range gain {
		altitude += v
		maxAltitude = max(maxAltitude, altitude)
	}

	return maxAltitude
}
