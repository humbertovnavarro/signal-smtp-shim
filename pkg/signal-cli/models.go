package signalcli

type Group struct {
	ID                    string   `json:"id"`
	Name                  string   `json:"name"`
	Description           string   `json:"description"`
	IsMember              bool     `json:"isMember"`
	IsBlocked             bool     `json:"isBlocked"`
	MessageExpirationTime int      `json:"messageExpirationTime"`
	Members               []Member `json:"members"`
	PendingMembers        []string `json:"pendingMembers"`
	RequestingMembers     []string `json:"requestingMembers"`
	Admins                []Admin  `json:"admins"`
	Banned                []string `json:"banned"`
	PermissionAddMember   string   `json:"permissionAddMember"`
	PermissionEditDetails string   `json:"permissionEditDetails"`
	PermissionSendMessage string   `json:"permissionSendMessage"`
	GroupInviteLink       string   `json:"groupInviteLink"`
}

type Member struct {
	Number string `json:"number"`
	UUID   string `json:"uuid"`
}

type Admin struct {
	Number string `json:"number"`
	UUID   string `json:"uuid"`
}

type GroupInfo struct {
	GroupID string `json:"groupId"`
	Type    string `json:"type"`
}

type Sticker struct {
	PackID    string `json:"packId"`
	StickerID int    `json:"stickerId"`
}

type Attachment struct {
	ContentType     string  `json:"contentType"`
	Filename        string  `json:"filename"`
	ID              string  `json:"id"`
	Size            int     `json:"size"`
	Width           *int    `json:"width,omitempty"`
	Height          *int    `json:"height,omitempty"`
	Caption         *string `json:"caption,omitempty"`
	UploadTimestamp *int    `json:"uploadTimestamp,omitempty"`
}

type Reaction struct {
	Emoji               string `json:"emoji"`
	TargetAuthor        string `json:"targetAuthor"`
	TargetAuthorNumber  string `json:"targetAuthorNumber"`
	TargetAuthorUUID    string `json:"targetAuthorUuid"`
	TargetSentTimestamp int64  `json:"targetSentTimestamp"`
	IsRemove            bool   `json:"isRemove"`
}

type DataMessage struct {
	Timestamp        int64        `json:"timestamp"`
	Message          *string      `json:"message,omitempty"`
	ExpiresInSeconds int          `json:"expiresInSeconds"`
	ViewOnce         bool         `json:"viewOnce"`
	Sticker          *Sticker     `json:"sticker,omitempty"`
	Attachments      []Attachment `json:"attachments,omitempty"`
	GroupInfo        GroupInfo    `json:"groupInfo"`
	Reaction         *Reaction    `json:"reaction,omitempty"`
}

type EditMessage struct {
	TargetSentTimestamp int64       `json:"targetSentTimestamp"`
	DataMessage         DataMessage `json:"dataMessage"`
}

type Envelope struct {
	Source       string      `json:"source"`
	SourceNumber string      `json:"sourceNumber"`
	SourceUUID   string      `json:"sourceUuid"`
	SourceName   string      `json:"sourceName"`
	SourceDevice int         `json:"sourceDevice"`
	Timestamp    int64       `json:"timestamp"`
	DataMessage  DataMessage `json:"dataMessage"`
	Account      string      `json:"account"`
}

type Payload struct {
	Envelope Envelope `json:"envelope"`
	Account  string   `json:"account"`
}
