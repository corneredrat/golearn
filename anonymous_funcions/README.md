# Anonymous functions

- If anonymous function modifies variable that is outside of its scope, then the variable is changed for the external scope too!

```golang
package main

import "fmt"

func main() {
	var moviesILike []string
	var allTimeFavouriteMovie string

	moviesILike = []string{"Ulidavaru Kandanthe", "Kingdom of the Crystal Skull", "The Guest", "Shawshanl Redemption", "Shutter Island"}
	allTimeFavouriteMovie = "Raiders of the lost Ark"
	for _, movie := range moviesILike {
		func() {
			fmt.Printf("%v is one of my favorite movies; %v is my all time favorite movie.\n", movie, allTimeFavouriteMovie)
			movie = "A generic Marvel Movie"
		}()
		fmt.Printf("%v is a movie I would watch again\n", movie)
	}
}
```

output
```
Ulidavaru Kandanthe is one of my favorite movies; Raiders of the lost Ark is my all time favorite movie.
A generic Marvel Movie is a movie I would watch again
Kingdom of the Crystal Skull is one of my favorite movies; Raiders of the lost Ark is my all time favorite movie.
A generic Marvel Movie is a movie I would watch again
The Guest is one of my favorite movies; Raiders of the lost Ark is my all time favorite movie.
A generic Marvel Movie is a movie I would watch again
Shawshanl Redemption is one of my favorite movies; Raiders of the lost Ark is my all time favorite movie.
A generic Marvel Movie is a movie I would watch again
Shutter Island is one of my favorite movies; Raiders of the lost Ark is my all time favorite movie.
A generic Marvel Movie is a movie I would watch again
```

To avoid this, pass the scoped variables as parameter to the anonymous function
```golang
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
			myFavoriteMovie = "A generic Marvel Movie"
		}(movie)
		fmt.Printf("%v is a movie I would watch again\n", movie)
	}
}
```

output
```
Ulidavaru Kandanthe is one of my favorite movies; Raiders of the lost Ark is my all time favorite movie.
Ulidavaru Kandanthe is a movie I would watch again
Kingdom of the Crystal Skull is one of my favorite movies; Raiders of the lost Ark is my all time favorite movie.
Kingdom of the Crystal Skull is a movie I would watch again
The Guest is one of my favorite movies; Raiders of the lost Ark is my all time favorite movie.
The Guest is a movie I would watch again
Shawshanl Redemption is one of my favorite movies; Raiders of the lost Ark is my all time favorite movie.
Shawshanl Redemption is a movie I would watch again
Shutter Island is one of my favorite movies; Raiders of the lost Ark is my all time favorite movie.
Shutter Island is a movie I would watch again
```