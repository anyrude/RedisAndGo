//# RedisAndGo
//# run redis server and client


package main

import "gopkg.in/redis.v5"
import "time"

func main() {
 var client *redis.Client

client = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		//Password: "",
		//DB: 0,
	})
//	client.FlushDB()


 key:="k"
 value:="ckcckck"
  err := client.Set(key, value, 0).Err()
if err != nil {
    panic(err)
}

tim := 30 * time.Second
 key1 := "l"
 value1 := "ckcckck"
  err1 := client.Set(key1, value1, tim).Err()
if err1 != nil {
    panic(err1)
}


//only if list is already present named as 'lis'


//key2 := "m"
key3:="lis"
var b int64
b= 2
a:= "ccccccc"
err2 := client.LSet(key3, b, a).Err()
if err2 != nil {
    panic(err2)
}

//only if hash is already present named as 'hashyo' with key 'no1'

key4:="hashyo"
bb:="no1"
ab := "ccccccccccccc"
err3 := client.HSet(key4, bb, ab).Err()
if err3 != nil {
    panic(err3) 

}
}
