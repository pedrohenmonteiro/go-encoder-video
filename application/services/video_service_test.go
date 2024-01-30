package services_test

import (
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/pedrohenmonteiro/go-encoder-video/application/repositories"
	"github.com/pedrohenmonteiro/go-encoder-video/application/services"
	"github.com/pedrohenmonteiro/go-encoder-video/domain"
	"github.com/pedrohenmonteiro/go-encoder-video/framework/database"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func prepare() (*domain.Video, repositories.VideoRepositoryDB) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "convite.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDB{Db: db}
	repo.Insert(video)

	return video, repo
}

func TestVideoServiceDownload(t *testing.T) {
	video, repo := prepare()

	videoService := services.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download("testandoencode")
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)

	err = videoService.Encode()
	require.Nil(t, err)

	err = videoService.Finish()
	require.Nil(t, err)
}
