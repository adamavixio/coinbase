package coinbaseclient

//
// Definitions
//

type Order struct {
	CreatedAt     string `json:"created_at,omitempty" bson:"created_at,omitempty"`
	ExecutedValue string `json:"executed_value,omitempty" bson:"executed_value,omitempty"`
	FillFees      string `json:"fill_fees,omitempty" bson:"fill_fees,omitempty"`
	FilledSize    string `json:"filled_size,omitempty" bson:"filled_size,omitempty"`
	ID            string `json:"id,omitempty" bson:"id,omitempty"`
	PostOnly      bool   `json:"post_only,omitempty" bson:"post_only,omitempty"`
	Price         string `json:"price,omitempty" bson:"price,omitempty"`
	ProductID     string `json:"product_id,omitempty" bson:"product_id,omitempty"`
	ProfileID     string `json:"profile_id,omitempty" bson:"profile_id,omitempty"`
	Settled       bool   `json:"settled,omitempty" bson:"settled,omitempty"`
	Side          string `json:"side,omitempty" bson:"side,omitempty"`
	Size          string `json:"size,omitempty" bson:"size,omitempty"`
	Status        string `json:"status,omitempty" bson:"status,omitempty"`
	TimeInForce   string `json:"time_in_force,omitempty" bson:"time_in_force,omitempty"`
	Type          string `json:"type,omitempty" bson:"type,omitempty"`
}
