package minhaodb

import (
	"database/sql"
	"fmt"
	"log"
)

type MinhaoDB struct {
	*sql.DB
}

// NewMinhaoDB 初始化数据库连接并返回 MinhaoDB 的指针，封装了sql指针
func NewMinhaoDB(dsn string) (*MinhaoDB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping不通数据库: %w", err)
	}

	return &MinhaoDB{DB: db}, nil
}

// 封装sql的query方法
func (minhaodb *MinhaoDB) QueryWrapper(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := minhaodb.Query(query, args...)
	if err != nil {
		log.Printf("参数错误: %v\n", err)
		return nil, err
	}
	return rows, nil
}
