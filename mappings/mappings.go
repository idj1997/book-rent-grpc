package mappings

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/idj1997/book-rent-core/domain"
	"github.com/idj1997/book-rent-core/service"
	"github.com/idj1997/book-rent-grpc/proto/proto_book"
	"github.com/idj1997/book-rent-grpc/proto/proto_common"
	"time"
)

func BookToProtoBook(book *domain.Book) *proto_book.Book {
	return &proto_book.Book{
		Metadata: &proto_common.Metadata{
			CreatedAt: TimeToTimestamp(book.CreatedAt),
			UpdatedAt: TimeToTimestamp(book.UpdatedAt),
		},
		Id:       int32(book.ID),
		Title:    book.Title,
		Content:  book.Content,
		Stock:    int32(book.Stock),
	}
}

func ErrorToServiceError(err error) *service.ServiceError {
	return err.(*service.ServiceError)
}


func TimeToTimestamp(time time.Time) *timestamp.Timestamp {
	timestamp, _ := ptypes.TimestampProto(time)
	return timestamp
}