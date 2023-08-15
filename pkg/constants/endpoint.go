package constants

const (
	ApiGroup        string = "/api"
	WsGroup         string = "/ws"
	DocGroup        string = "/docs/*"
	StaticGroupPath string = "/"
	StaticGroupName string = "public"
)

const (
	AuthGroupEndPoint string = "/auth"
	SignUpEndPoint    string = "/signup"
	SignInEndPoint    string = "/signin"
	SignOutEndPoint   string = "/signout"
)

const (
	RoomGroupEndPoint      string = "/rooms"
	CreateRoomEndPoint     string = ""
	GetAllRoomsEndPoint    string = ""
	GetRoomsOfUserEndPoint string = "/:userId"
	GetChatHistoryEndPoint string = "/chat-history"
	GetContactListEndPoint string = "/contact-list"
)

const (
	streamEndPoint       string = "/stream"
	CreateStreamEndPoint string = streamEndPoint + "/create"
	JoinStreamEndPoint   string = streamEndPoint + "/join"
)

const (
	ChatConnectEndPoint string = WsGroup + ""
)

const (
	ResourceGroupEndPoint               string = "/resources"
	localResourceEndPoint               string = "/local"
	UploadSingleLocalResourceEndPoint   string = localResourceEndPoint + "/single"
	UploadMultipleLocalResourceEndPoint string = localResourceEndPoint + "/multiple"
	DeleteSingleLocalResourceEndPoint   string = localResourceEndPoint + "/single/:fileName"
	DeleteMultipleLocalResourceEndPoint string = localResourceEndPoint + "/multiple"
)

const (
	SearchGroupEndPoint string = "/search"
	SearchRoomEndPoint  string = "/room"
)

const (
	UserGroupEndPoint string = "/users"
	GetUserByIdOrName string = "/:idOrName"
)
