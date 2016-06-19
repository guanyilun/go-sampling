package sampling

import (
    "math/rand"
    "time"
)

type Sampling struct {
    values map[int]float32
}

func (s *Sampling) AddValue(val int, prob float32) {
    s.values[val] = prob
}

func (s *Sampling) SetProb(val int, prob float32) {
    s.values[val] = prob
}

func (s *Sampling) Normalize() {
    var denorm float32 = 0
    
    for _, v := range s.values {
	denorm += v
    }
    
    for k, v := range s.values {
	s.SetProb(k, v/denorm)
    }
}

func (s *Sampling) Sample() int {
    var lower float32 = 0
    var upper float32 = 0
    
    rand.Seed(time.Now().UnixNano())
    r := rand.Float32()
    
    for k, v := range s.values {
	upper += v
	if r >= lower && r < upper {
	    return k
	}
	lower = upper
    }
    return -1 //ERR
}

func (s *Sampling) GetProb(val int) float32 {
    return s.values[val]
}

func NewSampling() *Sampling {
    var s Sampling
    s.values = make(map[int]float32)
    return &s
}

/*
func main () {
    s := NewSampling() 
    s.AddValue(1, 1)
    s.AddValue(0, 1)
    s.Normalize()
    
    total := 0
    for i:=0;i<1000;i++ {
	val := s.Sample()
	total += val
    }
    fmt.Println(total)
}

*/