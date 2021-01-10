package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {

	fmt.Println("hello world")
	elliot := &Person{
		Name:   "Monu",
		Salary: 24,
		SocialFollowers: &SocialFollowers{
			Youtube: 2500,
			Twitter: 1400,
		},
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println(data)
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.GetSalary())
	fmt.Println("twitters ", newElliot.SocialFollowers.GetTwitter())
	fmt.Println("Youtube ", newElliot.SocialFollowers.GetYoutube())
}
