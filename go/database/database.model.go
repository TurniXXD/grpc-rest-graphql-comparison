package database

type Image struct {
	Id     string
	Url    string
	Width  int
	Height int
	Alt    string
}

type User struct {
	Id       string
	username string
	password string
	role     string
}

// type Database interface {
// 	Set(username, password, role string) (string, error)
// 	Get(id int32, username, password, role string) (string, error)
// 	Delete(id int32, username, password, role string) (string, error)
// }

// func Factory(databaseName string) (Database, error) {
// 	switch databaseName {
// 	case "redis":
// 		return createRedisDatabase()
// 	default:
// 		return nil, &NotImplementedDatabaseError{databaseName}
// 	}
// }
