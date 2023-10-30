package redis

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/go-redis/redis/v8"
	"grapefruit/config"
	"grapefruit/kit/log"
	"io/ioutil"
	"os"
	"time"
)

var RDB *redis.Client

//GetRedis 集群的情况下不可以变更db
func GetRedis(db ...int) *redis.Client {
	redisConf := config.GetRedisConfig()
	//dir, _ := os.Getwd()
	//fmt.Println(os.Getwd())
	//lastIndex := strings.Index(dir, config.ProjectModuleName)
	//tmp := dir[:lastIndex] + config.ProjectModuleName + redisConf.CrtPath
	//fmt.Println(tmp)
	caCert, err := ioutil.ReadFile(redisConf.CrtPath)
	if err != nil {
		log.Fatal("Failed to load CA certificate: %v", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tLSConfig := &tls.Config{
		RootCAs:            caCertPool,
		InsecureSkipVerify: false, // Ensure this is set to 'false' to perform server's certificate chain and hostname verification.
	}
	dbIndex := 0
	if len(db) != 0 {
		dbIndex = db[0]
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:      redisConf.Host,
		Username:  redisConf.User,
		Password:  redisConf.Pwd, // 没有密码，默认值
		DB:        dbIndex,       // 默认DB 0
		TLSConfig: tLSConfig,
	})
	rdb.Ping(context.Background()).Val()
	RDB = rdb
	return rdb
}

func ParseRedisOption() *redis.Options {
	opt, err := redis.ParseURL(os.Getenv("REDIS_CONN_STRING"))
	if err != nil {
		log.Error(err.Error())
	}
	return opt
}

func RedisSet(key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	return RDB.Set(ctx, key, value, expiration).Err()
}

func RedisGet(key string) (string, error) {
	ctx := context.Background()
	return RDB.Get(ctx, key).Result()
}

func RedisDel(key string) error {
	ctx := context.Background()
	return RDB.Del(ctx, key).Err()
}

func RedisDecrease(key string, value int64) error {
	ctx := context.Background()
	return RDB.DecrBy(ctx, key, value).Err()
}
