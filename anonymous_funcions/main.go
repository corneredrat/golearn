package main

import "fmt"

func main() {
	var moviesILike []string
	var allTimeFavouriteMovie string

	moviesILike = []string{"Ulidavaru Kandanthe", "Kingdom of the Crystal Skull", "The Guest", "Shawshanl Redemption", "Shutter Island"}
	allTimeFavouriteMovie = "Raiders of the lost Ark"
	for _, movie := range moviesILike {
		func(myFavoriteMovie string) {
			fmt.Printf("%v is one of my favorite movies; %v is my all time favorite movie.\n", myFavoriteMovie, allTimeFavouriteMovie)
			myFavoriteMovie = "Some generic Marvel Movie"
		}(movie)
		fmt.Printf("%v is a movie I would watch again\n", movie)
	}
}
