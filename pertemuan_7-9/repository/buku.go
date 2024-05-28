package repository

import (
	"context"
	"fmt"
	"log"
	"pertemuan_7/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var _Buku = "buku"

type mongoRepoBuku struct {
	DB *mongo.Database
}

func NewMongoRepoBuku(client *mongo.Database) domain.RepositoryBuku {
	return &mongoRepoBuku{
		DB: client,
	}
}

func (b *mongoRepoBuku) GetAll(ctx context.Context) ([]domain.Buku, error) {
	BookList := b.DB.Collection(_Buku)

	var Dbuku []domain.Buku

	data, err := BookList.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer data.Close(ctx)

	for data.Next(ctx) {
		var buku domain.Buku
		if err := data.Decode(&buku); err != nil {
			log.Println("Error decoding book:", err)
			continue
		}
		Dbuku = append(Dbuku, buku)
	}

	if err := data.Err(); err != nil {
		return nil, err
	}

	return Dbuku, err
}

func (b *mongoRepoBuku) CreateBuku(ctx context.Context, buku *domain.Buku) (domain.Buku, error) {
	BookList := b.DB.Collection(_Buku)

	data, err := BookList.InsertOne(ctx, buku)
	if err != nil {
		return domain.Buku{}, err
	}

	InsertedID, ok := data.InsertedID.(domain.Buku)
	if !ok {
		return domain.Buku{}, fmt.Errorf("inserted ID is not correct")
	}

	return InsertedID, nil
}

func (b *mongoRepoBuku) GetBukuByID(ctx context.Context, id string) (domain.Buku, error) {
	BookList := b.DB.Collection(_Buku)

	var buku domain.Buku
	err := BookList.FindOne(ctx, bson.M{"id": id}).Decode(&buku)
	if err != nil {
		return domain.Buku{}, err
	}

	return buku, nil
}

func (b *mongoRepoBuku) UpdateBuku(ctx context.Context, body *domain.Buku) error {
	BookList := b.DB.Collection(_Buku)
	body.UpdatedAt = time.Now()

	ID := bson.M{"_id": body.ID, "updated_at": nil}
	Upd := bson.M{
		"judul":   body.Judul,
		"penulis": body.Penulis,
		"status":  body.Status,
		"halaman": body.Halaman,
	}
	update := bson.M{"$set": Upd}

	_, err := BookList.UpdateOne(ctx, update, ID)
	if err != nil {
		return err
	}

	return nil
}

func (b *mongoRepoBuku) DeleteBuku(ctx context.Context, id string) error {
	BookList := b.DB.Collection(_Buku)

	_, err := BookList.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}

	return nil
}
