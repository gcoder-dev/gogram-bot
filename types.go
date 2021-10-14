package gogram

// Update from webhook
type Update struct {
	UpdateId      int           `json:"update_id"`
	Message       Message       `json:"message"`
	CallbackQuery CallbackQuery `json:"callback_query"`
	Poll          Poll          `json:"poll"`
}

type CallbackQuery struct {
	Id           string  `json:"id"`
	Message      Message `json:"message"`
	Chat         Chat    `json:"chat"`
	ChatInstance string  `json:"chat_instance"`
	Data         string  `json:"data"`
}

type MessageEntity struct {
	Type     string
	offset   int
	length   int
	url      string
	user     User
	language string
}

type Poll struct {
	Id              string       `json:"id,omitempty"`
	Question        string       `json:"question,omitempty"`
	Options         []PollOption `json:"options,omitempty"`
	TotalVoterCount int          `json:"total_voter_count,omitempty"`
	IsAnonymous     bool         `json:"is_anonymous,omitempty"`
	CloseDate       int          `json:"close_date,omitempty"`
	// Pass True, if the poll needs to be immediately closed.
	// This can be useful for poll preview.
	IsClosed                 bool            `json:"is_closed,omitempty"`
	Type                     string          `json:"type,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	CorrectOptionId          int             `json:"correct_option_id,omitempty"`
	Explanation              string          `json:"explanation,omitempty"`
	ExplanationEntities      []MessageEntity `json:"explanation_entities,omitempty"`
	// Amount of time in seconds the poll will be active after
	// creation, 5-600. Can't be used together with close_date.
	OpenPeriod int `json:"open_period,omitempty"`
}

type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size"`
}

type Message struct {
	MessageId       int                  `json:"message_id"`
	User            User                 `json:"from"`
	Chat            Chat                 `json:"chat"`
	Text            string               `json:"text"`
	Animation       Animation            `json:"animation"`
	Photo           []PhotoSize          `json:"photo"`
	Date            int                  `json:"date"`
	ReplyMarkup     inlineKeyboardMarkup `json:"reply_markup"`
	Poll            Poll                 `json:"poll"`
	NewChatPhoto    []PhotoSize          `json:"new_chat_photo"`
	NewChatTitle    string               `json:"new_chat_title"`
	NewChatMembers  []User               `json:"new_chat_members"`
	DeleteChatPhoto bool                 `json:"delete_chat_photo"`
}

// TypeIndicator function returns the type of message
// This make it easier to know which fields are empty and which aren't
// TypeIndicator may return "Text", "Animation", "Photo" and etc
func (m Message) TypeIndicator() string {
	switch {
	case m.Text != "":
		return "Text"
	case m.Animation != Animation{}:
		return "Animation"
	case m.Photo != nil:
		return "Photo"
	case m.DeleteChatPhoto == true:
		return "DeleteChatPhoto"
	case m.NewChatPhoto != nil:
		return "NewChatPhoto"
	default:
		return "Unknown"
	}
}

type ReplyAble struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// User id is a unique identification number of a particular Telegram user.
// However, the Telegram Chat id is a unique identification
// number of a Telegram chat (personal or group chat).
// Use Chat id for groups, and User id for a specific user
type User struct {
	// SupportsInlineQueries shows if Bot supports inline queries
	// This field is only for bots
	SupportsInlineQueries bool   `json:"supports_inline_queries"`
	LanguageCode          string `json:"language_code"`
	IsBot                 bool   `json:"is_bot"`
	ReplyAble
}

// Chat id is a unique identification number of a Telegram chat (personal or group chat).
// However, the Telegram User id is a unique identification number of a particular Telegram user.
// Use Chat id for groups, and User id for a specific user
type Chat struct {
	Type                  string          `json:"type"`
	Title                 string          `json:"title"`
	Bio                   string          `json:"bio"`
	Description           string          `json:"description"`
	InviteLink            string          `json:"invite_link"`
	PinnedMessage         *Message        `json:"pinned_message"`
	Permissions           ChatPermissions `json:"permissions"`
	SlowModeDelay         int             `json:"slow_mode_delay"`
	MessageAutoDeleteTime int             `json:"message_auto_delete_time"`
	StickerSetName        string          `json:"sticker_set_name"`
	CanSetStickerSet      bool            `json:"can_set_sticker_set"`
	LinkedChatId          int             `json:"linked_chat_id"`
	ReplyAble
}

type Animation struct {
	FileId string `json:"file_id"`
}

// inputMediaPhoto Represents a photo to be sent.
type inputMediaPhoto struct {
	// Type of the result, must be "photo"
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name>
	Media string `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
}

type inputMediaVideo struct {
	// Type of the result, must be "video"
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name>
	Media string `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption           string          `json:"caption"`
	ParseMode         string          `json:"parse_mode"`
	Width             int             `json:"width"`
	Height            int             `json:"height"`
	Duration          int             `json:"duration"`
	SupportsStreaming bool            `json:"supports_streaming"`
	CaptionEntities   []MessageEntity `json:"caption_entities"`
}

type inputMediaDocument struct {
	// Type of the result, must be "document"
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name>
	Media string `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Optional. Disables automatic server-side content type detection for files uploaded using
	// multipart/form-data. Always true, if the document is sent as part of an album.
	DisableContentTypeDetection bool `json:"disable_content_type_detection"`
}

type inputMediaAudio struct {
	// Type of the result, must be "audio"
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name>
	Media string `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	Duration        int             `json:"duration"`
	Performer       string          `json:"performer"`
	Tile            string          `json:"tile"`
}

type ChatInviteLink struct {
	InviteLink  string `json:"invite_link"`
	Creator     User   `json:"creator"`
	IsPrimary   bool   `json:"is_primary"`
	IsRevoked   bool   `json:"is_revoked"`
	ExpireDate  int    `json:"expire_date"`
	MemberLimit int    `json:"member_limit"`
}

type UserProfilePhotos struct {
	// Total number of profile pictures the target user has
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

type UserProfileResponse struct {
	Ok     bool              `json:"ok"`
	Result UserProfilePhotos `json:"result"`
}

type Response struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

type BooleanResponse struct {
	Result bool `json:"result"`
	Response
}

type StringResponse struct {
	Result string `json:"result"`
	Response
}

type MessageResponse struct {
	Result Message `json:"result"`
	Response
}

type InviteLinkResponse struct {
	Result ChatInviteLink `json:"result"`
	Response
}

type ChatResponse struct {
	Result Chat `json:"result"`
	Response
}

type ChatMember struct {
	Status                string `json:"status"`
	User                  User   `json:"user"`
	IsAnonymous           bool   `json:"is_anonymous"`
	CustomTitle           string `json:"custom_title"`
	IsMember              bool   `json:"is_member"`
	CanBeEdited           bool   `json:"can_be_edited"`
	CanManageChat         bool   `json:"can_manage_chat"`
	CanDeleteMessages     bool   `json:"can_delete_messages"`
	CanManageVoiceChats   bool   `json:"can_manage_voice_chats"`
	CanRestrictMembers    bool   `json:"can_restrict_members"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanPromoteMembers     bool   `json:"can_promote_members"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanPostMessages       bool   `json:"can_post_messages"`
	CanEditMessages       bool   `json:"can_edit_messages"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	// if member is restricted, UntilDate is the date when restrictions will be lifted for this user;
	// unix time. If 0, then the user is restricted forever. If -1 user is not
	// restricted.
	UntilDate int `json:"until_date"`
}

type ChatMemberResponse struct {
	Result []ChatMember `json:"result"`
	Response
}

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
}
