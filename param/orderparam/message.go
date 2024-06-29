package orderparam

import "github.com/iam-benyamin/hellofresh/entity/orderentity"

type Payload struct {
	Order orderentity.Order `json:"order"`
}

type Message struct {
	Producer string  `json:"producer"`
	SentAt   string  `json:"sent_at"`
	Type     string  `json:"type"`
	Payload  Payload `json:"payload"`
}
