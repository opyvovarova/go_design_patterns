package email

type Email struct {
	to string
	from string
	subject string
	body string
}

func (e *Email) To() string {
	return e.to
}

func (e *Email) From() string {
	return e.from
}

func (e *Email) Subject() string {
	return e.subject
}

func (e *Email) Body() string {
	return e.body
}

type Builder struct {
	to string
	from string
	subject string
	body string
}

func (b *Builder) To(to string) *Builder {
	b.to = to
	return b
}

func (b *Builder) From(from string) *Builder {
	b.from = from
	return b
}

func (b *Builder) Subject(subject string) *Builder {
	b.subject = subject
	return b
}

func (b *Builder) Body(body string) *Builder {
	b.body = body
	return b
}

func (b *Builder) Build() *Email {
	e := new(Email)
	e.to = b.to
	e.from = b.from
	e.subject = b.subject
	e.body = b.body
	return e
}

func EmailBuilder() *Builder {
	return new(Builder)
}


