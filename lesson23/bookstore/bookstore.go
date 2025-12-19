package main

import (
	"context"
	"errors"
	pb "lesson23/bookstore/pb"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

const (
	defaultShelfSize = 100
	defaultCursor    = 0 // 默认游标
	defaultPageSize  = 2 // 默认每页显示数量
)

type server struct {
	pb.UnimplementedBookstoreServer
	bs *bookstore
}

// CreateShelf 创建一个新的书架
func (s *server) CreateShelf(ctx context.Context, req *pb.CreateShelfRequest) (*pb.Shelf, error) {
	// 实现创建书架的逻辑
	if len(req.GetShelf().GetTheme()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid theme")
	}
	var size int64 = defaultShelfSize
	if req.GetShelf().GetSize() > 0 {
		size = req.GetShelf().GetSize()
	}
	data := &Shelf{
		Theme: req.GetShelf().GetTheme(),
		Size:  size,
	}
	shelf, err := s.bs.CreateShelf(ctx, data)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Shelf{
		Id:    shelf.ID,
		Theme: shelf.Theme,
		Size:  shelf.Size,
	}, nil
}

// GetShelf 获取一个指定的书架
func (s *server) GetShelf(ctx context.Context, req *pb.GetShelfRequest) (*pb.Shelf, error) {
	// 实现获取书架的逻辑
	if req.GetShelfId() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}
	shelf, err := s.bs.GetShelf(ctx, req.ShelfId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Shelf{
		Id:    shelf.ID,
		Theme: shelf.Theme,
		Size:  shelf.Size,
	}, nil
}

// ListShelves 列出所有书架
func (s *server) ListShelves(ctx context.Context, _ *emptypb.Empty) (*pb.ListShelvesResponse, error) {
	// 实现列出所有书架的逻辑
	shelves, err := s.bs.ListShelves(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) { // 没有数据
		return &pb.ListShelvesResponse{}, nil
	}
	if err != nil { // 查询数据库失败
		return nil, status.Error(codes.Internal, "query failed")
	}
	ns := make([]*pb.Shelf, len(shelves), len(shelves))
	for i, shelf := range shelves {
		ns[i] = &pb.Shelf{
			Id:    shelf.ID,
			Theme: shelf.Theme,
			Size:  shelf.Size,
		}
	}
	return &pb.ListShelvesResponse{
		Shelves: ns,
	}, nil
}

// DeleteShelf 删除一个指定的书架
func (s *server) DeleteShelf(ctx context.Context, req *pb.DeleteShelfRequest) (*emptypb.Empty, error) {
	// 实现删除书架的逻辑
	if req.GetShelfId() <= 0 {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, "invalid shelf id")
	}
	err := s.bs.DeleteShelf(ctx, req.ShelfId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *server) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	if req.GetShelfId() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}
	var (
		hasNext             = false
		cursor        int64 = defaultCursor
		pageSize            = defaultPageSize
		realSize      int
		nextPageToken Token
	)
	if req.PageToken != "" {
		token := Token(req.PageToken)
		pageInfo := token.Decode()
		if pageInfo.InValid() {
			return nil, status.Error(codes.InvalidArgument, "invalid page_token")
		}
		cursor = pageInfo.NextID
		pageSize = int(pageInfo.PageSize)
	}

	books, err := s.bs.GetBookListByShelfID(ctx, req.ShelfId, cursor, pageSize+1)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(books) > pageSize {
		hasNext = true
		nextPageToken = Page{
			NextID:        books[pageSize-1].ID,
			NextTimeAtUTC: time.Now().Unix(),
			PageSize:      int64(pageSize),
		}.Encode()
	}
	if hasNext {
		realSize = pageSize
	} else {
		realSize = len(books)
	}
	ns := make([]*pb.Book, realSize, realSize)
	for i := 0; i < realSize; i++ {
		book := books[i]
		ns[i] = &pb.Book{
			Id:     book.ID,
			Author: book.Author,
			Title:  book.Title,
		}
	}
	return &pb.ListBooksResponse{Books: ns, NextPageToken: string(nextPageToken)}, nil
}
