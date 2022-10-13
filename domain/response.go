package domain

type Response struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

type RespRegister struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RespLogin struct {
	Token string `json:"token"`
}

type RespUpdateUser struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Age       int    `json:"age"`
	UpdatedAt string `json:"updated_at"`
}

type RespCreatePhoto struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type RespGetPhotoUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type RespGetPhoto struct {
	Id        int              `json:"id"`
	Title     string           `json:"title"`
	Caption   string           `json:"caption"`
	PhotoUrl  string           `json:"photo_url"`
	UserId    int              `json:"user_id"`
	CreatedAt string           `json:"created_at"`
	UpdatedAt string           `json:"updated_at"`
	User      RespGetPhotoUser `json:"User"`
}

type RespUpdatePhoto struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url"`
	UserId    int    `json:"user_id"`
	UpdatedAt string `json:"updated_at"`
}

type RespCreateComment struct {
	Id        int    `json:"id"`
	Message   string `json:"message"`
	PhotoId   int    `json:"photo_id"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type RespGetCommentUser struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type RespGetCommentPhoto struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}

type RespGetComment struct {
	Id        int                 `json:"id"`
	Message   string              `json:"message"`
	PhotoId   int                 `json:"photo_id"`
	UserId    int                 `json:"user_id"`
	CreatedAt string              `json:"created_at"`
	UpdatedAt string              `json:"updated_at"`
	User      RespGetCommentUser  `json:"User"`
	Photo     RespGetCommentPhoto `json:"Photo"`
}

type RespUpdateComment struct {
	Id        int    `json:"id"`
	Message   string `json:"message"`
	PhotoId   int    `json:"photo_id"`
	UserId    int    `json:"user_id"`
	UpdatedAt string `json:"updated_at"`
}

type RespCreateSocialMedia struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         int    `json:"user_id"`
	CreatedAt      string `json:"created_at"`
}

type RespGetSocialMediaUser struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type RespGetSocialMediaUserItem struct {
	Id             int                    `json:"id"`
	Name           string                 `json:"name"`
	SocialMediaUrl string                 `json:"social_media_url"`
	UserId         int                    `json:"user_id"`
	CreatedAt      string                 `json:"created_at"`
	UpdatedAt      string                 `json:"updated_at"`
	User           RespGetSocialMediaUser `json:"User"`
}

type RespGetSocialMedia struct {
	SosialMedias []RespGetSocialMediaUserItem `json:"sosial_medias"`
}

type RespUpdateSocialMedia struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         int    `json:"user_id"`
	UpdatedAt      string `json:"updated_at"`
}
