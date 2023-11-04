package update

type Unlike Like

type Like struct {
	LikerID       int    `json:"liker_id,omitempty"`
	ObjectType    string `json:"object_type,omitempty"`
	ObjectOwnerID int    `json:"object_owner_id,omitempty"`
	ThreadReplyID int    `json:"thread_reply_id,omitempty"`
	PostID        int    `json:"post_id,omitempty"`
	ObjectID      int    `json:"object_id,omitempty"`
}

type Likes struct {
	Count      int `json:"count,omitempty"`
	Likes      int `json:"user_likes,omitempty"`
	CanLike    int `json:"can_like,omitempty"`
	CanPublish int `json:"can_publish,omitempty"`
}
