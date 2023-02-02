include "feed.thrift"

namespace go publish

// 基础返回信息 包括状态码和状态描述
struct BaseResp {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string statsu_msg, // 返回状态描述
}

// 视频投稿
struct douyin_publish_action_request {
    1: required string token,                       // 用户鉴权token
    2: required binary data,                        // 视频数据
    3: required string title ( vt.min_size = "1" ), // 视频标题
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
    1: required list<feed.Video> video_list, // 用户发布的视频列表
    2: required BaseResp base_resp,
}

service PublishService {
    douyin_publish_action_response PublishAction(1: douyin_publish_action_request req)
    douyin_publish_list_response PublishList(1: douyin_publish_list_request req)
}