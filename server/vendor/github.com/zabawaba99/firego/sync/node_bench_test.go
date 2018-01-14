package sync

import (
	"testing"
)

func BenchmarkNewNode(b *testing.B) {
	nodesToCreate := []interface{}{
		// primitives
		"somestirng",
		1,
		int8(2),
		int16(12),
		int32(222),
		int64(3245),
		float32(123123123123234),
		float64(7654323456543245654),
		true,

		// slice
		[]interface{}{
			"foobar",
			1234,
			true,
		},

		// simple map
		map[string]interface{}{
			"hello":  "world",
			"foo":    false,
			"number": 222,
		},

		// something that is not interface type
		map[string]bool{
			"1": true,
			"2": true,
			"3": true,
		},

		// nest map
		map[string]interface{}{
			"aaa": []interface{}{
				"ahhhh",
				map[string]interface{}{
					"wat?": "another map?!",
					"yup":  true,
				},
			},
			"doras-backpack": map[string]interface{}{
				"si se puede":     "yes you can!",
				"map map map map": 333,
			},
			"something-simple": 1,
			"inception": map[string]interface{}{
				"aaa": []interface{}{
					"ahhhh",
					map[string]interface{}{
						"wat?": "another map?!",
						"yup":  true,
					},
				},
				"doras-backpack": map[string]interface{}{
					"si se puede":     "yes you can!",
					"map map map map": 333,
				},
				"something-simple": 1,
				"inception": map[string]interface{}{
					"aaa": []interface{}{
						"ahhhh",
						map[string]interface{}{
							"wat?": "another map?!",
							"yup":  true,
						},
					},
					"doras-backpack": map[string]interface{}{
						"si se puede":     "yes you can!",
						"map map map map": 333,
					},
					"something-simple": 1,
					"inception": map[string]interface{}{
						"aaa": []interface{}{
							"ahhhh",
							map[string]interface{}{
								"wat?": "another map?!",
								"yup":  true,
							},
						},
						"doras-backpack": map[string]interface{}{
							"si se puede":     "yes you can!",
							"map map map map": 333,
						},
						"something-simple": 1,
					},
				},
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, node := range nodesToCreate {
			_ = NewNode("", node)
		}
	}
}
