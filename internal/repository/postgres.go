package repository

import (
	"context"
	"os"
	"test/pkg/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type SqlHandler struct {
	DB  *pgxpool.Pool
	Dsn string
}

// Создание соединения для подключения к базе данных
func New() *SqlHandler {
	// Здесь подгружаются данные для подключения к базе данных из файла .env (os.Getenv("DB_USER") и т.д.)
	// Обычно загрузка переменных окружения происходит в другом месте, но для читаемости кода этот файл загружается здесь
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dsn := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME")
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}
	config.MaxConns = 10
	config.MinConns = 5
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}
	err = pool.Ping(context.Background())
	if err != nil {
		panic("SqlHandler error ping database:" + err.Error())
	}
	return &SqlHandler{DB: pool, Dsn: dsn}
}

func (s *SqlHandler) Close() {
	s.DB.Close()
}

func (s *SqlHandler) InsertProgram(ctx context.Context, program models.Program) error {
	query := `INSERT INTO programs (id, name, name_en, is_public, project_id) 
              VALUES ($1, $2, $3, $4, $5)
              ON CONFLICT (id) DO NOTHING`

	_, err := s.DB.Exec(ctx, query, program.ID, program.Name, program.NameEn, program.IsPublic, program.ProjectID)
	return err
}

func (s *SqlHandler) GetProgram(id uuid.UUID) (*models.Program, error) {
	query := `SELECT id, name, name_en, is_public, project_id FROM programs WHERE id = $1`

	var program models.Program
	err := s.DB.QueryRow(context.Background(), query, id).Scan(&program.ID, &program.Name, &program.NameEn, &program.IsPublic, &program.ProjectID)
	if err != nil {
		return nil, err
	}

	return &program, nil
}

func (s *SqlHandler) UpdateProgram(ctx context.Context, program models.Program) error {
	query := `UPDATE programs SET name = $1, name_en = $2, is_public = $3, project_id = $4 WHERE id = $5`

	_, err := s.DB.Exec(ctx, query, program.Name, program.NameEn, program.IsPublic, program.ProjectID, program.ID)
	return err
}

func (s *SqlHandler) DeleteProgram(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM programs WHERE id = $1`

	_, err := s.DB.Exec(ctx, query, id)
	return err
}
