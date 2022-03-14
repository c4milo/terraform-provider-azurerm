package connections

type ConnectionError struct {
	Etag       *string                    `json:"etag,omitempty"`
	Id         *string                    `json:"id,omitempty"`
	Location   *string                    `json:"location,omitempty"`
	Name       *string                    `json:"name,omitempty"`
	Properties *ConnectionErrorProperties `json:"properties,omitempty"`
	Tags       *map[string]string         `json:"tags,omitempty"`
	Type       *string                    `json:"type,omitempty"`
}
