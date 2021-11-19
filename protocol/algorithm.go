package protocol

type FaceInputProto struct {
	Id     int64  `json:"id"`
	FaceId string `json:"feature"`
}

type RecordInputProto struct {
	FaceId string `json:"feature"`
}

type ImageForwardProto struct {
	Img    string  `json:"img"`
	X1     float64 `json:"x1"`
	X2     float64 `json:"x2"`
	Y1     float64 `json:"y1"`
	Y2     float64 `json:"y2"`
	IsMask bool    `json:"mask"`
}
