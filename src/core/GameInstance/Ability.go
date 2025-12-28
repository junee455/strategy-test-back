package gameinstance

type TaskContext[Payload any] struct {
	Instigator ActorID
	Payload    Payload
}

type IAbility[Payload any] interface {
	CreateTask(taskContext TaskContext[Payload]) GameplayTask
}

type Ability struct {
}
