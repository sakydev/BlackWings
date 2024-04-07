package types

import (
	"github.com/guregu/null/v5"
	"time"
)

type EmailResponse struct {
	Subject     string
	SenderName  null.String
	SenderEmail string
	Date        time.Time
	Snippet     string
}
