package main

import "testing"

func TestRapper_Distance(t *testing.T) {
	tests := []struct {
		rapper     Rapper
		sen1, sen2 Sentence
		distance   float64
	}{
		{
			Rapper{
				weights: []Weight{
					{1.0, 10.0},
					{2.0, 20.0},
					{3.0, 30.0},
				},
				maxWeight: 1.0 + 2.0 + 3.0 + 10.0 + 20.0 + 30.0,
			},
			Sentence{
				&Morph{"", "", "", "", "", "", "", "", "", "アイ"},
				&Morph{"", "", "", "", "", "", "", "", "", "ウ"},
				&Morph{"", "", "", "", "", "", "", "", "", "エオ"},
			},
			Sentence{
				&Morph{"", "", "", "", "", "", "", "", "", "カキクケ"},
				&Morph{"", "", "", "", "", "", "", "", "", "コ"},
			},
			(10.0 + 20.0 + 30.0) / (1.0 + 2.0 + 3.0 + 10.0 + 20.0 + 30.0),
		},
		{
			Rapper{
				weights: []Weight{
					{1.0, 10.0},
					{2.0, 20.0},
					{3.0, 30.0},
				},
				maxWeight: 1.0 + 2.0 + 3.0 + 10.0 + 20.0 + 30.0,
			},
			Sentence{
				&Morph{"", "", "", "", "", "", "", "", "", "アイ"},
				&Morph{"", "", "", "", "", "", "", "", "", ""},
				&Morph{"", "", "", "", "", "", "", "", "", "エオ"},
			},
			Sentence{
				&Morph{"", "", "", "", "", "", "", "", "", "カキクケ"},
				&Morph{"", "", "", "", "", "", "", "", "", "コ"},
			},
			0.0,
		},
		{
			Rapper{
				weights: []Weight{
					{1.0, 10.0},
					{2.0, 20.0},
					{3.0, 30.0},
				},
				maxWeight: 1.0 + 2.0 + 3.0 + 10.0 + 20.0 + 30.0,
			},
			Sentence{
				&Morph{"", "", "", "", "", "", "", "", "", "エオ"},
			},
			Sentence{
				&Morph{"", "", "", "", "", "", "", "", "", "カキクケ"},
				&Morph{"", "", "", "", "", "", "", "", "", "コ"},
			},
			(20.0 + 30.0) / (1.0 + 2.0 + 3.0 + 10.0 + 20.0 + 30.0),
		},
		{
			Rapper{
				weights: []Weight{
					{3.0, 30.0},
				},
				maxWeight: 3.0 + 30.0,
			},
			Sentence{
				&Morph{"", "", "", "", "", "", "", "", "", "アイ"},
				&Morph{"", "", "", "", "", "", "", "", "", "ウ"},
				&Morph{"", "", "", "", "", "", "", "", "", "エオ"},
			},
			Sentence{
				&Morph{"", "", "", "", "", "", "", "", "", "カキクケ"},
				&Morph{"", "", "", "", "", "", "", "", "", "コ"},
			},
			30.0 / (3.0 + 30.0),
		},
		{
			Rapper{
				weights: []Weight{
					{1.0, 10.0},
					{2.0, 20.0},
					{3.0, 30.0},
				},
				maxWeight: 1.0 + 2.0 + 3.0 + 10.0 + 20.0 + 30.0,
			},
			Sentence{
				&Morph{"", "", "", "", "", "", "", "", "", "アイ"},
				&Morph{"", "", "", "", "", "", "", "", "", "ウ"},
				&Morph{"", "", "", "", "", "", "", "", "", "エコ"},
			},
			Sentence{
				&Morph{"", "", "", "", "", "", "", "", "", "カキクケ"},
				&Morph{"", "", "", "", "", "", "", "", "", "コ"},
			},
			(10.0 + 20.0 + 30.0 + 3.0) / (1.0 + 2.0 + 3.0 + 10.0 + 20.0 + 30.0),
		},
	}

	for idx, test := range tests {
		distance := test.rapper.Distance(test.sen1, test.sen2)
		if test.distance != distance {
			t.Errorf("[%d] expected %f, but got %f", idx, test.distance, distance)
		}
	}
}
