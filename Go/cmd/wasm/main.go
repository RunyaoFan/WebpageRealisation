package main

import (
	"syscall/js"
  "strings"
)

var slideIndex = 1;

// go function used to facilitate carousel functionalities
func showSlides(n int) {
	var i int
	var slides = js.Global().Get("document").Call("getElementsByClassName", "mySlides")
	var dots = js.Global().Get("document").Call("getElementsByClassName", "carousel-indicator")

	if n > slides.Length() {slideIndex = 1}
	if n < 1 {slideIndex = slides.Length()}
	for i = 0; i < slides.Length(); i++ {
			slides.Index(i).Set("style", "display: none;")
	}

	for i = 0; i < dots.Length(); i++ {
			classes := dots.Index(i).Get("className").String()
	    dots.Index(i).Set("className", strings.Replace(classes, " active", "", 1))
	}

	slides.Index(slideIndex - 1).Set("style", "display: block;")
	classes := dots.Index(slideIndex - 1).Get("className")
	dots.Index(slideIndex - 1).Set("className", classes.String() + " active")
}

func main() {
  // these variables represent functions callable by JavaScript, they are
  // written in golang and added with a wrapper
	var cb, prevcb, nextcb, show1, show2, show3, enableAuto, disableAuto js.Func

	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		collapsibleContent := js.Global().Get("document").Call("getElementById", "collapsible")

		var elementStatus string = collapsibleContent.Get("style").Get("display").String()

		if elementStatus == "none" {
			collapsibleContent.Set("style", "display: block;")
		} else {
			collapsibleContent.Set("style", "display: none;")
		}

		return nil
	})

	nextcb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    slideIndex = slideIndex + 1
    showSlides(slideIndex)

		return nil
	})

  prevcb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    slideIndex = slideIndex - 1
    showSlides(slideIndex)

		return nil
	})

  show1 = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    slideIndex = 1
    showSlides(slideIndex)

		return nil
	})

  show2 = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    slideIndex = 2
    showSlides(slideIndex)

		return nil
	})

  show3 = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    slideIndex = 3
    showSlides(slideIndex)

		return nil
	})

  var myVar = js.Global().Call("setInterval", nextcb, 5000)

  enableAuto = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    myVar = js.Global().Call("setInterval", nextcb, 5000)

		return nil
	})

  disableAuto = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
    js.Global().Call("clearInterval", myVar)

		return nil
	})

	js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", cb)
	js.Global().Get("document").Call("getElementById", "nextSlide").Call("addEventListener", "click", nextcb)
  js.Global().Get("document").Call("getElementById", "prevSlide").Call("addEventListener", "click", prevcb)
  js.Global().Get("document").Call("getElementById", "slide1").Call("addEventListener", "click", show1)
  js.Global().Get("document").Call("getElementById", "slide2").Call("addEventListener", "click", show2)
  js.Global().Get("document").Call("getElementById", "slide3").Call("addEventListener", "click", show3)
  js.Global().Get("document").Call("getElementById", "carousel-area").Call("addEventListener", "mouseleave", enableAuto)
  js.Global().Get("document").Call("getElementById", "carousel-area").Call("addEventListener", "mouseenter", disableAuto)
	showSlides(1)



	<-make(chan bool)
}
