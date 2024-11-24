package main

func Hello(color string) string {
	//replicate Hello depending of colour
	if color == "black" {
		return "Hello Shadow"
	} else if color == "blue" {
		return "Hello Sonic"
	}
	return "You're not even good enough to be my faker, you " + color
}
