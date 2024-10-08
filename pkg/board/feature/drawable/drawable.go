package drawable

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Visualiser/pkg/board/feature/drawable/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Drawable struct {
	rectangles []components.Rectangle
	triangles  []components.Triangle
	circles    []components.Circle
	color      rl.Color
}

func New(color rl.Color) Drawable {
	return Drawable{
		rectangles: make([]components.Rectangle, 0),
		triangles:  make([]components.Triangle, 0),
		circles:    make([]components.Circle, 0),
		color:      color,
	}
}

func (drawable *Drawable) AddRectangle(offsetOnTile rl.Vector2, size rl.Vector2) {
	drawable.rectangles = append(drawable.rectangles, components.NewRectangle(offsetOnTile, size))
}

func (drawable *Drawable) AddTriangle(offsetsOnTile []rl.Vector2) {
	drawable.triangles = append(drawable.triangles, components.NewTriangle(offsetsOnTile))
}

func (drawable *Drawable) AddCircle(offsetOnTile rl.Vector2, radius float32) {
	drawable.circles = append(drawable.circles, components.NewCircle(offsetOnTile, radius))
}

func (drawable Drawable) Draw(tilePosition rl.Vector2) {
	for _, rectangle := range drawable.rectangles {
		rectangle.Draw(tilePosition, drawable.color)
	}
	for _, triangle := range drawable.triangles {
		triangle.Draw(tilePosition, drawable.color)
	}
	for _, circle := range drawable.circles {
		circle.Draw(tilePosition, drawable.color)
	}
}
