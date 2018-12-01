package donut

import (
	"math"
	"math/rand"
)

type Donut struct {
	Radius     float32 //4
	Thick      float32 //4
	Toppings   []string //24
	GlutenFree bool // 4
	Hole       bool // 4
	Filling    string // 16
}

type DonutPreferences struct {
	Radius     float32
	Thick      float32
	Toppings   map[string]float32
	GlutenFree float32
	Hole       float32
	Filling    map[string]float32
}

const maxToppings = 3

var radiuses = []float32{5, 10, 15}
var thicks = []float32{2, 3, 4}
var toppings = []string{"Chocolate", "Nuts", "Sugar", "Caramel"}
var fillings = []string{"", "Mermelade", "Chocolate", "Cream"}

var rnd = rand.New(rand.NewSource(321))

func RndPtr() *Donut {
	d := &Donut{
		Radius: radiuses[rnd.Intn(len(radiuses))],
		Thick:  thicks[rnd.Intn(len(radiuses))],
	}
	if rnd.Int()%2 == 0 {
		d.GlutenFree = true
	} else {
		d.GlutenFree = false
	}
	if rnd.Int()%2 == 0 {
		d.Hole = true
	} else {
		d.Hole = false
		d.Filling = fillings[rnd.Intn(len(fillings))]
	}
	numToppings := rnd.Intn(maxToppings)
	d.Toppings = make([]string, 0, numToppings)
	for i := 0; i < numToppings; i++ {
		d.Toppings = append(d.Toppings,
			toppings[rnd.Intn(len(toppings))])
	}
	return d
}

func RndVal() Donut {
	d := Donut{
		Radius: radiuses[rnd.Intn(len(radiuses))],
		Thick:  thicks[rnd.Intn(len(radiuses))],
	}
	if rnd.Int()%2 == 0 {
		d.GlutenFree = true
	} else {
		d.GlutenFree = false
	}
	if rnd.Int()%2 == 0 {
		d.Hole = true
	} else {
		d.Hole = false
		d.Filling = fillings[rnd.Intn(len(fillings))]
	}
	numToppings := rnd.Intn(maxToppings)
	d.Toppings = make([]string, 0, numToppings)
	for i := 0; i < numToppings; i++ {
		d.Toppings = append(d.Toppings,
			toppings[rnd.Intn(len(toppings))])
	}
	return d
}

func RndPreferences() DonutPreferences {
	dp := DonutPreferences{
		Radius:     radiuses[rnd.Intn(len(radiuses))],
		Thick:      thicks[rnd.Intn(len(radiuses))],
		GlutenFree: float32(10 - rnd.Intn(20)),
		Hole:       float32(10 - rnd.Intn(20)),
		Toppings:   map[string]float32{},
		Filling:    map[string]float32{},
	}
	for _, topping := range toppings {
		dp.Toppings[topping] = rnd.Float32() * 10
	}
	for _, filling := range fillings {
		dp.Filling[filling] = rnd.Float32() * 10
	}
	return dp
}

//var out = bytes.NewBuffer(make([]byte, 0, 1000))
//var logger = log.New(out, "", 0)

func ScorePtr(d *Donut, p *DonutPreferences) float32 {
	//logger.Print(d)    //f***ng escape analysis
	score := p.Filling[d.Filling] + float32(math.Abs(float64(p.Radius-d.Radius))) +
		float32(math.Abs(float64(p.Thick-d.Thick)))
	if d.GlutenFree {
		score += p.GlutenFree
	}
	if d.Hole {
		score += p.Hole
	}
	for _, topping := range d.Toppings {
		score += p.Toppings[topping]
	}
	return score
}

func ScoreVal(d Donut, p DonutPreferences) float32 {
	//logger.Print(d)    //f***ng escape analysis
	score := p.Filling[d.Filling] + float32(math.Abs(float64(p.Radius-d.Radius))) +
		float32(math.Abs(float64(p.Thick-d.Thick)))
	if d.GlutenFree {
		score += p.GlutenFree
	}
	if d.Hole {
		score += p.Hole
	}
	for _, topping := range d.Toppings {
		score += p.Toppings[topping]
	}
	return score
}
