package problem

import (
	"container/list"
	"fmt"
	"math"

	"github.com/mrsmileg/transp/logger"
)

//Problem structure, describes transportation problem states
type Problem struct {
	Providers  []int        `json:"providers"`
	Consumers  []int        `json:"consumers"`
	Prices     [][]float64  `json:"prices"`
	Deliveries [][]Delivery `json:"deliveries"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Delivery struct {
	Count float64 `json:"count"`
	Price float64 `json:"price"`
	Point `json:"point"`
}

//Logger instance
var Logger = logger.NewLogger()

//EpsPotential represents epsila constant to check degenerative function
const EpsPotential = 0.001

var defaultDelivery = Delivery{}

//GetSolution gets solution of transportation problem by using of petentials method
func (problem *Problem) GetSolution() [][]Delivery {
	problem.normalize()
	problem.basePlan()
	problem.calculate()

	return problem.Deliveries
}

func (problem Problem) String() string {
	var out string
	totalCost := 0.0
	for x := 0; x < len(problem.Providers); x++ {
		for y := 0; y < len(problem.Consumers); y++ {
			delivery := problem.Deliveries[x][y]
			if delivery != defaultDelivery && delivery.X == x && delivery.Y == y {
				out = fmt.Sprintf("%s\t%d", out, int(delivery.Count))
				totalCost += delivery.Count * delivery.Price
			} else {
				out = fmt.Sprintf("%s\t0 ", out)
			}
		}
		out = fmt.Sprintf("%s\n", out)
	}
	out = fmt.Sprintf("%s\n\tTotal Cost: %.2f\n\n", out, totalCost)

	return out
}

func (problem *Problem) calculate() {
	var step []Delivery
	maxCut := 0.0
	last := defaultDelivery
	problem.repair(EpsPotential)
	for x := range problem.Providers {
		for y := range problem.Consumers {
			if problem.Deliveries[x][y] != defaultDelivery {
				continue
			}
			temp := Delivery{0, problem.Prices[x][y], Point{x, y}}
			position := problem.getClosedPositions(temp)
			cut := 0.0
			minCount := float64(math.MaxInt32)
			lastDelivery := defaultDelivery
			plus := true
			for _, delivery := range position {
				if plus {
					cut += delivery.Price
				} else {
					cut -= delivery.Price
					if delivery.Count < minCount {
						lastDelivery = delivery
						minCount = delivery.Count
					}
				}
				plus = !plus
			}
			if cut < maxCut {
				step = position
				last = lastDelivery
				maxCut = cut
			}
		}
	}

	if step != nil {
		q := last.Count
		plus := true
		for _, s := range step {
			if plus {
				s.Count += q
			} else {
				s.Count -= q
			}
			if s.Count == 0 {
				problem.Deliveries[s.X][s.Y] = defaultDelivery
			} else {
				problem.Deliveries[s.X][s.Y] = s
			}
			plus = !plus
		}
		problem.calculate()
	}
}

func (problem *Problem) normalize() {
	provCount, customCount := len(problem.Providers), len(problem.Consumers)
	diff := sum(problem.Providers) - sum(problem.Consumers)

	if diff > 0 {
		problem.Consumers = append(problem.Consumers, diff)
	} else if diff < 0 {
		problem.Providers = append(problem.Providers, -diff)
	}

	prices := make([][]float64, len(problem.Providers))
	for i := 0; i < len(problem.Providers); i++ {
		prices[i] = make([]float64, len(problem.Consumers))
	}

	for i := 0; i < provCount; i++ {
		for j := 0; j < customCount; j++ {
			prices[i][j] = problem.Prices[i][j]
		}
	}

	deliveries := make([][]Delivery, len(problem.Providers))
	for i := 0; i < len(problem.Providers); i++ {
		deliveries[i] = make([]Delivery, len(problem.Consumers))
	}

	problem.Prices = prices
	problem.Deliveries = deliveries
}

func sum(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}

	return sum
}

func (problem *Problem) basePlan() {
	for x, start := 0, 0; x < len(problem.Providers); x++ {
		for y := start; y < len(problem.Consumers); y++ {
			count := math.Min(float64(problem.Providers[x]), float64(problem.Consumers[y]))
			if count <= 0 {
				continue
			}

			problem.Deliveries[x][y] = Delivery{count, problem.Prices[x][y], Point{x, y}}
			problem.Providers[x] -= int(count)
			problem.Consumers[y] -= int(count)
			if problem.Providers[x] == 0 {
				start = y
				break
			}
		}
	}
}

func (problem *Problem) repair(eps float64) {
	if len(problem.Providers)+len(problem.Consumers)-1 == problem.toList().Len() {
		return
	}

	for x := 0; x < len(problem.Providers); x++ {
		for y := 0; y < len(problem.Consumers); y++ {
			if problem.Deliveries[x][y] == defaultDelivery {
				temp := Delivery{eps, problem.Prices[x][y], Point{x, y}}
				if len(problem.getClosedPositions(temp)) == 0 {
					problem.Deliveries[x][y] = temp
					break
				}
			}
		}
	}
}

func (problem *Problem) toList() *list.List {
	list := list.New()
	for _, deliveries := range problem.Deliveries {
		for _, delivery := range deliveries {
			if delivery != defaultDelivery {
				list.PushBack(delivery)
			}
		}
	}
	return list
}

func (problem *Problem) getClosedPositions(d Delivery) []Delivery {
	positions := problem.toList()
	positions.PushFront(d)

	var next *list.Element
	for ok, toremove := true, 0; ok; ok = toremove == 0 {
		for element := positions.Front(); element != nil; element = next {
			next = element.Next()
			neighbors := problem.getNeighbors(element.Value.(Delivery), positions)
			if neighbors[0] == defaultDelivery || neighbors[1] == defaultDelivery {
				positions.Remove(element)
				toremove++
			}
		}
	}

	deliveries := make([]Delivery, positions.Len())
	previous := d
	for i := 0; i < len(deliveries); i++ {
		deliveries[i] = previous
		previous = problem.getNeighbors(previous, positions)[i%2]
	}
	return deliveries
}

func (problem *Problem) getNeighbors(d Delivery, lst *list.List) [2]Delivery {
	var neighbors [2]Delivery
	for element := lst.Front(); element != nil; element = element.Next() {
		value := element.Value.(Delivery)
		if value != d {
			if value.X == d.X && neighbors[0] == defaultDelivery {
				neighbors[0] = value
			} else if value.Y == d.Y && neighbors[1] == defaultDelivery {
				neighbors[1] = value
			}
			if neighbors[0] != defaultDelivery && neighbors[1] != defaultDelivery {
				break
			}
		}
	}
	return neighbors
}
