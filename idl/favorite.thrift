include "feed.thrift"

namespace go favorite

// 基础返回信息 包括状态码和状态描述
struct BaseResp {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string statsu_msg, // 返回状态描述
}

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
    1: required list<feed.Video> video_list, // 用户点赞视频列表
    2: required BaseResp base_resp,
}

service FavoriteService {
    douyin_favorite_action_response FavoriteAction(1: douyin_favorite_action_request req),
    douyin_favorite_list_response FavoriteList(1: douyin_favorite_list_request req),
}