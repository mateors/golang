package shortener

type Redirect struct {
	Code       string `json:"code" bson:"code" msgpack:"code"`
	URL        string `json:"url" bson:"url" msgpack:"url" validate:"empty=false & format=url"`
	CreateTime int64  `json:"create_time" bson:"create_time" msgpack:"create_time"`
}
