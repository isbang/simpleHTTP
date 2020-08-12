package simple

type HandlerFunc func(c Context)

type Handler interface {
	Handle(c Context)
}

type Getter interface {
	Get(c Context)
}

type Poster interface {
	Post(c Context)
}

type Putter interface {
	Put(c Context)
}

type Deleter interface {
	Delete(c Context)
}

type BasicHandler interface {
	Getter
	Poster
}

type RestHandler interface {
	Getter
	Poster
	Putter
	Deleter
}
