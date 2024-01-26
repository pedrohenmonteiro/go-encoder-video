package repositories

import (
	"database/sql"
	"fmt"

	"github.com/pedrohenmonteiro/go-encoder-video/domain"
)

type JobRepository interface {
	Insert(job *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
}

type JobRepositoryDB struct {
	Db *sql.DB
}

func NewJobRepository(db *sql.DB) *JobRepositoryDB {
	return &JobRepositoryDB{
		Db: db,
	}
}

func (repo JobRepositoryDB) Insert(job *domain.Job) (*domain.Job, error) {

	stmt, err := repo.Db.Prepare(`
		INSERT INTO jobs (id, output_bucket_path, status, video_id, error, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(job.ID, job.OutputBucketPath, job.Status, job.Video.ID, job.Error, job.CreatedAt, job.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (repo JobRepositoryDB) Find(id string) (*domain.Job, error) {
	var job domain.Job

	row := repo.Db.QueryRow(`
		SELECT id, output_bucket_path, status, video_id, error, created_at, updated_at
		FROM jobs
		WHERE id = ?
	`, id)

	err := row.Scan(&job.ID, &job.OutputBucketPath, &job.Status, &job.VideoID, &job.Error, &job.CreatedAt, &job.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("job not found")
		}
		return nil, err
	}

	return &job, nil
}
