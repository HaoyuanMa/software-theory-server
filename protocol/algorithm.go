package protocol

type FaceInputProto struct {
	Id     int    `json:"id"`
	FaceId string `json:"face_id"`
}

type RecordInputProto struct {
	FaceId string `json:"face_id"`
}
