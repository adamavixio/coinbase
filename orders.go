package coinbaseclient

type Order struct {
	ProfileID   string `json:"profile_id,omitempty" bson:"profile_id,omitempty"`
	Type        string `json:"type,omitempty" bson:"type,omitempty"`
	Side        string `json:"side,omitempty" bson:"side,omitempty"`
	Stp         string `json:"stp,omitempty" bson:"stp,omitempty"`
	Stop        string `json:"stop,omitempty" bson:"stop,omitempty"`
	TimeInForce string `json:"time_in_force,omitempty" bson:"time_in_force,omitempty"`
	CancelAfter string `json:"cancel_after,omitempty" bson:"cancel_after,omitempty"`
	PostOnly    bool   `json:"post_only,omitempty" bson:"post_only,omitempty"`
	StopPrice   string `json:"stop_price,omitempty" bson:"stop_price,omitempty"`
	Funds       string `json:"funds,omitempty" bson:"funds,omitempty"`
}
