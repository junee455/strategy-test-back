package gameinstance

import (
	"fmt"
	"strategy-test-back/src/core/helpers"
	"time"
)

type IGameInstance interface {
	AddGameplayTask(task IGameplayTask)
	RemoveGameplayTask(task IGameplayTask)

	AddPersisentEffect(effect *GameStateEffect)
	RemovePersisentEffect(effect *GameStateEffect)

	AddImmediateEffect(effect *GameStateEffect)
	RemoveImmediateEffect(effect *GameStateEffect)

	GetGameState() *GameInstance

	Start()
	Pause()
	Resume()
}

type TickCallback = func(c TickContext, gi *GameInstance)

type GameInstance struct {
	Characters    []*Character
	EventNotifier IEventNotifier
	gameplayTasks []IGameplayTask

	gameStartTime time.Time
	prevTickTime  time.Time

	persisentEffects []*GameStateEffect

	nextImmediateEffects []*GameStateEffect

	currentTime float64
	currentDt   float64
	currentTick int

	tickCallbacks []*TickCallback
}

func (gi *GameInstance) AddGameplayTask(task ...IGameplayTask) {
	var tickContext = TickContext{
		Dt:          gi.currentDt,
		Time:        gi.currentTime,
		CurrentTick: gi.currentTick,
	}

	for _, t := range task {
		var description = t.GetDescription()

		if description.ID == "move" {
			// remove old move tasks from characters
			for _, currTask := range gi.gameplayTasks {
				var ttDescription = currTask.GetDescription()

				var isMoveTask = ttDescription.ID == "move"
				var isSameCharacter = description.Metadata == ttDescription.Metadata

				if isMoveTask && isSameCharacter {
					currTask.SetShouldRemove()
				}
			}
		}

		t.OnAdd(tickContext, gi)
	}

	gi.gameplayTasks = append(gi.gameplayTasks, task...)
}

func (gi *GameInstance) RemoveGameplayTask(task IGameplayTask) {
	gi.gameplayTasks = helpers.Filter(gi.gameplayTasks, func(tt IGameplayTask) bool {
		return tt != task
	})
}

func (gi *GameInstance) AddPersisentEffect(effect *GameStateEffect) {
	gi.persisentEffects = append(gi.persisentEffects, effect)

	var tickContext = TickContext{
		Dt:          gi.currentDt,
		Time:        gi.currentTime,
		CurrentTick: gi.currentTick,
	}

	switch effect.Type {
	case Bash:
		{
			ProcessBashEffect(effect, tickContext, gi)
		}
	}
}

func (gi *GameInstance) RemovePersisentEffect(effect *GameStateEffect) {
	gi.persisentEffects = helpers.Filter(gi.persisentEffects, func(e *GameStateEffect) bool {
		return e != effect
	})
}

func (gi *GameInstance) AddImmediateEffect(effect *GameStateEffect) {
	gi.nextImmediateEffects = append(gi.nextImmediateEffects, effect)
}

func (gi *GameInstance) RemoveImmediateEffect(effect *GameStateEffect) {
	gi.nextImmediateEffects = helpers.Filter(gi.nextImmediateEffects, func(e *GameStateEffect) bool {
		return e != effect
	})
}

func (gi *GameInstance) GetGameState() *GameInstance {
	return gi
}

func (gi *GameInstance) Start() {
	fmt.Printf("start new game")

	var ticker = time.NewTicker(time.Duration(TickScale * 1000_000_000))

	gi.gameStartTime = time.Now()
	gi.prevTickTime = time.Now()

	for {
		select {
		case tick := <-ticker.C:
			{
				var dt = tick.Sub(gi.prevTickTime)

				gi.currentDt = float64(dt.Microseconds()) / 1000_000.0

				var timeNow = time.Now()

				gi.prevTickTime = timeNow

				gi.currentTime = float64(timeNow.Sub(gi.gameStartTime).Microseconds() / 1000_000.0)

				var tickContext = TickContext{
					Dt:          gi.currentDt,
					Time:        gi.currentTime,
					CurrentTick: gi.currentTick,
				}

				for _, tc := range gi.tickCallbacks {
					(*tc)(tickContext, gi)
				}

				go gi.onTick()
			}
		default:
		}
	}
}
func (gi *GameInstance) Pause()  {}
func (gi *GameInstance) Resume() {}

func NewGameInstance() *GameInstance {
	return &GameInstance{
		Characters:    make([]*Character, 0),
		gameplayTasks: make([]IGameplayTask, 0),
		tickCallbacks: make([]*TickCallback, 0),
	}
}

func (gi *GameInstance) onTick() {
	gi.currentTick++

	var tickContext = TickContext{
		Dt:          gi.currentDt,
		Time:        gi.currentTime,
		CurrentTick: gi.currentTick,
	}

	gi.gameplayTasks = helpers.Filter(gi.gameplayTasks, func(t IGameplayTask) bool {
		return !t.GetShouldRemove()
	})

	for _, gt := range gi.gameplayTasks {
		gt.OnFirstTick(tickContext, gi)
		gt.OnTasksTick(tickContext, gi)
	}

	gi.nextImmediateEffects = helpers.Filter(gi.nextImmediateEffects, func(t *GameStateEffect) bool {
		return t != nil
	})

	// process immediate and persistend effects

	for _, effect := range gi.nextImmediateEffects {
		// effect.Type
		switch effect.Type {
		case ApplyDamageEffect:
			ProcessApplyDamageEffect(effect, tickContext, gi)
		case MoveEffect:
			ProcessMoveEffect(effect, tickContext, gi)
		case TeleportEffect:
			ProcessTeleportEffect(effect, tickContext, gi)
		case MoveProjectileEffect:
		case RestoreHealthEffect:
		case RestoreManaEffect:
		case Spawn:
			ProcessSpawnEffect(effect, tickContext, gi)
		case Dispel:

		default:
			//
			// }
		}
	}

	gi.nextImmediateEffects = make([]*GameStateEffect, 0)
}

func (gi *GameInstance) FindCharacterByID(id ActorID) *Character {
	for _, v := range gi.Characters {
		if v.RuntimeId == id {
			return v
		}
	}

	return nil
}

func (gi *GameInstance) AddTickCallback(c *TickCallback) {
	gi.tickCallbacks = append(gi.tickCallbacks, c)
}

func (gi *GameInstance) RemoveTickCallback(c *TickCallback) {
	gi.tickCallbacks = helpers.Filter(gi.tickCallbacks, func(t *TickCallback) bool {
		return t != c
	})

}

func (gi *GameInstance) GetTasksAmount() int {
	return len(gi.gameplayTasks)
}
