include "user.thrift"

namespace go relation

// 基础返回信息 包括状态码和状态描述
struct BaseResp {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: optional string statsu_msg, // 返回状态描述
}

struct douyin_relation_action_request {
    1:required i64 user_id // 用户id
    3:required i64 to_user_id // 对方用户id
    4:required i32 action_type // 1-关注，2-取消关注
}

struct douyin_relation_action_response {
    1:BaseResp base_resp
}

struct douyin_relation_follow_list_request {
    1:required i64 user_id // 用户id
}

struct douyin_relation_follow_list_response {
    1:required list<user.User> user_list // 用户信息列表
    2:BaseResp base_resp
}

struct douyin_relation_follower_list_request {
    1:required i64 user_id // 用户id
}

struct douyin_relation_follower_list_response {
    1:required list<user.User> user_list // 用户列表
    2:BaseResp base_resp
}

service RelationService {
    douyin_relation_action_response RelationAction(1: douyin_relation_action_request req)
    douyin_relation_follow_list_response RelationFollowList(1: douyin_relation_follow_list_request req)
    douyin_relation_follower_list_response RelationFollowerList(1: douyin_relation_follower_list_request req)
}