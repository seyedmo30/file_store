package ports

import "context"
type StorageDB interface {

	AddFile(ctx context.Context , )( string ,error)
	
}