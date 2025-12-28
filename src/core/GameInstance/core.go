package gameinstance

// number of seconds per tick
const TickScale float64 = 0.1

type TickDuration int64

type TickContext struct {
	Time        float64
	Dt          float64
	CurrentTick int
}

type ITickable interface {
	Tick(c TickContext)
}
