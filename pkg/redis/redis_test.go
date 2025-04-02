package redis

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestNewRedis(t *testing.T) {
	client := InitRedis()
	// Test connection
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Redis connected:", pong)
}
