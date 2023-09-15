# geometry

Go Package to perform operations on 2D and 3D geometric figures

## Installation

Run this command in your project directory:

```
go get github.com/nemzyxt/geometry
```

## Example

```
package main

import (
    "fmt"

    "github.com/nemzyxt/geometry/2d" // planar : 2d shapes
    "github.com/nemzyxt/geometry/3d" // spatial: 3d shapes
)

func main() {
    r := planar.NewRectangle(5, 4)
    c := spatial.NewCube(4)

    fmt.Println(r.Area())      // 20
    fmt.Println(r.Perimeter()) // 18
    fmt.Println(c.Volume())    // 64
}
```

## 2D Shapes (area, perimeter)

- Circle
- Square
- Rectangle
- Triangle

## 3D Shapes (volume)

- Cube
- Cuboid
- Cylinder
