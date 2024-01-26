package repositories_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pedrohenmonteiro/go-encoder-video/application/repositories"
	"github.com/pedrohenmonteiro/go-encoder-video/domain"
	"github.com/pedrohenmonteiro/go-encoder-video/framework/database"
	"github.com/stretchr/testify/assert"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDB{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	assert.Nil(t, err)
	assert.NotEmpty(t, v.ID)
	assert.Equal(t, v.ID, video.ID)
}
