package event_dispatcher

import "sync"

type listener[E any] func(E)

type node[E any] struct {
	fn   listener[E]
	next *node[E]
}

type list[E any] struct {
	first *node[E]
	last  *node[E]
	m     sync.RWMutex
}

func (l *list[E]) add(fn listener[E]) {
	node := node[E]{
		fn:   fn,
		next: nil,
	}
	l.m.Lock()
	defer l.m.Unlock()
	if l.first == nil {
		l.first = &node
		l.last = &node
		return
	}
	l.last.next = &node
}

func (l *list[E]) forEach(fn func(ls listener[E])) {
	node := l.first
	if l.first == nil {
		return
	}
	l.m.RLock()
	defer l.m.RLock()
	for node != nil {
		fn(node.fn)
		node = node.next
	}
}

type EventDispatcher[T comparable, E any] struct {
	ch        map[T]chan E
	listeners map[T]*list[E]

	buffer   int
	dmu, lmu sync.Mutex
}

func NewEventDispatcher[T comparable, E any](buffer int) *EventDispatcher[T, E] {
	return &EventDispatcher[T, E]{
		ch:        map[T]chan E{},
		listeners: map[T]*list[E]{},
		buffer:    buffer,
	}
}

func (e *EventDispatcher[T, E]) DispatchEvent(eventType T, event E) {
	ch, ok := e.ch[eventType]
	if !ok {
		ch = make(chan E, e.buffer)
		e.dmu.Lock()
		e.ch[eventType] = ch
		e.dmu.Unlock()
		go e.startDispatch(eventType)
	}
	ch <- event
}

func (e *EventDispatcher[T, E]) AddEventListener(eventType T, fn listener[E]) {
	l, ok := e.listeners[eventType]
	if !ok {
		l = &list[E]{}
		e.lmu.Lock()
		e.listeners[eventType] = l
		e.lmu.Unlock()
	}
	l.add(fn)
}

/*func (e *EventDispatcher[E, T]) Close() {
	for _, es := range e.ch {
		close(es)
	}
}*/

func (e *EventDispatcher[T, E]) startDispatch(eventType T) {
	ch := e.ch[eventType]
	for event := range ch {
		ls, ok := e.listeners[eventType]
		if !ok {
			continue
		}

		ls.forEach(func(fn listener[E]) {
			go fn(event)
		})
	}
}
