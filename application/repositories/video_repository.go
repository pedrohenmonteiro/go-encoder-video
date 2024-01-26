package repositories

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/pedrohenmonteiro/go-encoder-video/domain"
)

type VideoRepository interface {
	Insert(vide *domain.Video) (*domain.Video, error)
	Find(id string) (*domain.Video, error)
}

type VideoRepositoryDB struct {
	Db *sql.DB
}

func NewVideoRepository(db *sql.DB) *VideoRepositoryDB {
	return &VideoRepositoryDB{
		Db: db,
	}
}

func (r VideoRepositoryDB) Insert(video *domain.Video) (*domain.Video, error) {
	if video.ID == "" {
		video.ID = uuid.New().String()
	}

	stmt, err := r.Db.Prepare("INSERT INTO videos (encoded_video_folder, resource_id, file_path, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(video.ID, video.ResourceID, video.FilePath, video.CreatedAt)
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (r VideoRepositoryDB) Find(id string) (*domain.Video, error) {

	row := r.Db.QueryRow("SELECT * FROM videos WHERE id = ?", id)

	var video domain.Video

	err := row.Scan(&video.ID, &video.ResourceID, &video.FilePath, &video.CreatedAt)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("video not found")
		}

		return nil, err
	}

	return &video, nil
}

// ID         string    `json:"encoded_video_folder" valid:"uuid"`
// ResourceID string    `json:"resource_id" :"notnull"`
// FilePath   string    `json:"file_path" valid:"notnull"`
// CreatedAt  time.Time `json:"-" valid:"-"`
// Jobs       []*Job    `json:"-" valid:"-"`
