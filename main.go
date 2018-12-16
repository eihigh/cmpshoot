package main

import (
	"image/color"
	"math"
	"math/cmplx"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func main() {
	if err := ebiten.Run(update, 640, 480, 1, "shoot by complexes"); err != nil {
		panic(err)
	}
}

func update(screen *ebiten.Image) error {

	// 発射元の位置
	pos := complex(320, 240)
	// 一発目のベクトル
	vec := complex(100, 0)
	// 9方向に発射
	n := 9
	// 9分の1回転するベクトル
	rot := cmplx.Rect(1, 2*math.Pi/float64(n))
	// 実際に発射してみる
	for i := 0; i < n; i++ {
		bullet := pos + vec
		x, y := real(bullet), imag(bullet)
		ox, oy := real(pos), imag(pos)
		ebitenutil.DrawLine(screen, ox, oy, x, y, color.White)
		vec *= rot
	}

	return nil
}
