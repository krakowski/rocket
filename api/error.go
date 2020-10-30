package rocket

// The last error returned from the RocketChat API
var LastError Error

type Error struct {
	Status	string	`json:"status"`
	Error	string	`json:"status"`
	Message	string	`json:"status"`
}