package plex

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e Error) Error() string {
	return error(e).Error()
}
