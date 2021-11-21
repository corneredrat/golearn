package main

import "fmt"

func main() {
	var moviesILike []string
	var allTimeFavouriteMovie string

	moviesILike = []string{"Ulidavaru Kandanthe", "Kingdom of the Crystal Skull", "The Guest", "Shawshanl Redemption", "Shutter Island"}
	allTimeFavouriteMovie = "Raiders of the lost Ark"
	for _, movie := range moviesILike {
		func(myFavouriteMovie string) {
			fmt.Printf("%v is one of my favorite movies; %v is my all time favorite movie.\n", myFavouriteMovie, allTimeFavouriteMovie)
		}(movie)
	}
}
