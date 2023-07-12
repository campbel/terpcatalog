package email

import "github.com/campbel/terpcatalog/shared/types"

type Sender interface {
	SendEmail(to []types.EmailAddress, subject, content string) error
}
