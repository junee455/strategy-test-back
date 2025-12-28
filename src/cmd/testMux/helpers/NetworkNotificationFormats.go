package helpers

type WsMovePayload struct {
	ActorID string     `json:"actorId"`
	From    [2]float64 `json:"from"`
	To      [2]float64 `json:"to"`
	Speed   float64    `json:"speed"`
}

type WsApplyDamagePayload struct {
	InstigatorID  string `json:"instigatorId"`
	TargetID      string `json:"targetId"`
	DamageApplied int    `json:"damageApplied"`
	HpLeft        int    `json:"hpLeft"`
}

type WsCastProjectilePayload struct {
	Type string `json:"type"`

	InstigatorID string     `json:"instigatorId,omitempty"`
	From         [2]float64 `json:"from,omitempty"`

	TargetID string     `json:"targetId,omitempty"`
	To       [2]float64 `json:"to,omitempty"`

	Speed float64 `json:"speed"`
}

type WsApplyEffectPayload struct {
	EffectType int    `json:"effect"`
	TargetID   string `json:"targetId"`
	Remove     bool   `json:"remove"`
}

type WsMoveStopPayload struct {
	ActorID string `json:"actorId"`
}

type WsNotification struct {
	Type    string `json:"type"`
	Tick    int    `json:"tick"`
	Payload any    `json:"payload"`
}
