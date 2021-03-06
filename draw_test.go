package draw

import (
	"fmt"
	px "github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"image/color"
)

// ImdOp

func ExampleCircle() {
	circle := Circle(25, px.V(50, 100), 2)
	smallCircle := Circle(3, px.V(1, 2), 4)
	fmt.Println(circle.String())
	fmt.Println(smallCircle.String())
	// Output:
	// Circle radius 25 center Vec(50, 100) thickness 2
	// Circle radius 3 center Vec(1, 2) thickness 4
}

func ExampleLine() {
	line := Line(px.V(50, 100), px.V(101, 202), 2)
	fmt.Println(line.String())
	fmt.Println(Line(px.V(1, 2), px.V(3, 4), 5).String())
	// Output:
	// Line from Vec(50, 100) to Vec(101, 202) thickness 2
	// Line from Vec(1, 2) to Vec(3, 4) thickness 5
}

func ExampleRectangle() {
	rectangle := Rectangle(px.V(50, 100), px.V(101, 202), 0)
	fmt.Println(rectangle.String())
	fmt.Println(Rectangle(px.V(1, 2), px.V(3, 4), 5).String())
	// Output:
	// Rectangle from Vec(50, 100) to Vec(101, 202) (filled)
	// Rectangle from Vec(1, 2) to Vec(3, 4) thickness 5
}

func ExampleColor() {
	circle := Circle(25, px.V(50, 100), 2)
	smallCircle := Circle(3, px.V(1, 2), 4)
	green := color.RGBA{R: 0, G: 1, B: 0}
	fmt.Println(Colored(green, circle))
	white := color.RGBA{R: 1, G: 1, B: 1}
	fmt.Println(Colored(white, smallCircle))
	// Output:
	// Color {0 1 0 0}:
	//   Circle radius 25 center Vec(50, 100) thickness 2
	// Color {1 1 1 0}:
	//   Circle radius 3 center Vec(1, 2) thickness 4
}

func Example_imdOpSequence() {
	circle := Circle(25, px.V(50, 100), 2)
	smallCircle := Circle(3, px.V(1, 2), 4)
	fmt.Println(ImdOpSequence(circle, smallCircle).String())
	fmt.Println(ImdOpSequence(smallCircle, circle).String())
	// Output:
	// ImdOp Sequence:
	//   Circle radius 25 center Vec(50, 100) thickness 2
	//   Circle radius 3 center Vec(1, 2) thickness 4
	// ImdOp Sequence:
	//   Circle radius 3 center Vec(1, 2) thickness 4
	//   Circle radius 25 center Vec(50, 100) thickness 2
}

func Example_nestedSequence() {
	circle := Circle(25, px.V(50, 100), 2)
	smallCircle := Circle(3, px.V(1, 2), 4)
	fmt.Println(ImdOpSequence(ImdOpSequence(smallCircle, circle)).String())
	// Output:
	// ImdOp Sequence:
	//   ImdOp Sequence:
	//     Circle radius 3 center Vec(1, 2) thickness 4
	//     Circle radius 25 center Vec(50, 100) thickness 2
}

func Example_thenSequence() {
	sequence := ImdOpSequence().
		Then(Circle(25, px.V(50, 100), 2)).
		Then(Circle(3, px.V(1, 2), 4))
	fmt.Println(sequence.String())
	// Output:
	// ImdOp Sequence:
	//   Circle radius 25 center Vec(50, 100) thickness 2
	//   Circle radius 3 center Vec(1, 2) thickness 4
}

// TextOp

func ExampleText() {
	fmt.Println(Text("First line", "Second line"))
	// Output:
	// Text:
	//   First line
	//   Second line
}

// WinOp

func Example_liftImdOp() {
	fmt.Println(ToWinOp(Circle(5, px.V(0, 4), 1)).String())
	fmt.Println(ToWinOp(Line(px.V(0, 4), px.V(0, 4), 1)).String())
	// Output:
	// WinOp from ImdOp:
	//   Circle radius 5 center Vec(0, 4) thickness 1
	// WinOp from ImdOp:
	//   Line from Vec(0, 4) to Vec(0, 4) thickness 1
}

func Example_movedLineWinOp() {
	fmt.Print(Moved(px.V(50, 100.41001), ToWinOp(Line(px.V(0, 4), px.V(5, 6), 10))).String())
	// Output:
	// Moved 50 pixels right 100 pixels up:
	//   WinOp from ImdOp:
	//     Line from Vec(0, 4) to Vec(5, 6) thickness 10
}

func Example_movedRectangleWinOp() {
	fmt.Println(Moved(px.V(-1, -2), ToWinOp(Rectangle(px.V(0, 4), px.V(5, 6), 10))).String())
	// Output:
	// Moved 1 pixels left 2 pixels down:
	//   WinOp from ImdOp:
	//     Rectangle from Vec(0, 4) to Vec(5, 6) thickness 10
}

func Example_movedTileLayerWinOp() {
	fmt.Println(Moved(px.V(100, -80), TileLayer(nil, "Foreground")).String())
	// Output:
	// Moved 100 pixels right 80 pixels down:
	//   TileLayer "Foreground"
}

func Example_movedImageWinOp() {
	fmt.Println(Moved(px.V(55, -88), Image(nil, "IMap")).String())
	// Output:
	// Moved 55 pixels right 88 pixels down:
	//   Image "IMap"
}

func Example_colorImageWinOp() {
	fmt.Println(Color(colornames.Red, Image(nil, "IMap")).String())
	// Output:
	// Color {255 0 0 255}:
	//   Image "IMap"
}

func Example_sequencedWinOps() {
	mapImage := Color(colornames.Red, Image(nil, "IMap"))
	ghostImage := Color(colornames.Yellow, Image(nil, "IGhost"))
	sequence := OpSequence(mapImage, ghostImage)
	fmt.Println(sequence.String())
	// Output:
	// WinOp Sequence:
	//   Color {255 0 0 255}:
	//     Image "IMap"
	//   Color {255 255 0 255}:
	//     Image "IGhost"
}

func ExampleMirrored() {
	mapImage := Image(nil, "IMap")
	ghostImage := Color(colornames.Yellow, Image(nil, "IGhost"))
	seq := OpSequence(mapImage, ghostImage)
	mirrored := Mirrored(seq)
	fmt.Println(mirrored.String())
	// Output:
	// Mirrored around Y axis:
	//   WinOp Sequence:
	//     Image "IMap"
	//     Color {255 255 0 255}:
	//       Image "IGhost"
}
