package main

import (
	"reflect"
	"testing"
)

func TestRateLimitation_I(t *testing.T) {
	test_data := [][]string{
		{
			"www.xyz.com",
			"www.abc.com",
			"www.xyz.com",
			"www.pqr.com",
			"www.abc.com",
			"www.xyz.com",
			"www.xyz.com",
		},
		{
			"www.xyz.com",
			"www.abc.com",
			"www.xyz.com",
			"www.xyz.com",
			"www.pqr.com",
			"www.abc.com",
			"www.abc.com",
			"www.abc.com",
			"www.xyz.com",
			"www.xyz.com",
		},
		{
			"www.xyz.com",
			"www.abc.com",
			"www.xyz.com",
			"www.xyz.com",
			"www.pqr.com",
			"www.abc.com",
			"www.abc.com",
			"www.abc.com",
			"www.xyz.com",
			"www.xyz.com",
			"www.xyz.com",
			"www.xyz.com",
		},
	}

	want := [][]string {
		{
			"{status: 200, message: OK}",
			"{status: 200, message: OK}",
			"{status: 200, message: OK}",
			"{status: 200, message: OK}",
			"{status: 200, message: OK}",
			"{status: 200, message: OK}",
			"{status: 429, message: Too many requests}",
		},
		{
			"{status: 200, message: OK}",
			"{status: 200, message: OK}",
			"{status: 200, message: OK}",
			"{status: 429, message: Too many requests}",
			"{status: 200, message: OK}",
			"{status: 200, message: OK}",
			"{status: 200, message: OK}",
			"{status: 429, message: Too many requests}",
			"{status: 200, message: OK}",
			"{status: 200, message: OK}",
		},
		{
			"{status: 200, message: OK}",
			"{status: 200, message: OK}",
			"{status: 200, message: OK}",
			"{status: 429, message: Too many requests}",
			"{status: 200, message: OK}",
			"{status: 200, message: OK}",
			"{status: 200, message: OK}",
			"{status: 429, message: Too many requests}",
			"{status: 200, message: OK}",
			"{status: 200, message: OK}",
			"{status: 429, message: Too many requests}",
			"{status: 429, message: Too many requests}",
		},
	}

	for i := range test_data {
		got := RateLimitation_I(test_data[i])
		if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("Test Failed for data: %v\noutput: %v\nexpected: %v", test_data[i], got, want[i])
		}
	}
}