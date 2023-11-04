package update

type AttachmentAudio Audio

type Audio struct {
	ID       int    `json:"id,omitempty"`
	OwnerID  int    `json:"owner_id,omitempty"`
	Artist   string `json:"artist,omitempty"`
	Title    string `json:"title,omitempty"`
	Duration int    `json:"duration,omitempty"`
	URL      string `json:"url,omitempty"`
	LyricsID int    `json:"lyrics_id,omitempty"`
	AlbumID  int    `json:"album_id,omitempty"`
	GenreID  int    `json:"genre_id,omitempty"`
	Date     int    `json:"date,omitempty"`
	NoSearch int    `json:"no_search,omitempty"`
	IsHQ     int    `json:"is_hq,omitempty"`

	IsExplicit          bool          `json:"is_explicit,omitempty"`
	IsFocusTrack        bool          `json:"is_focus_track,omitempty"`
	TrackCode           string        `json:"track_code,omitempty"`
	Artists             []AudioArtist `json:"main_artists,omitempty"`
	ShortVideosAllowed  bool          `json:"short_videos_allowed,omitempty"`
	StoriesAllowed      bool          `json:"stories_allowed,omitempty"`
	StoriesCoverAllowed bool          `json:"stories_cover_allowed,omitempty"`
}

type AudioArtist struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Followed  bool   `json:"is_followed,omitempty"`
	CanFollow bool   `json:"can_follow,omitempty"`
}
