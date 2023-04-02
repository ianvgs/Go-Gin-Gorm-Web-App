package models

import (
	"time"
)

type Noticia struct {
	Id            int       `gorm:"column:id" binding:"required" gorm:"primaryKey"`
	ColaboradorId int       `gorm:"column:idColaborador" binding:"required" `
	Titulo        string    `gorm:"column:titulo" binding:"required"`
	Resumo        string    `gorm:"column:resumo" binding:"required"`
	Texto         string    `gorm:"column:texto" binding:"required"`
	IdCategoria   int       `gorm:"column:idCategoria" binding:"required"`
	Ativo         string    `gorm:"column:ativo" binding:"required"`
	CreatedAt     time.Time `gorm:"column:createdAt" binding:"required"`
	UpdatedAt     time.Time `gorm:"column:updatedAt" binding:"required"`
	Views         string    `gorm:"column:views" binding:"required"`
	Colaborador   Colaborador
}
