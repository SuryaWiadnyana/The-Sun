package services

import (
	"context"
	"pertemuan_7/domain"
	"time"
)

type ServiceBuku struct {
	RepositoryBuku domain.RepositoryBuku
	contextTimeout time.Duration
}

func NewServiceBuku(Rb domain.RepositoryBuku, T time.Duration) domain.ServiceBuku {
	return &ServiceBuku{
		RepositoryBuku: Rb,
		contextTimeout: T,
	}
}

func (b *ServiceBuku) GetAll(ctx context.Context) ([]domain.Buku, error) {
	ctx, cancel := context.WithTimeout(context.Background(), b.contextTimeout)
	defer cancel()

	return b.RepositoryBuku.GetAll(ctx)
}

func (b *ServiceBuku) CreateBuku(ctx context.Context, buku *domain.Buku) (domain.Buku, error) {
	ctx, cancel := context.WithTimeout(context.Background(), b.contextTimeout)
	defer cancel()

	return b.RepositoryBuku.CreateBuku(ctx, buku)
}

func (b *ServiceBuku) GetBukuByID(ctx context.Context, id string) (domain.Buku, error) {
	ctx, cancel := context.WithTimeout(context.Background(), b.contextTimeout)
	defer cancel()

	return b.RepositoryBuku.GetBukuByID(ctx, id)
}

func (b *ServiceBuku) UpdateBuku(ctx context.Context, body *domain.Buku) error {
	ctx, cancel := context.WithTimeout(context.Background(), b.contextTimeout)
	defer cancel()

	return b.RepositoryBuku.UpdateBuku(ctx, body)
}

func (b *ServiceBuku) DeleteBuku(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), b.contextTimeout)
	defer cancel()

	return b.RepositoryBuku.DeleteBuku(ctx, id)
}
