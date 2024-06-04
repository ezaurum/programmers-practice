package main

// solution https://programmers.co.kr/learn/courses/30/lessons/42583
func solution(bridgeLength int, weight int, truckWeights []int) int {
	var passedTrucks []int
	completed := len(truckWeights)
	var bridge []int
	time := 0
	weightOnBridge := 0
	var timeOnBridge []int
	for len(passedTrucks) < completed {
		time++
		if len(bridge) > 0 && time-timeOnBridge[0] == bridgeLength {
			passedTruck := bridge[0]
			weightOnBridge -= passedTruck
			passedTrucks = append(passedTrucks, passedTruck)
			bridge = bridge[1:]
			timeOnBridge = timeOnBridge[1:]
		}
		if len(truckWeights) > 0 {
			nextTruck := truckWeights[0]
			// 무게 체크
			if weightOnBridge+nextTruck > weight {
				continue
			}

			//길이 체크
			if len(bridge) == bridgeLength {
				continue
			}
			bridge = append(bridge, nextTruck)
			timeOnBridge = append(timeOnBridge, time)
			weightOnBridge += nextTruck
			truckWeights = truckWeights[1:]
		}
	}
	return time
}
