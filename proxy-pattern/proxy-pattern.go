package proxy_pattern

import "fmt"

type Image interface {
	Display()
}

type RealImage struct {
}

func (r *RealImage) Display() {
	fmt.Println("Display")
}

func (r *RealImage) Load() {
	fmt.Println("Load")
}

type ProxyImage struct {
	realImage *RealImage
}

func (p *ProxyImage) Display() {
	if nil == p.realImage {
		p.realImage = &RealImage{}
		p.realImage.Load()
	}
	p.realImage.Display()
}
