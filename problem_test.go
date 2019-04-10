package problem

import (
	"reflect"
	"testing"
)

var problemOne = Problem{
	Providers: []int{12, 40, 33},
	Consumers: []int{20, 30, 10},
	Prices: [][]float64{
		{3, 5, 7},
		{2, 4, 6},
		{9, 1, 8},
	},
}

var problemTwo = Problem{
	Providers: []int{25, 35},
	Consumers: []int{20, 30, 10},
	Prices: [][]float64{
		{3, 5, 7},
		{3, 2, 5},
	},
}

var problemThree = Problem{
	Providers: []int{14, 10, 15, 12},
	Consumers: []int{10, 15, 12, 15},
	Prices: [][]float64{
		{10, 30, 25, 15},
		{20, 15, 20, 10},
		{10, 30, 20, 20},
		{30, 40, 35, 45},
	},
}

func TestNormalization(t *testing.T) {
	t.Run("The first task normalization data", testNormalizeFunc(&problemOne, Problem{
		Providers: []int{12, 40, 33},
		Consumers: []int{20, 30, 10, 25},
		Prices: [][]float64{
			{3, 5, 7, 0},
			{2, 4, 6, 0},
			{9, 1, 8, 0},
		},
	}))

	t.Run("The first task normalization data", testNormalizeFunc(&problemTwo, Problem{
		Providers: []int{25, 35},
		Consumers: []int{20, 30, 10},
		Prices: [][]float64{
			{3, 5, 7},
			{3, 2, 5},
		},
	}))

	t.Run("The second task normalization data", testNormalizeFunc(&problemThree, Problem{
		Providers: []int{14, 10, 15, 12, 1},
		Consumers: []int{10, 15, 12, 15},
		Prices: [][]float64{
			{10, 30, 25, 15},
			{20, 15, 20, 10},
			{10, 30, 20, 20},
			{30, 40, 35, 45},
			{0, 0, 0, 0},
		},
	}))
}

func TestBasePlan(t *testing.T) {
	t.Run("The first task base plan", testBasePlanFunc(&problemOne, [][]Delivery{
		{
			Delivery{
				Count: 12,
				Price: 3,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
		},
		{
			Delivery{
				Count: 8,
				Price: 2,
				Point: Point{1, 0},
			},
			Delivery{
				Count: 30,
				Price: 4,
				Point: Point{1, 1},
			},
			Delivery{
				Count: 2,
				Price: 6,
				Point: Point{1, 2},
			},
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
		},
		{
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 8,
				Price: 8,
				Point: Point{2, 2},
			},
			Delivery{
				Count: 25,
				Price: 0,
				Point: Point{2, 3},
			},
		},
	}))

	t.Run("The second task base plan", testBasePlanFunc(&problemTwo, [][]Delivery{
		{
			Delivery{
				Count: 20,
				Price: 3,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 5,
				Price: 5,
				Point: Point{0, 1},
			},
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
		},
		{
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 25,
				Price: 2,
				Point: Point{1, 1},
			},
			Delivery{
				Count: 10,
				Price: 5,
				Point: Point{1, 2},
			},
		},
	}))
}

func TestCalculation(t *testing.T) {
	t.Run("The first task solution calc", testCalcFunc(&problemOne, [][]Delivery{
		{
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 12,
				Price: 0,
				Point: Point{0, 3},
			},
		},
		{
			Delivery{
				Count: 20,
				Price: 2,
				Point: Point{1, 0},
			},
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 10,
				Price: 6,
				Point: Point{1, 2},
			},
			Delivery{
				Count: 10,
				Price: 0,
				Point: Point{1, 3},
			},
		},
		{
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 30,
				Price: 1,
				Point: Point{2, 1},
			},
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 3,
				Price: 0,
				Point: Point{2, 3},
			},
		},
	}))

	t.Run("The second task task solution calc", testCalcFunc(&problemTwo, [][]Delivery{
		{
			Delivery{
				Count: 20,
				Price: 3,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 5,
				Price: 7,
				Point: Point{0, 2},
			},
		},
		{
			Delivery{
				Count: 0,
				Price: 0,
				Point: Point{0, 0},
			},
			Delivery{
				Count: 30,
				Price: 2,
				Point: Point{1, 1},
			},
			Delivery{
				Count: 5,
				Price: 5,
				Point: Point{1, 2},
			},
		},
	}))
}

func testNormalizeFunc(problem *Problem, expected Problem) func(*testing.T) {
	return func(t *testing.T) {
		problem.normalize()
		if !reflect.DeepEqual(problem.Providers, expected.Providers) {
			t.Errorf("Expected providers: %v, but got providers: %v", expected.Providers, problem.Providers)
		}
		if !reflect.DeepEqual(problem.Consumers, expected.Consumers) {
			t.Errorf("Expected consumers: %v, but got consumers: %v", expected.Consumers, problem.Consumers)
		}
		if !reflect.DeepEqual(problem.Prices, expected.Prices) {
			t.Errorf("Expected prices matrix: %v, but got prices matrix: %v", expected.Prices, problem.Prices)
		}
	}
}

func testBasePlanFunc(problem *Problem, expected [][]Delivery) func(*testing.T) {
	return func(t *testing.T) {
		problem.basePlan()
		if !reflect.DeepEqual(problem.Deliveries, expected) {
			t.Errorf("Expected base plan: %v, but got base plan: %v", expected, problem.Deliveries)
		}
	}
}

func testCalcFunc(problem *Problem, expected [][]Delivery) func(*testing.T) {
	return func(t *testing.T) {
		problem.calculate()
		if !reflect.DeepEqual(problem.Deliveries, expected) {
			t.Errorf("Expected problem solution: %v, but got problem solution: %v", expected, problem.Deliveries)
		}
	}
}
