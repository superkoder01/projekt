package mailhog

import "time"

type Messages struct {
	Content Content
	Raw     Raw
	MIME    MIME
	Created time.Time
}

type Content struct {
	Headers Headers
	Body    string
}

type Headers struct {
	Subject []string
}

type Raw struct {
	From string
	To   []string
}

type MIME struct {
	Parts []Parts
}

type Parts struct {
	Body string
}
