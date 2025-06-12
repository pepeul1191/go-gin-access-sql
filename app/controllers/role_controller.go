// controllers/foot_controller.go
package controllers

import (
	"access/app/configs"
	"access/app/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveRoles(c *gin.Context) {
	systemIdStr := c.Param("system-id")

	// Convertir el ID del sistema
	var systemID uint
	if _, err := fmt.Sscanf(systemIdStr, "%d", &systemID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de sistema inválido"})
		return
	}

	// Leer JSON del cuerpo
	var req models.RoleCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar conexión a la base de datos
	if configs.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Conexión a DB no disponible"})
		return
	}

	// Iniciar transacción
	tx := configs.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var response []map[string]string

	// 1. Crear nuevos roles
	for _, incoming := range req.News {
		newRole := models.Role{
			Name:     incoming.Name,
			Created:  time.Now(),
			Updated:  time.Now(),
			SystemID: systemID,
		}

		if err := tx.Create(&newRole).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear rol", "message": err.Error()})
			return
		}

		response = append(response, map[string]string{
			"tmp": incoming.ID,
			"id":  strconv.Itoa(int(newRole.ID)),
		})
	}

	// 2. Actualizar roles existentes
	for _, incoming := range req.Edits {
		var role models.Role
		if err := tx.First(&role, uint(incoming.ID)).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, gin.H{"error": "Rol no encontrado", "id": incoming.ID})
			return
		}

		role.Name = incoming.Name
		role.Updated = time.Now()

		if err := tx.Model(&role).Select("Name", "Updated").Updates(role).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar rol", "id": incoming.ID, "message": err.Error()})
			return
		}
	}

	// 3. Eliminar roles
	for _, idToDelete := range req.Deletes {
		var role models.Role
		if err := tx.First(&role, uint(idToDelete)).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, gin.H{"error": "Rol no encontrado", "id": idToDelete})
			return
		}

		if err := tx.Delete(&role).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar rol", "message": err.Error()})
			return
		}
	}

	// Confirmar transacción
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al confirmar cambios", "message": err.Error()})
		return
	}

	// Responder con éxito
	c.JSON(http.StatusOK, gin.H{
		"created": response,
	})
}
