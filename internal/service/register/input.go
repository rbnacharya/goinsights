package register

type RegisterClickInput struct {
	CustomerID int    `json:"customerID" validate:"required"`
	TagID      int    `json:"tagId" validate:"required"`
	UserID     string `json:"userId" validate:"required"`
	Timestamp  int64  `json:"timestamp" validate:"required,gte=0"`
	RemoteIP   string `json:"remoteIp" validate:"ipv4"`
	UserAgent  string `json:"userAgent" validate:"required"`
}
