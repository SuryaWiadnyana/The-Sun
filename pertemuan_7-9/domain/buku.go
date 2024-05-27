package domain

import (
	"context"
	"time"
)

type Buku struct {
	ID        string     `json:"id" bson:"id"`
	Judul     string     `json:"judul" bson:"judul"`
	Penulis   string     `json:"penulis" bson:"penulis"`
	Halaman   int        `json:"halaman" bson:"halaman"`
	Status    string     `json:"status" bson:"status"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
}

type RepositoryBuku interface {
	GetAll(ctx context.Context) ([]Buku, error)
	CreateBuku(ctx context.Context, buku *Buku) (Buku, error)
	GetBukuByID(ctx context.Context, id string) (Buku, error)
	DeleteBuku(ctx context.Context, id string) error
	UpdateBuku(ctx context.Context, db *Buku) error
}

type ServiceBuku interface {
	GetAll(ctx context.Context) ([]Buku, error)
	CreateBuku(ctx context.Context, buku *Buku) (Buku, error)
	GetBukuByID(ctx context.Context, id string) (Buku, error)
	DeleteBuku(ctx context.Context, id string) error
	UpdateBuku(ctx context.Context, body *Buku) error
}
