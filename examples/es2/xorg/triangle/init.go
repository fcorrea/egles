// +build !raspberry

package main

import (
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/mortdeus/egles/egl"
	"github.com/mortdeus/egles/egl/platform/xorg"
	"log"
)

const (
	INITIAL_WINDOW_WIDTH  = 640
	INITIAL_WINDOW_HEIGHT = 480
)

var X *xgbutil.XUtil

func initialize() {
	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	mousebind.Initialize(X)
	xWindow := newWindow(X, INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT)
	go xevent.Main(X)
	xorg.Initialize(
		egl.NativeWindowType(uintptr(xWindow.Id)),
		xorg.DefaultConfigAttributes,
		xorg.DefaultContextAttributes)
}
