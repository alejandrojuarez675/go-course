package main

import "fmt"

// person students has a 1-1 relation
type Person struct {
	name     string
	lastname string
	age      int
	email    string
}

type Student struct {
	user   Person
	code   string
	active bool
}

// course has a one to many relation to videos
type Course struct {
	title  string
	videos []*Video
}

type Video struct {
	title  string
	url    string
	course *Course
}

func main() {

	v1 := Video{title: "first video", url: "https://first.com"}
	v2 := Video{title: "second video", url: "https://second.com"}

	course := Course{
		title:  "go course",
		videos: []*Video{&v1, &v2},
	}

	v1.course = &course
	v2.course = &course

	fmt.Println(course)

	for _, video := range course.videos {
		fmt.Println(video)
	}

}
