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
