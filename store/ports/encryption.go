package ports

import "context"

type Encryption interface{

	Encrypt(ctx context.Context , file  *[]byte )( *[]byte ,error)
}