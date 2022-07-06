package database

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func GetImages() (images []Image) {

	keys, err := Client.Keys("*").Result()

	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		image := GetImage(key)
		images = append(images, image)
	}

	return images
}

func GetImage(id string) (image Image) {

	value, err := Client.Get(id).Result()

	if err != nil {
		panic(err)
	}

	if err != redis.Nil {
		err = json.Unmarshal([]byte(value), &image)
	}

	return image
}

func SaveImage(image Image) {
	imageBytes, err := json.Marshal(image)
	if err != nil {
		panic(err)
	}

	err = Client.Set(image.Id, imageBytes, 0).Err()
	if err != nil {
		panic(err)
	}
}

func SaveImages(images []Image) {
	for _, image := range images {
		SaveImage(image)
	}
}

func Database() {
	redis_sentinels := os.Getenv("REDIS_SENTINELS")
	redis_master := os.Getenv("REDIS_MASTER_NAME")
	redis_password := os.Getenv("REDIS_PASSWORD")
	sentinelAddr := strings.Split(redis_sentinels, ",")

	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    redis_master,
		SentinelAddrs: sentinelAddr,
		Password:      redis_password,
		DB:            0, // use default DB
	})

	Client = client

	_, err := client.Ping().Result() // ping database to check for connection
	if err != nil {
		panic(err)
	}
}

// import (
// 	"github.com/go-redis/redis"
// 	"github.com/turnixxd/grpc-rest-graphql-comparison/go/env"
// )

// func createRedisDatabase() (Database, error) {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     env.Process("REDIS_ADDR"),
// 		Password: "",
// 		DB:       0, // use default DB
// 	})
// 	_, err := client.Ping().Result() // ping database to check for connection
// 	if err != nil {
// 		return nil, &CreateDatabaseError{}
// 	}
// 	return &redisDatabase{client: client}, nil
// }

// func (r *redisDatabase) Set(username, password, role string) (string, error) {
// 	_, err := r.client.Set(username, password, role, 0).Result()
// 	if err != nil {
// 		return generateError("set", err)
// 	}
// 	return username, nil
// }

// func (r *redisDatabase) Get(id int32, username, password, role string) (string, error) {
// 	value, err := r.client.Get(id, username, password, role, 0).Result()
// 	if err != nil {
// 		return generateError("get", err)
// 	}
// 	return value, nil
// }

// func (r *redisDatabase) Delete(id int, username, password, role string) (string, error) {
// 	_, err := r.client.Del(key).Result()
// 	if err != nil {
// 		return generateError("delete", err)
// 	}
// 	return id, nil
// }

// func generateError(operation string, err error) (string, error) {
// 	if err == redis.Nil {
// 		return "", &OperationError{operation}
// 	}
// 	return "", &DownError{}
// }
