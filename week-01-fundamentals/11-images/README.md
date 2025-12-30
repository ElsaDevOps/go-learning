## Images
 Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.

Define your own Image type, implement the necessary methods, and call pic.ShowImage.

Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

ColorModel should return color.RGBAModel.

At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one. 

## What I learned

- **Implementing interfaces:** `image.Image` requires three methods: `ColorModel()`, `Bounds()`, `At(x, y int)`
- **Struct fields:** Store the dimensions your methods need to reference
- **Return types matter:** `At` returns `color.Color` (the interface), not `color.RGBA` — but `color.RGBA` satisfies the interface
- **Type conversions:** `color.RGBA` needs `uint8` values, so convert your calculation
- **Imports:** When using types from packages, you need to import all the packages you're using, not just the one the exercise provides

## Running
go run main.go