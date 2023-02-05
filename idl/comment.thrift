include "user.thrift"

namespace go comment

// 互动接口
struct Comment {
    1: required i64 id,             // 视频评论id
    2: required user.User user,          // 评论用户信息
    3: required string content,     // 评论内容
    4: required string create_date, // 评论发布日期，格式 mm-dd
}

// 基础返回信息 包括状态码和状态描述
struct BaseResp {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string statsu_msg, // 返回状态描述
}

struct douyin_comment_action_request {
    1:required i64 video_id // 视频id
    2:required i32 action_type // 1-发布评论，2-删除评论
    3:optional string comment_text // 用户填写的评论内容，在action_type=1的时候使用
    4:optional i64 comment_id // 要删除的评论id，在action_type=2的时候使用
}

struct douyin_comment_action_response {
    1:optional Comment comment // 评论成功返回评论内容，不需要重新拉取整个列表
    2:BaseResp base_resp
}

struct douyin_comment_list_request {
    1:required i64 video_id // 视频id
}

struct douyin_comment_list_response {
    1:BaseResp base_resp
}

service CommentService {
    douyin_comment_action_response CommentAction(1: douyin_comment_action_request req)
    douyin_comment_list_response CommentList(1: douyin_comment_list_request req)
}