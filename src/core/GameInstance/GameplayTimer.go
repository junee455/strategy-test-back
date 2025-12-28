package gameinstance

func NewGameplayTimer(duration float64, gi *GameInstance) chan struct{} {
	var startTime = gi.currentTime

	var doneChan = make(chan struct{}, 1)

	var resultChan = make(chan struct{}, 1)

	var onTick TickCallback = func(c TickContext, tgi *GameInstance) {

		if c.Time-startTime >= duration {
			resultChan <- struct{}{}
			doneChan <- struct{}{}
		}
	}

	go func() {
		<-doneChan

		gi.RemoveTickCallback(&onTick)
	}()

	gi.AddTickCallback(&onTick)

	return resultChan
}
