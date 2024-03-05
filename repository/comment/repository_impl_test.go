package comment

import (
	"context"
	"fmt"
	"testing"

	"github.com/banggibima/go-database/entity"
	"github.com/banggibima/go-database/sql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(sql.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "bimarepository@mail.com",
		Comment: "first repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(sql.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindById(ctx, 36)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(sql.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}
