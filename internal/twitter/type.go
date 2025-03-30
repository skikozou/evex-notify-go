package twitter

type TwitterNotifications map[string]TwitterNotification

type TwitterNotification struct {
	ID          string `json:"id"`
	TimestampMs string `json:"timestampMs"`
	Icon        struct {
		ID string `json:"id"`
	} `json:"icon"`
	Message struct {
		Text     string `json:"text"`
		Entities []struct {
			FromIndex int `json:"fromIndex"`
			ToIndex   int `json:"toIndex"`
			Ref       struct {
				User struct {
					ID string `json:"id"`
				} `json:"user"`
			} `json:"ref"`
		} `json:"entities"`
		RTL bool `json:"rtl"`
	} `json:"message"`
	Template struct {
		AggregateUserActionsV1 struct {
			TargetObjects []struct {
				Tweet struct {
					ID string `json:"id"`
				} `json:"tweet"`
			} `json:"targetObjects"`
			FromUsers []struct {
				User struct {
					ID string `json:"id"`
				} `json:"user"`
			} `json:"fromUsers"`
		} `json:"aggregateUserActionsV1"`
	} `json:"template"`
}
