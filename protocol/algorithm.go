package protocol

type FaceInputProto struct {
	Id     int    `json:"id"`
	FaceId string `json:"face_id"`
}

type RecordInputProto struct {
	FaceId string `json:"face_id"`
}

type ImageForwardProto struct {
	Img    string  `json:"img"`
	X1     float64 `json:"x1"`
	X2     float64 `json:"x2"`
	X3     float64 `json:"x3"`
	X4     float64 `json:"x4"`
	IsMask bool    `json:"is_mask"`
}
