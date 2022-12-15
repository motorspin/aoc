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

// After the solution was submitted, we can still do a lot better:
//  - Changed solvePart2: Change when we generate the ranges so we can save
//  on iterations over y (and memory)
//  - Changed solvePart1: Like part 2, use ranges for the row. Also consolidate
//  part1 and part2 range logic

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

func consolidateOverlaps(in []Range) (out []Range) {
	if len(in) <= 1 {
		return in
	}

	out = make([]Range, 0)
	out = append(out, in[0])
	current := 0

	for i := 1; i < len(in); i++ {
		// No overlap
		if in[i].start > out[current].end {
			out = append(out, in[i])
			current = len(out) - 1
			continue
		}

		// There is overlap, modify the end of our current out
		out[current].end = max(out[current].end, in[i].end)
	}

	return
}

func getRanges(sensors []SensorT, row int) (ranges []Range) {
	ranges = make([]Range, 0)

	for _, sensor := range sensors {
		yDistanceDiff := abs(sensor.y - row)

		// Can the current sensor reach this y?
		if sensor.distance-yDistanceDiff < 0 {
			continue
		}

		xDistanceDiff := abs(sensor.distance - yDistanceDiff)
		xStart, xEnd := sensor.x-xDistanceDiff, sensor.x+xDistanceDiff
		ranges = append(ranges, Range{xStart, xEnd})
	}

	sort.Slice(ranges, func(i int, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	ranges = consolidateOverlaps(ranges)

	return ranges
}

func solvePart1(sensors []SensorT, rowOfInterest int) int {
	count := 0
	ranges := getRanges(sensors, rowOfInterest)

	for _, r := range ranges {
		count += r.end - r.start
	}

	return count
}

func solvePart2(sensors []SensorT, lStart int, lEnd int) int {
	var location Coord

	for y := lStart; y <= lEnd; y++ {
		ranges := getRanges(sensors, y)

		if len(ranges) > 1 {
			location.y = y

			if ranges[0].start > 0 {
				break
			}

			location.x = ranges[1].start - 1
			break
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
