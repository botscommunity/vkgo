package responses

type Users struct {
	Error *Error `json:"error"`
	Users []User `json:"response,omitempty"`
}

type User struct {
	Error                *Error           `json:"error"`
	ID                   int              `json:"id,omitempty"`
	Name                 string           `json:"first_name,omitempty"`
	Surname              string           `json:"last_name,omitempty"`
	Deactivated          string           `json:"deactivated,omitempty"`
	Closed               bool             `json:"is_closed"`
	CanAccessClosed      bool             `json:"can_access_closed"`
	About                string           `json:"about,omitempty"`
	Activities           string           `json:"activities,omitempty"`
	BirthDate            string           `json:"bdate,omitempty"`
	Banned               int              `json:"blacklisted,omitempty"`
	BannedByMe           int              `json:"blacklisted_by_me,omitempty"`
	Books                string           `json:"books,omitempty"`
	CanPost              int              `json:"can_post,omitempty"`
	CanSeePosts          int              `json:"can_see_all_posts,omitempty"`
	CanSeeAudio          int              `json:"can_see_audio,omitempty"`
	CanSendFriendRequest int              `json:"can_send_friend_request,omitempty"`
	CanWriteMessage      int              `json:"can_write_private_message,omitempty"`
	Career               []UserCareer     `json:"career,omitempty"`
	City                 UserCity         `json:"city,omitempty"`
	MutualFriends        int              `json:"common_count,omitempty"`
	Connections          any              `json:"connections,omitempty"`
	Contacts             UserContacts     `json:"contacts,omitempty"`
	Counters             UserCounters     `json:"counters,omitempty"`
	Country              UserCountry      `json:"country,omitempty"`
	CropPhoto            CropPhoto        `json:"crop_photo,omitempty"`
	Domain               string           `json:"domain,omitempty"`
	Education            UserEducation    `json:"education,omitempty"`
	Exports              any              `json:"exports,omitempty"`
	NameNominative       string           `json:"first_name_nom,omitempty"`
	NameGenitive         string           `json:"first_name_gen,omitempty"`
	NameDative           string           `json:"first_name_dat,omitempty"`
	NameAccusative       string           `json:"first_name_acc,omitempty"`
	NameInstrumental     string           `json:"first_name_ins,omitempty"`
	NameAblative         string           `json:"first_name_abl,omitempty"`
	FollowersCount       int              `json:"followers_count,omitempty"`
	Games                string           `json:"games,omitempty"`
	HasMobile            int              `json:"has_mobile,omitempty"`
	HasPhoto             int              `json:"has_photo,omitempty"`
	HomeTown             string           `json:"home_town,omitempty"`
	Interests            string           `json:"interests,omitempty"`
	IsFavorite           int              `json:"is_favorite,omitempty"`
	IsFriend             int              `json:"is_friend,omitempty"`
	IsHiddenFromFeed     int              `json:"is_hidden_from_feed,omitempty"`
	IsNoIndex            int              `json:"is_no_index,omitempty"`
	SurnameNominative    string           `json:"last_name_nom,omitempty"`
	SurnameGenitive      string           `json:"last_name_gen,omitempty"`
	SurnameDative        string           `json:"last_name_dat,omitempty"`
	SurnameAccusative    string           `json:"last_name_acc,omitempty"`
	SurnameInstrumental  string           `json:"last_name_ins,omitempty"`
	SurnameAblative      string           `json:"last_name_abl,omitempty"`
	LastSeen             UserLastSeen     `json:"last_seen,omitempty"`
	Lists                string           `json:"lists,omitempty"`
	MaidenName           string           `json:"maiden_name,omitempty"`
	Military             []UserMilitary   `json:"military,omitempty"`
	Movies               string           `json:"movies,omitempty"`
	Music                string           `json:"music,omitempty"`
	Nickname             string           `json:"nickname,omitempty"`
	Occupation           UserOccupation   `json:"occupation,omitempty"`
	Online               int              `json:"online,omitempty"`
	Personal             UserPersonal     `json:"personal,omitempty"`
	Photo50              string           `json:"photo_50,omitempty"`
	Photo100             string           `json:"photo_100,omitempty"`
	PhotoOriginal200     string           `json:"photo_200_orig,omitempty"`
	Photo200             string           `json:"photo_200,omitempty"`
	PhotoOriginal400     string           `json:"photo_400_orig,omitempty"`
	PhotoID              string           `json:"photo_id,omitempty"`
	PhotoMax             string           `json:"photo_max,omitempty"`
	PhotoOriginalMax     string           `json:"photo_max_orig,omitempty"`
	Quotes               string           `json:"quotes,omitempty"`
	Relatives            []UserRelatives  `json:"relatives,omitempty"`
	Relation             int              `json:"relation,omitempty"`
	Schools              []UserSchools    `json:"schools,omitempty"`
	ScreenName           string           `json:"screen_name,omitempty"`
	Gender               int              `json:"sex,omitempty"`
	Site                 string           `json:"site,omitempty"`
	Status               string           `json:"status,omitempty"`
	Timezone             any              `json:"timezone,omitempty"`
	Trending             int              `json:"trending,omitempty"`
	TV                   string           `json:"tv,omitempty"`
	Universities         UserUniversities `json:"universities,omitempty"`
	Verified             int              `json:"verified,omitempty"`
	WallDefault          string           `json:"wall_default,omitempty"`
}

type UserCareer struct {
	GroupID   int    `json:"group_id,omitempty"`
	Company   string `json:"company,omitempty"`
	CountryID int    `json:"country_id,omitempty"`
	CityID    int    `json:"city_id,omitempty"`
	CityName  string `json:"city_name,omitempty"`
	From      int    `json:"from,omitempty"`
	Until     int    `json:"until,omitempty"`
	Position  string `json:"position,omitempty"`
}

type UserCity struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type UserContacts struct {
	MobilePhone string `json:"mobile_phone,omitempty"`
	HomePhone   string `json:"home_phone,omitempty"`
}

type UserCounters struct {
	Albums        int `json:"albums,omitempty"`
	Videos        int `json:"videos,omitempty"`
	Audios        int `json:"audios,omitempty"`
	Photos        int `json:"photos,omitempty"`
	Notes         int `json:"notes,omitempty"`
	Friends       int `json:"friends,omitempty"`
	Gifts         int `json:"gifts,omitempty"`
	Groups        int `json:"groups,omitempty"`
	OnlineFriends int `json:"online_friends,omitempty"`
	MutualFriends int `json:"mutual_friends,omitempty"`
	UserVideos    int `json:"user_videos,omitempty"`
	UserPhotos    int `json:"user_photos,omitempty"`
	Followers     int `json:"followers,omitempty"`
	Pages         int `json:"pages,omitempty"`
	Subscriptions int `json:"subscriptions,omitempty"`
}

type UserCountry struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type UserEducation struct {
	University     int    `json:"university,omitempty"`
	UniversityName string `json:"university_name,omitempty"`
	Faculty        int    `json:"faculty,omitempty"`
	FacultyName    string `json:"faculty_name,omitempty"`
	Graduation     int    `json:"graduation,omitempty"`
}

type UserLastSeen struct {
	Time     int `json:"time,omitempty"`
	Platform int `json:"platform,omitempty"`
}

type UserMilitary struct {
	Unit      string `json:"unit,omitempty"`
	UnitID    int    `json:"unit_id,omitempty"`
	CountryID int    `json:"country_id,omitempty"`
	From      int    `json:"from,omitempty"`
	Until     int    `json:"until,omitempty"`
}

type UserOccupation struct {
	Type string `json:"type,omitempty"`
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type UserPersonal struct {
	Political  int      `json:"political,omitempty"`
	Languages  []string `json:"langs,omitempty"`
	Religion   string   `json:"religion,omitempty"`
	Inspired   string   `json:"inspired_by,omitempty"`
	PeopleMain int      `json:"people_main,omitempty"`
	LifeMain   int      `json:"life_main,omitempty"`
	Smoking    int      `json:"smoking,omitempty"`
	Alcohol    int      `json:"alcohol,omitempty"`
}

type UserRelatives struct {
	ID   int `json:"id,omitempty"`
	Name int `json:"name,omitempty"`
	Type int `json:"type,omitempty"`
}

type UserSchools struct {
	ID            string `json:"id,omitempty"`
	Country       int    `json:"country,omitempty"`
	City          int    `json:"city,omitempty"`
	Name          string `json:"name,omitempty"`
	YearFrom      int    `json:"year_from,omitempty"`
	YearTo        int    `json:"year_to,omitempty"`
	YearGraduated int    `json:"year_graduated,omitempty"`
	Class         string `json:"class,omitempty"`
	Speciality    string `json:"speciality,omitempty"`
	Type          int    `json:"type,omitempty"`
	TypeName      string `json:"type_str,omitempty"`
}

type UserUniversities []UserUniversity
type UserUniversity struct {
	ID              int    `json:"id,omitempty"`
	Country         int    `json:"country,omitempty"`
	City            int    `json:"city,omitempty"`
	Name            string `json:"name,omitempty"`
	Faculty         int    `json:"faculty,omitempty"`
	FacultyName     string `json:"faculty_name,omitempty"`
	Chair           int    `json:"chair,omitempty"`
	ChairName       string `json:"chair_name,omitempty"`
	Graduation      int    `json:"graduation,omitempty"`
	EducationForm   string `json:"education_form,omitempty"`
	EducationStatus string `json:"education_status,omitempty"`
}
