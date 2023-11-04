package responses

import "github.com/botscommunity/vkgo/update"

type GroupsReply struct {
	Error    *Error `json:"error"`
	Response Groups `json:"response,omitempty"`
}

type Groups struct {
	Error  *Error  `json:"error"`
	Groups []Group `json:"groups,omitempty"`
}

type Group struct {
	Error            *Error    `json:"error"`
	ID               int       `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Domain           string    `json:"screen_name,omitempty"`
	Closed           int       `json:"is_closed,omitempty"`
	Deactivated      string    `json:"deactivated,omitempty"`
	IsAdmin          int       `json:"is_admin,omitempty"`
	AdminLevel       int       `json:"admin_level,omitempty"`
	IsMember         int       `json:"is_member,omitempty"`
	IsAdvertiser     int       `json:"is_advertiser,omitempty"`
	InvitedBy        int       `json:"invited_by,omitempty"`
	Type             string    `json:"type,omitempty"`
	Photo50          string    `json:"photo_50,omitempty"`
	Photo100         string    `json:"photo_100,omitempty"`
	Photo200         string    `json:"photo_200,omitempty"`
	Activity         string    `json:"activity,omitempty"`
	Addresses        Addresses `json:"addresses,omitempty"`
	Like             Like      `json:"like,omitempty"`
	AgeLimits        int       `json:"age_limits,omitempty"`
	BanInfo          BanInfo   `json:"ban_info,omitempty"`
	CanCreateTopic   int       `json:"can_create_topic,omitempty"`
	CanMessage       int       `json:"can_message,omitempty"`
	CanPost          int       `json:"can_post,omitempty"`
	CanSuggest       int       `json:"can_suggest,omitempty"`
	CanSeePosts      int       `json:"can_see_all_posts,omitempty"`
	CanUploadDoc     int       `json:"can_upload_doc,omitempty"`
	CanUploadStory   int       `json:"can_upload_story,omitempty"`
	CanUploadVideo   int       `json:"can_upload_video,omitempty"`
	City             City      `json:"city,omitempty"`
	Contacts         []Contact `json:"contacts,omitempty"`
	Counters         Counters  `json:"counters,omitempty"`
	Country          Country   `json:"country,omitempty"`
	Cover            Cover     `json:"cover,omitempty"`
	CropPhoto        CropPhoto `json:"crop_photo,omitempty"`
	Description      string    `json:"description,omitempty"`
	PinnedPost       int       `json:"fixed_post,omitempty"`
	HasPhoto         int       `json:"has_photo,omitempty"`
	IsFavorite       int       `json:"is_favorite,omitempty"`
	IsHiddenFromFeed int       `json:"is_hidden_from_feed,omitempty"`
	MessagesBlocked  int       `json:"is_messages_blocked,omitempty"`
	Links            []Link    `json:"links,omitempty"`
	AlbumID          int       `json:"main_album_id,omitempty"`
	Section          int       `json:"main_section,omitempty"`
	Market           Market    `json:"market,omitempty"`
	MemberStatus     int       `json:"member_status,omitempty"`
	MembersCount     int       `json:"members_count,omitempty"`
	Place            Place     `json:"place,omitempty"`
	PublicDate       string    `json:"public_date_label,omitempty"`
	Site             string    `json:"site,omitempty"`
	StartDate        string    `json:"start_date,omitempty"`
	FinishDate       string    `json:"finish_date,omitempty"`
	Status           string    `json:"status,omitempty"`
	Trending         int       `json:"trending,omitempty"`
	Verified         int       `json:"verified,omitempty"`
	Wall             int       `json:"wall,omitempty"`
	WikiPage         string    `json:"wiki_page,omitempty"`
}

type Addresses struct {
	Enabled   bool `json:"is_enabled,omitempty"`
	AddressID int  `json:"main_address_id,omitempty"`
}

type Like struct {
	Liked   bool        `json:"is_liked,omitempty"`
	Friends FriendsLike `json:"friends,omitempty"`
}

type FriendsLike struct {
	Count   int `json:"count,omitempty"`
	Preview any `json:"preview,omitempty"`
}

type BanInfo struct {
	EndDate int    `json:"end_date,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type City struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type Counter struct {
	Albums          int `json:"albums,omitempty"`
	Market          int `json:"market,omitempty"`
	Photos          int `json:"photos,omitempty"`
	Topics          int `json:"topics,omitempty"`
	Articles        int `json:"articles,omitempty"`
	Narratives      int `json:"narratives,omitempty"`
	ClassifiedYoula int `json:"classified_youla,omitempty"`
}

type Country struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type Cover struct {
	Enabled int     `json:"enabled,omitempty"`
	Images  []Image `json:"images,omitempty"`
}

type Image struct {
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type CropPhoto struct {
	Photo update.Photo `json:"photo,omitempty"`
	Crop  Crop         `json:"crop,omitempty"`
	Rect  Crop         `json:"rect,omitempty"`
}

type Crop struct {
	X  any `json:"x,omitempty"`
	Y  any `json:"y,omitempty"`
	X2 any `json:"x2,omitempty"`
	Y2 any `json:"y2,omitempty"`
}

type Link struct {
	ID          int    `json:"id,omitempty"`
	URL         string `json:"url,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"desc,omitempty"`
	Photo50     string `json:"photo_50,omitempty"`
	Photo100    string `json:"photo_100,omitempty"`
}

type Market struct {
	Type                     string   `json:"type,omitempty"`
	ContactID                int      `json:"contact_id,omitempty"`
	Currency                 Currency `json:"currency,omitempty"`
	Symbol                   string   `json:"currency_text,omitempty"`
	IsShowHeaderItemsLink    int      `json:"is_show_header_items_link,omitempty"`
	Enabled                  int      `json:"enabled,omitempty"`
	PriceMax                 string   `json:"price_max,omitempty"`
	PriceMin                 string   `json:"price_min,omitempty"`
	UnViewedOrdersCount      int      `json:"unviewed_orders_count,omitempty"`
	CommunityManageEnabled   int      `json:"is_community_manage_enabled,omitempty"`
	HasNotInMarketTab        bool     `json:"has_not_in_market_tab,omitempty"`
	HasModerationRejectedTab bool     `json:"has_moderation_rejected_tab,omitempty"`
}

type Currency struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Symbol string `json:"title,omitempty"`
}

type Place struct {
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Latitude  int    `json:"latitude,omitempty"`
	Longitude int    `json:"longitude,omitempty"`
	Type      string `json:"type,omitempty"`
	Country   int    `json:"country,omitempty"`
	City      int    `json:"city,omitempty"`
	Address   string `json:"address,omitempty"`
}
