package main


func RateLimitation_I(requests []string) []string {
	responses := make([]string, 0)

	// Keeping request ith second for a particular domain into
	// a map with 2D array.
	// 0th index array contains the ith second for 5 second window and 
	// 1st index array contains the ith second for 30 second window. 
	requestCount := make(map[string][][]int, 0)

	// Initiating or updating map and slice for each domain
	for i, val := range requests {
		if _, ok := requestCount[val]; !ok {
			requestCount[val] = make([][]int, 0)
			requestCount[val] = append(requestCount[val], []int{}, []int{})
		}

		// Registering current second for each domain
		requestCount[val][0] = append(requestCount[val][0], i)
		requestCount[val][1] = append(requestCount[val][1], i)

		// Recalculating the seconds for 5 second window
		times_5 := make([]int, 0)
		for _, v := range requestCount[val][0] {
			if i-v < 5{
				times_5 = append(times_5, v)
			}
		}

		// Recalculating the seconds for 30 second window
		times_30 := make([]int, 0)
		for _, v := range requestCount[val][1] {
			if i-v < 30{
				times_30 = append(times_30, v)
			}
		}

		// Checking if it has crossed any limit on either of the windows
		// and adding respective response
		if len(times_5) > 2 || len(times_30) > 5 {
			responses = append(responses, "{status: 429, message: Too many requests}")
		} else {
			responses = append(responses, "{status: 200, message: OK}")
		}

		// clearing old data on 5 & 30 second window
		requestCount[val][0] = times_5
		requestCount[val][1] = times_30
	}

	return responses
}


// t for t in requests[val] if i-t < 5