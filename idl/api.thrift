namespace go api

struct User {
    1: required i64 id,             // 用户id
    2: required string name,        // 用户名称
    3: optional i64 follow_count,   // 关注总数
    4: optional i64 follower_count, // 粉丝总数
    5: required bool is_follow,     // true-已关注，false-未关注
}

struct Video {
    1: required i64 id,             // 视频唯一表示
    2: required User author,        // 视频作者信息
    3: required string play_url,    // 视频播放地址
    4: required string cover_url,   // 视频封面地址
    5: required i64 favorite_count, // 视频的点赞总数
    6: required i64 comment_count,  // 视频的评论总数
    7: required bool is_favorite,   // true-已点赞，false-未点赞
    8: required string title,       // 视频标题
}

// 互动接口
struct Comment {
    1: required i64 id,             // 视频评论id
    2: required User user,          // 评论用户信息
    3: required string content,     // 评论内容
    4: required string create_date, // 评论发布日期，格式 mm-dd
}

// 基础返回信息 包括状态码和状态描述
struct BaseResp {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string statsu_msg, // 返回状态描述
}

// User

// 用户注册接口
struct douyin_user_register_request {
    1: required string username ( api.form="username", api.vd="len($) < 32" ), // 注册用户名，最长32个字符
    2: required string password ( api.form="password", api.vd="len($) < 32" ), // 密码，最长32个字符
}

struct douyin_user_register_response {
    1: required i64 user_id,        // 用户id
    2: required string token,       // 用户鉴权token
    3: required BaseResp base_resp,
}

// 用户登陆接口
struct douyin_user_login_request {
    1: required string username ( api.form="username", api.vd="len($) < 32" ), // 登录用户名
    2: required string password ( api.form="username", api.vd="len($) < 32" ), // 登录密码
}

struct douyin_user_login_response {
    1: required i64 user_id,        // 用户id
    2: required string token,       // 用户鉴权token
    3: required BaseResp base_resp,
}

// 用户信息
struct douyin_user_request {
    1: required i64 user_id ( api.query="user_id" ),   // 用户id
    2: required string token,  // 用户鉴权token
}

struct douyin_user_response {
    1: required User user,          // 用户信息
    2: required BaseResp base_resp,
}

// Feed

// 视频流接口
struct douyin_feed_request {
    1: optional i64 latest_time,        // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional string token,           // 可选参数，登录用户设置
}

struct douyin_feed_response {
    1: required list<Video> vedio_list, // 视频列表
    2: optional i64 next_time,                // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
    3: required BaseResp base_resp,
}

// Rublish

// 视频投稿
struct douyin_publish_action_request {
    1: required string token,                       // 用户鉴权token
    2: required binary data,                        // 视频数据
    3: required string title ( api.form="title", api.vd="len($) > 0" ), // 视频标题
}

struct douyin_publish_action_response {
    1: required BaseResp base_resp,
}

// 发布列表
struct douyin_publish_list_request {
    1: required i64 user_id ( vt.gt = "0" ), // 用户id
    2: required string token,                // 用户鉴权token
}

struct douyin_publish_list_response {
    1: required list<Video> video_list, // 用户发布的视频列表
    2: required BaseResp base_resp,
}

// Comment

struct douyin_comment_action_request {
    1:required string token // 用户鉴权token
    2:required i64 video_id // 视频id
    3:required i32 action_type // 1-发布评论，2-删除评论
    4:optional string comment_text // 用户填写的评论内容，在action_type=1的时候使用
    5:optional i64 comment_id // 要删除的评论id，在action_type=2的时候使用
}

struct douyin_comment_action_response {
    1:optional Comment comment // 评论成功返回评论内容，不需要重新拉取整个列表
    2:BaseResp base_resp
}

struct douyin_comment_list_request {
    1:required string token // 用户鉴权token
    2:required i64 video_id // 视频id
}

struct douyin_comment_list_response {
    1:BaseResp base_resp
}

// FAVORITE

struct douyin_favorite_action_request {
    1: required string token,    // 用户鉴权token
    2: required i64 video_id,    // 视频id
    3: required i32 action_type, // 1-点赞，2-取消点赞
}

struct douyin_favorite_action_response {
    1: required BaseResp base_resp,
}

struct douyin_favorite_list_request {
    1: required i64 user_id,  // 用户id
    2: required string token, // 用户鉴权token
}

struct douyin_favorite_list_response {
    1: required list<Video> video_list, // 用户点赞视频列表
    2: required BaseResp base_resp,
}

// RELATION

struct douyin_relation_action_request {
    1:required i64 user_id // 用户id
    2:required string token // 用户鉴权token
    3:required i64 to_user_id // 对方用户id
    4:required i32 action_type // 1-关注，2-取消关注
}

struct douyin_relation_action_response {
    1:BaseResp base_resp
}

struct douyin_relation_follow_list_request {
    1:required i64 user_id // 用户id
    2:required string token // 用户鉴权token
}

struct douyin_relation_follow_list_response {
    1:required list<User> user_list // 用户信息列表
    2:BaseResp base_resp
}

struct douyin_relation_follower_list_request {
        1:required i64 user_id // 用户id
        2:required string token // 用户鉴权token
}

struct douyin_relation_follower_list_response {
    1:required list<User> user_list // 用户列表
        2:BaseResp base_resp
}

service ApiService {
    // basic service
    douyin_feed_response GetUserFeed(1:douyin_feed_request req) (api.get="/douyin/feed/")
    douyin_user_register_response UserRegister(1: douyin_user_register_request req) (api.post="/douyin/user/register/")
    douyin_user_login_response UserLogin(1: douyin_user_login_request req) (api.post="/douyin/user/login/")
    douyin_user_response UserInfo(1: douyin_user_request req) (api.get="/douyin/user/")
    douyin_publish_action_response PublishAction(1: douyin_publish_action_request req) (api.post="/douyin/publish/action/")
    douyin_publish_list_response PublishList(1: douyin_publish_list_request req) (api.get="/douyin/publish/list/")

    // interaction service
    douyin_favorite_action_response FavoriteAction(1: douyin_favorite_action_request req) (api.post="/douyin/favorite/action/")
    douyin_favorite_list_response FavoriteList(1: douyin_favorite_list_request req) (api.get="/douyin/favorite/list/")
    douyin_comment_action_response CommentAction(1: douyin_comment_action_request req) (api.post="/douyin/comment/action/")
    douyin_comment_list_response CommentList(1: douyin_comment_list_request req) (api.get="/douyin/comment/list/")

    // social service
    douyin_relation_action_response RelationAction(1: douyin_relation_action_request req) (api.get="/douyin/relation/action/")
    douyin_relation_follow_list_response RelationFollowList(1: douyin_relation_follow_list_request req) (api.get="/douyin/relatioin/follow/list/")
    douyin_relation_follower_list_response RelationFollowerList(1: douyin_relation_follower_list_request req) (api.get="/douyin/relation/follower/list/")
}