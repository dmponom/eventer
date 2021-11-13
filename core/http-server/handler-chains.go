package httpserver

import "net/http"

type Handler func(w http.ResponseWriter, r *http.Request)

type HandlesChain struct {
	head *HandleChainLeaf
	tail *HandleChainLeaf
}

func (hc *HandlesChain) append(handle Handler) {
	leaf := HandleChainLeaf{handler: handle}

	if hc.head == nil && hc.tail == nil {
		hc.head = &leaf
		hc.tail = &leaf
	} else {
		hc.tail.next = &leaf
		hc.tail = &leaf
	}
}

type HandleChainLeaf struct {
	handler Handler
	next    *HandleChainLeaf
}

func (hcl *HandleChainLeaf) Handle(w http.ResponseWriter, r *http.Request) {
	hcl.handler(w, r)
	if hcl.next != nil {
		hcl.next.Handle(w, r)
	}
}

func MakeHandlesChain(handlers ...Handler) http.HandlerFunc {
	chain := HandlesChain{}

	for _, h := range handlers {
		chain.append(h)
	}

	return chain.head.Handle
}
