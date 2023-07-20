package smtp_security

type Security string

const (
	None     Security = "None"
	TLS               = "TLS"
	STARTTLS          = "STARTTLS"
)
