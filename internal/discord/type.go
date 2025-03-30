package discord

type DiscordNotifications []struct {
	ID              string `json:"id"`
	Type            int    `json:"type"`
	Content         string `json:"content"`
	ChannelID       string `json:"channel_id"`
	Author          User   `json:"author"`
	Attachments     []any  `json:"attachments"`
	Embeds          []any  `json:"embeds"`
	Mentions        []User `json:"mentions"`
	MentionRoles    []any  `json:"mention_roles"`
	MentionEveryone bool   `json:"mention_everyone"`
	Pinned          bool   `json:"pinned"`
	TTS             bool   `json:"tts"`
	Timestamp       string `json:"timestamp"`
	EditedTimestamp any    `json:"edited_timestamp"`
	Flags           int    `json:"flags"`
	Components      []any  `json:"components"`
	MessageRef      struct {
		Type      int    `json:"type"`
		ChannelID string `json:"channel_id"`
		GuildID   string `json:"guild_id"`
		MessageID string `json:"message_id"`
	} `json:"message_reference"`
	Position int `json:"position"`
}

type User struct {
	ID               string `json:"id"`
	Username         string `json:"username"`
	GlobalName       string `json:"global_name"`
	Avatar           string `json:"avatar"`
	AvatarDecoration any    `json:"avatar_decoration_data"`
	Discriminator    string `json:"discriminator"`
	PublicFlags      int    `json:"public_flags"`
	Clan             any    `json:"clan"`
}
