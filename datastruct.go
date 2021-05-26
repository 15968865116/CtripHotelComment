package apiiget

type AllStruct struct {
	Responsestatus ResponseStatus `json:"ResponseStatus"`
	Responses      Response       `json:"Response"`
	ErrorCodes     int            `json:"ErrorCode"`
}

type ResponseStatus struct {
	Timestamp     string      `json:"Timestamp"`
	Ack           string      `json:"Ack"`
	Errors        []string    `json:"Errors"`
	ExtensionList []Extension `json:"Extension"`
}

type Extension struct {
	Id    string `json:"Id"`
	Value string `json:"Value"`
}

type Response struct {
	ReviewBaseInfos  ReviewBaseInfo `json:"ReviewBaseInfo"`
	ReviewList       []Review       `json:"ReviewList"`
	ReviewFilterList []ReviewFilter `json:"ReviewFilterList"`
}

type ReviewBaseInfo struct {
	Categoryscore               []CategoryScore `json:"categoryScore"`
	RecommendPercent            string          `json:"recommendPercent"`
	Score                       string          `json:"score"`
	TotalReviews                int             `json:"totalReviews"`
	ScoreMax                    int             `json:"scoreMax"`
	ScoreDesc                   string          `json:"scoreDesc"`
	TotalReviewsTA              int             `json:"totalReviewsTA"`
	AllTotalReviews             int             `json:"allTotalReviews"`
	CtripTotalReviews           int             `json:"ctripTotalReviews"`
	CtripTotalReviewsForPage    int             `json:"ctripTotalReviewsForPage"`
	TotalUnusefulReviewsForPage int             `json:"totalUnusefulReviewsForPage"`
}

type CategoryScore struct {
	ScoreName string `json:"scoreName"`
	ItemScore string `json:"itemScore"`
}

type ReviewFilter struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Type  string `json:"type"`
	Count string `json:"count"`
}

type Review struct {
	Userprofile   UserProfile   `json:"userProfile"`
	Reviewdetails ReviewDetails `json:"reviewDetails"`
	ReviewId      string        `json:"reviewId"`
	HelpfulCount  int           `json:"helpfulCount"`
	MarkUseful    bool          `json:"markUseful"`
	ShowFold      bool          `json:"showFold"`
}

type UserProfile struct {
	AvatarUrl     string `json:"avatarUrl"`
	UserName      string `json:"userName"`
	ReviewedCount int    `json:"reviewedCount"`
}

type ReviewDetails struct {
	Reviewscore        ReviewScore `json:"reviewScore"`
	ReviewContent      string      `json:"reviewContent"`
	ReleaseDate        string      `json:"releaseDate"`
	TravelType         string      `json:"travelType"`
	RoomType           string      `json:"roomType"`
	CheckInDate        string      `json:"checkInDate"`
	ReviewUpdateImages []string    `json:"reviewUpdateImages"`
	FeedBackList       []FeedBack  `json:"feedbackList"`
}

type ReviewScore struct {
	Score            string `json:"score"`
	ScoreMax         string `json:"scoreMax"`
	ScoreDescription string `json:"scoreDescription"`
}

type FeedBack struct {
	ReviewId   string `json:"reviewId"`
	Type       int    `json:"type"`
	CreateDate string `json:"createDate"`
	Content    string `json:"content"`
}
