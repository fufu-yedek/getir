package memrecords

type SingleRecordResponse struct {
	// example: active-tabs
	Key string `json:"key"`
	// example: getir
	Value string `json:"value"`
}

type RecordSerializer struct {
	Record Record
}

func (r RecordSerializer) Response() interface{} {
	return SingleRecordResponse{
		Key:   r.Record.Key,
		Value: r.Record.Value,
	}
}
