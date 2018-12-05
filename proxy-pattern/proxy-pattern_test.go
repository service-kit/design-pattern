package proxy_pattern

import "testing"

func Test_ProxyPattern(t *testing.T) {
	p := &ProxyImage{}
	p.Display()
	p.Display()
}
