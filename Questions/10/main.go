package main

import (
	"fmt"
)

func GetBucketId(f float64) int {
	return int(f/10) * 10
}
func main() {
	m := make(map[string][]float64)
	for _, v := range []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 0, -0.5, 5} {
		key := ""
		k := (GetBucketId(v))
		switch {
		case k == 0:
			if v >= 0 {
				key = fmt.Sprintf("+%d", k)
			} else {
				key = fmt.Sprintf("-%d", k)
			}
		default:
			key = fmt.Sprintf("%d", k)
		}
		bucket := m[key]
		bucket = append(bucket, v)
		m[key] = bucket
	}
	fmt.Println(m)
}
