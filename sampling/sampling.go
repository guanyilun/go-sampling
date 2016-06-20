package sampling

import (
    "math/rand"
    "time"
)

type Sampling struct {
    values map[int]float64
}

func (s *Sampling) AddValue(val int, prob float64) {
    s.values[val] = prob
}

func (s *Sampling) SetProb(val int, prob float64) {
    s.values[val] = prob
}

func (s *Sampling) Normalize() {
    var denorm float64 = 0
    
    for _, v := range s.values {
	denorm += v
    }
    
    for k, v := range s.values {
	s.SetProb(k, v/denorm)
    }
}

func (s *Sampling) Sample() int {
    var lower float64 = 0
    var upper float64 = 0
    
    rand.Seed(time.Now().UnixNano())
    r := rand.Float64()
    
    for k, v := range s.values {
	upper += v
	if r >= lower && r < upper {
	    return k
	}
	lower = upper
    }
    return -1 //ERR
}

func (s *Sampling) GetProb(val int) float64 {
    return s.values[val]
}

func (s *Sampling) AddBundleProbs(probs []float64) {
    l := len(probs)
    for i:=0; i<l; i++ {
	s.AddValue(i, probs[i])
    }
}

func NewSampling() *Sampling {
    var s Sampling
    s.values = make(map[int]float64)
    return &s
}