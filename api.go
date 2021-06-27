package api_project

type ApiList struct {
	Id 			int    `json:"id"`
	Title 		string `string:"title"`
	Description string `string:"description"`
}

type UsersList struct {
	Id 		int
	UserId	int
	ListId  int
}