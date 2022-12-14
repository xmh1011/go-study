package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// 图像序列是确定的，除非我们使用当前时间播种伪随机数生成器
	rand.Seed(time.Now().UTC().UnixNano()) // 生成随机数种子
	if len(os.Args) > 1 && os.Args[1] == "web" {
		// web()
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		// HandleFunc registers the handler function for the given pattern
		// in the DefaultServeMux.
		// The documentation for ServeMux explains how patterns are matched.
		http.HandleFunc("/", handler) // each request calls handler
		// 回声请求调用处理程序
		http.HandleFunc("/count", counter)
		// Fatal is equivalent to Print() followed by a call to os.Exit(1).
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

//利萨茹图形实现函数
func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
