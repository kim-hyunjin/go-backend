package decorator_pattern

import (
	"log"
	"net/http"
	"time"
)

type DecoratorFunc func (http.ResponseWriter, *http.Request, http.Handler)

type DecoHandler struct {
	fn DecoratorFunc
	h http.Handler
}

func (self *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	self.fn(w, r, self.h)
}

func NewDecoHandler(h http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler{
		fn: fn,
		h: h,
	}
}

func logger(rw http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Print("[LOGGER1] Started")
	h.ServeHTTP(rw, r)
	log.Print("[LOGGER1] Completed time: ", time.Since(start).Milliseconds())
}