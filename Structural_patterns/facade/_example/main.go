package main

import "github.com/jiandahao/DesignPatterns/Structural_patterns/facade"

func main() {
	// 应用程序的类并不依赖于复杂框架中成千上万的类。同样，如果你决定更换框架，
	// 那只需重写外观类即可。
	converter := &facade.VideoConverter{}
	avi := converter.Convert("demo.mp4", "avi")
	avi.Save("./out.avi")

	mp4 := converter.Convert("demo.avi", "mp4")
	mp4.Save("./out.mp4")
}
