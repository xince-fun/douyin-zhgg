include "user.thrift"

namespace go feed

struct Video {
    1: required i64 id,             // 视频唯一表示
    2: required user.User author,        // 视频作者信息
    3: required string play_url,    // 视频播放地址
    4: required string cover_url,   // 视频封面地址
    5: required i64 favorite_count, // 视频的点赞总数
    6: required i64 comment_count,  // 视频的评论总数
    7: required bool is_favorite,   // true-已点赞，false-未点赞
    8: required string title,       // 视频标题
}

// 基础返回信息 包括状态码和状态描述
struct BaseResp {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string statsu_msg, // 返回状态描述
}

// 视频流接口
struct douyin_feed_request {
    1: optional i64 latest_time,        // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
}

struct douyin_feed_response {
    1: required list<Video> vedio_list, // 视频列表
    2: optional i64 next_time,                // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
    3: required BaseResp base_resp,
}

service FeedService {
    douyin_feed_response GetUserFeed(1:douyin_feed_request req)
}