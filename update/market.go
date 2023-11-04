package update

type MarketComment struct {
	PostComment

	OwnerID int `json:"market_owner_id,omitempty"`
	ItemID  int `json:"item_id,omitempty"`
}

type DeleteMarketComment struct {
	OwnerID   int `json:"owner_id,omitempty"`
	ID        int `json:"id,omitempty"`
	UserID    int `json:"user_id,omitempty"`
	DeleterID int `json:"deleter_id,omitempty"`
	ItemID    int `json:"item_id,omitempty"`
}

type OrderMarket struct {
	ID                 int                `json:"id,omitempty"`
	GroupID            int                `json:"group_id,omitempty"`
	UserID             int                `json:"user_id,omitempty"`
	Date               int                `json:"date,omitempty"`
	VariantsGroupingID int                `json:"variants_grouping_id,omitempty"`
	IsMainVariant      bool               `json:"is_main_variant,omitempty"`
	PropertyValues     []PropertyValues   `json:"property_values,omitempty"`
	CartQuantity       int                `json:"cart_quantity,omitempty"`
	Status             int                `json:"status,omitempty"`
	ItemsCount         int                `json:"items_count,omitempty"`
	TotalPrice         TotalPrice         `json:"total_price,omitempty"`
	DisplayOrderID     string             `json:"display_order_id,omitempty"`
	Comment            string             `json:"comment,omitempty"`
	PreviewOrderItems  []AttachmentMarket `json:"preview_order_items,omitempty"`
	Delivery           OrderDelivery      `json:"delivery,omitempty"`
	Recipient          Recipient          `json:"recipient,omitempty"`
}

type PropertyValues struct {
	VariantID    int    `json:"variant_id,omitempty"`
	VariantName  string `json:"variant_name,omitempty"`
	PropertyName string `json:"property_name,omitempty"`
}

type TotalPrice struct {
	Amount   string   `json:"amount,omitempty"`
	Currency Currency `json:"currency,omitempty"`
	Text     string   `json:"text,omitempty"`
}

type OrderDelivery struct {
	Address       string `json:"address,omitempty"`
	Type          string `json:"type,omitempty"`
	TrackNumber   string `json:"track_number,omitempty"`
	TrackLink     string `json:"track_link,omitempty"`
	DeliveryPoint any    `json:"delivery_point,omitempty"`
}

type Recipient struct {
	Name        string `json:"name,omitempty"`
	Phone       string `json:"phone,omitempty"`
	DisplayText string `json:"display_text,omitempty"`
}

type AttachmentMarket struct {
	ID          int               `json:"id,omitempty"`
	OwnerID     int               `json:"owner_id,omitempty"`
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	Price       Price             `json:"price,omitempty"`
	Dimensions  Dimensions        `json:"dimensions,omitempty"`
	Weight      int               `json:"weight,omitempty"`
	Category    Category          `json:"category,omitempty"`
	Cover       string            `json:"thumb_photo,omitempty"`
	Date        int               `json:"date,omitempty"`
	Access      int               `json:"availability,omitempty"`
	IsFavorite  bool              `json:"is_favorite,omitempty"`
	Vendor      string            `json:"sku,omitempty"`
	Photos      []AttachmentPhoto `json:"photos,omitempty"`
	CanComment  int               `json:"can_comment,omitempty"`
	CanRepost   int               `json:"can_repost,omitempty"`
	Likes       []Likes           `json:"likes,omitempty"`
	URL         string            `json:"url,omitempty"`
	ButtonTitle string            `json:"button_title,omitempty"`
}

type Price struct {
	Amount       string   `json:"amount,omitempty"`
	Currency     Currency `json:"currency,omitempty"`
	UnDiscounted string   `json:"old_amount,omitempty"`
	Text         string   `json:"text,omitempty"`
}

type Currency struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Dimensions struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
	Length int `json:"length,omitempty"`
}

type Category struct {
	ID      int     `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Section Section `json:"section,omitempty"`
}

type Section struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
