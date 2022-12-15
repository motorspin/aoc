package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// This code is cleaned up but the original solutions for each part are still
// represented.
//
// Solutions Attempted (Part 1):
//  - Naive, (i*j) for each sensor where sensor.x/y - sensor.distance < i/j <
//  sensor.x/y + sensor.distance, annotate each coordinate with a Cell value in
//  a sparse gri (map)d, skip rows where mDistance of new i,j coord is greater
//  than sensor distance (not optimal)
//  - Realize that only a specific row is of interest, don't bother inputting
//  anything into the sparse grid
//
// Solutions Attempted (Part 2):
//  - Optimization for Part 1 won't work, go back to Naive approach and see if
//  we can optimize some edge cases and still get a sparseGrid to then run
//  through (the answer is: this is definitely not the way)
//  - Realize we can save on some computations because when we can reduce the
//  range of x based on the current value of y and manhattan distance (see part
//  1 naive solution for what we were doing previously)
//  - Generate ranges for each sensor along the y-axis, sort them by their
//  start, and then iterate thru the x/y coordinates (lStart, lEnd) from the
//  problem statement to figure out where x is covered

type Cell int

const (
	Empty Cell = iota
	Sensor
	Beacon
	NoBeacon
)

type Coord struct {
	x int
	y int
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}

	return val
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func mDistance(point1 Coord, point2 Coord) int {
	return abs(point1.x-point2.x) + abs(point1.y-point2.y)
}

type SensorT struct {
	x        int
	y        int
	distance int
	beaconX  int
	beaconY  int
}

type Range struct {
	start int
	end   int
}

func solvePart1(sensors []SensorT, rowOfInterest int) int {
	count := 0
	cave := make(map[Coord]Cell, 0)

	for _, sensor := range sensors {
		sensorCoord := Coord{sensor.x, sensor.y}
		beaconCoord := Coord{sensor.beaconX, sensor.beaconY}
		cave[sensorCoord] = Sensor
		cave[beaconCoord] = Beacon

		for x := sensor.x - sensor.distance; x <= sensor.x+sensor.distance; x++ {
			coord := Coord{x, rowOfInterest}
			distance := mDistance(Coord{sensor.x, sensor.y}, coord)

			if distance > sensor.distance {
				continue
			}

			if cave[coord] == Empty {
				cave[coord] = NoBeacon
				count += 1
			}
		}
	}

	return count
}

func solvePart2(sensors []SensorT, lStart int, lEnd int) int {
	var location Coord
	ranges := make(map[int][]Range, 0)

	for y := lStart; y <= lEnd; y++ {
		for _, sensor := range sensors {
			yDistanceDiff := abs(sensor.y - y)

			// Can the current sensor reach this y?
			if sensor.distance-yDistanceDiff < 0 {
				continue
			}

			xDistanceDiff := abs(sensor.distance - yDistanceDiff)
			xStart, xEnd := sensor.x-xDistanceDiff, sensor.x+xDistanceDiff

			if xStart <= lStart && xEnd >= lEnd {
				// It can't be this line, it's full
				continue
			}

			ranges[y] = append(ranges[y], Range{max(xStart, lStart), min(xEnd, lEnd)})
		}
	}

	for y := range ranges {
		sort.Slice(ranges[y], func(i int, j int) bool {
			return ranges[y][i].start < ranges[y][j].start
		})
	}

mainloop:
	for y := lStart; y <= lEnd; y++ {
		for x := lStart; x <= lEnd; x++ {
			for _, r := range ranges[y] {
				if x < r.start {
					location = Coord{x + 1, y}
					break mainloop
				}

				// Move x forward by the entire range
				x = max(x, r.end)
			}
		}
	}

	return 4000000*location.x + location.y
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	rowOfInterest := 2000000
	lStart, lEnd := 0, 4000000
	isSample := false
	sensors := make([]SensorT, 0)

	if isSample {
		rowOfInterest = 10
		lEnd = 20
	}

	for scanner.Scan() {
		var sensor, beacon Coord
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.x, &sensor.y, &beacon.x, &beacon.y)
		distance := mDistance(sensor, beacon)
		sensors = append(sensors, SensorT{sensor.x, sensor.y, distance, beacon.x, beacon.y})
	}

	fmt.Println("Part 1:", solvePart1(sensors, rowOfInterest))
	fmt.Println("Part 2:", solvePart2(sensors, lStart, lEnd))
}
