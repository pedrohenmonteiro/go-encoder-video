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

func TestJobRepositoryInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDB{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	assert.Nil(t, err)

	repoJob := repositories.JobRepositoryDB{Db: db}
	repoJob.Insert(job)

	j, err := repoJob.Find(job.ID)
	assert.NotEmpty(t, j.ID)
	assert.Nil(t, err)
	assert.Equal(t, j.ID, job.ID)
	assert.Equal(t, j.VideoID, video.ID)

}
