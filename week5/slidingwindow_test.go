package sliding

import (
	"fmt"
	"testing"
	"time"
)

func TestMax(t *testing.T) {
	testCase := []struct {
		winTime    time.Duration
		bucketTime time.Duration
		data       []float64

		want float64
	}{
		{10 * time.Second, 1 * time.Second, []float64{10, 11, 9, 8, 12, 7, 6}, 12},
		{2 * time.Second, 2 * time.Second, []float64{10, 11, 9, 8, 12, 7, 6}, 12},
	}

	for _, tc := range testCase {
		t.Run(fmt.Sprintf("winTime: %+v, bucketTime: %+v", tc.winTime, tc.bucketTime), func(t *testing.T) {
			n := NewNumber(tc.winTime, tc.bucketTime)
			for _, x := range tc.data {
				n.UpdateMax(x)
				time.Sleep(1 * time.Second)
			}

			if got := n.Max(time.Now()); got != tc.want {
				t.Errorf("got %f, want %f", got, tc.want)
			}
		})
	}
}

func TestAvg(t *testing.T) {
	testCase := []struct {
		winTime    time.Duration
		bucketTime time.Duration
		data       []float64

		want float64
	}{
		{10 * time.Second, 1 * time.Second, []float64{0.5, 1.5, 2.5, 3.5, 4.5}, 1.25},
		{5 * time.Second, 500 * time.Millisecond, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5.5},
		{5 * time.Second, 500 * time.Millisecond, []float64{1, 2, 3, 4, 5}, 1.5},
	}

	for _, tc := range testCase {
		t.Run(fmt.Sprintf("winTime: %+v, bucketTime: %+v", tc.winTime, tc.bucketTime), func(t *testing.T) {
			n := NewNumber(tc.winTime, tc.bucketTime)
			for _, x := range tc.data {
				n.Increment(x)
				time.Sleep(tc.bucketTime)
			}

			if got := n.Avg(time.Now()); got != tc.want {
				t.Errorf("got %f, want %f", got, tc.want)
			}
		})
	}
}
