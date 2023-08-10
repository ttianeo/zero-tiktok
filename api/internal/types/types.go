// Code generated by goctl. DO NOT EDIT.
package types

type UserRegisterReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type UserLoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type UserTokenResp struct {
	Code   int64  `json:"status_code"`
	Msg    string `json:"status_msg"`
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type FeedReq struct {
	Latest string `form:"latest_time,optional"`
	Token  string `form:"token,optional"`
}

type Author struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	FollowCount    int64  `json:"follow_count"`
	FollowerCount  int64  `json:"follower_count"`
	IsFollow       bool   `json:"is_follow"`
	Avatar         string `json:"avatar"`
	Background     string `json:"background_image"`
	Signature      string `json:"signature"`
	TotalFavorited string `json:"total_favorited"`
	WorkCount      int64  `json:"work_count"`
	FavoriteCount  int64  `json:"favorite_count"`
}

type Video struct {
	ID             int64  `json:"id"`
	Author         Author `json:"author"`
	PlayUrl        string `json:"play_url"`
	CoverUrl       string `json:"cover_url"`
	FavouriteCount int64  `json:"favourite_count"`
	CommentCount   int64  `json:"comment_count"`
	IsFavourite    bool   `json:"is_favourite"`
	Title          string `json:"title"`
}

type FeedResp struct {
	Code int64   `json:"status_code"`
	Msg  string  `json:"status_msg"`
	Next int64   `json:"next_time"`
	List []Video `json:"video_list"`
}

type UserInfoReq struct {
	UserID int64  `form:"user_id"`
	Token  string `form:"token"`
}

type UserInfoResp struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
	User Author `json:"user"`
}

type PublishListResp struct {
	Code int64   `json:"status_code"`
	Msg  string  `json:"status_msg"`
	List []Video `json:"video_list"`
}

type FavoriteActionReq struct {
	Token      string `form:"token"`
	VideoID    int64  `form:"video_id"`
	ActionType int64  `form:"action_type"`
}

type FavoriteActionResp struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

type Comment struct {
	ID         int64  `json:"id"`
	User       Author `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type CommentActionReq struct {
	Token       string `form:"token"`
	VideoID     string `form:"token"`
	ActionType  string `form:"action_type"`
	CommentText string `form:"comment_text"`
	CommentID   string `form:"comment_id"`
}

type CommentActionResp struct {
	Code    int64   `json:"status_code"`
	Msg     string  `json:"status_msg"`
	Comment Comment `json:"comment"`
}

type CommentListReq struct {
	Token   string `form:"token"`
	VideoID string `form:"video_id"`
}

type CommentListResp struct {
	Code int64     `json:"status_code"`
	Msg  string    `json:"status_msg"`
	List []Comment `json:"comment_list"`
}
