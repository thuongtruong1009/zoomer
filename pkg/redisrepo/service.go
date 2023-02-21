package redisrepo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
	"zoomer/internal/models"
	"zoomer/db"
	"github.com/go-redis/redis/v8"
)

func RegisterNewUser(username,password string) error {
	err := db.RedisClient.Set(context.Background(), username, password, 0).Err()
	if err != nil{
		log.Println("error while adding new user", err)
		return err
	}

	err = db.RedisClient.SAdd(context.Background(), userSetKey(),username).Err()
	if err != nil {
		log.Println("error while adding new user to set", err)
		db.RedisClient.Del(context.Background(), username)
		return err
	}
	return nil
}

func IsUserExist(username string) bool {
	return db.RedisClient.SIsMember(context.Background(), userSetKey(), username).Val()
}

func IsUserAuthentic(username, password string) error {
	p:= db.RedisClient.Get(context.Background(), username).Val()

	if !strings.EqualFold(p, password) {
		return fmt.Errorf("invalid username or password")
	}
	return nil
}

func UpdateContactList(username, contact string) error {
	zs := &redis.Z{Score: float64(time.Now().Unix()), Member: contact}

	err := db.RedisClient.ZAdd(context.Background(), contactListZKey(username), zs).Err()

	if err != nil {
		log.Println("error while updating contact list username: ", username, "contact: ", contact, err)
		return err
	}
	return nil
}

func CreateChat(c *models.Chat) (string, error) {
	chatKey := chatKey()
	fmt.Println("chat key: ", chatKey)

	by, _ := json.Marshal(c)

	res, err := db.RedisClient.Do(context.Background(), "JSON.SET", chatKey, "$", string(by),).Result()

	if err !=nil {
		log.Println("error while creating chat", err)
		return "", err
	}

	log.Println("chat created: ", res)

	err = UpdateContactList(c.From, c.To)
	if err != nil {
		log.Println("error while updating contact list of", c.From)
	}

	err = UpdateContactList(c.To, c.From)
	if err != nil {
		log.Println("error while updating contact list of", c.To)
	}

	return chatKey, nil
}

func CreateFetchChatBetweenIndex() {
	res, err := db.RedisClient.Do(context.Background(), "FT.CREATE", "chatIndex", "ON", "JSON", "PREFIX", "1", "chat#", "SCHEMA", "$.from", "AS", "from", "$.to", "AS", "to", "TAG", "timestamp", "NUMERIC", "SORTABLE").Result()

	fmt.Println(res, err)
}

func FetchChatBetween(username1, username2, fromTS, toTS string)([]models.Chat, error) {
	query := fmt.Sprintf("@from:{%s|%s} @to:{%s|%s} @timestamp:[%s %s]", username1, username2, username1,username2, fromTS, toTS)

	res, err := db.RedisClient.Do(context.Background(), "FT.SEARCH", chatIndex(), query, "SORTBY", "timestamp", "DESC").Result()

	if err != nil {
		return nil, err
	}

	data := Deserialise(res)

	chats := DeserialiseChat(data)
	return chats, nil
}

func FetchContactList(username string) ([]models.ContactList, error) {
	zRangeArg := redis.ZRangeArgs {
		Key: contactListZKey(username),
		Start: 0,
		Stop: -1,
		Rev: true,
	}

	res, err := db.RedisClient.ZRangeArgsWithScores(context.Background(), zRangeArg).Result()

	if err != nil {
		log.Println("error while fetching contact list. username: ", username, err)
		return nil, err
	}

	contactList := DeserialiseContactList(res)
	return contactList, nil
}