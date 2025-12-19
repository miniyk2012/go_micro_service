package main

import (
	"context"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestServer(t *testing.T) {
	db, err := NewDB()
	if err != nil {
		t.Fatalf("new db failed: %v", err)
	}

	bs := &server{
		bs: &bookstore{db: db},
	}

	shelves, err := bs.ListShelves(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Fatalf("list shelves failed: %v", err)
	}

	t.Logf("shelves: %s", protojson.Format(shelves))
}
