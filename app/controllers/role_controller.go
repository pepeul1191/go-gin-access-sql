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
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID de sistema inválido", "error": err.Error()})
		return
	}

	// Leer JSON del cuerpo
	var req models.RoleCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No se pueden parsear los datos enviados", "error": err.Error()})
		return
	}

	// Conexión a la base de datos
	if err := configs.ConnectToDB(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "No se pudo conectar a la base de datos",
			"message": err.Error(),
		})
		return
	}

	// Iniciar transacción
	tx := configs.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	response := make([]models.CreatedRoleResponse, 0)

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
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al crear rol", "error": err.Error()})
			return
		}

		response = append(response, models.CreatedRoleResponse{
			Tmp: incoming.ID,
			ID:  strconv.Itoa(int(newRole.ID)),
		})
	}

	// 2. Actualizar roles existentes
	for _, incoming := range req.Edits {
		var role models.Role
		if err := tx.First(&role, uint(incoming.ID)).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, gin.H{"message": "Rol no encontrado", "id": incoming.ID})
			return
		}

		role.Name = incoming.Name
		role.Updated = time.Now()

		if err := tx.Model(&role).Select("Name", "Updated").Updates(role).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al actualizar rol", "id": incoming.ID, "error": err.Error()})
			return
		}
	}

	// 3. Eliminar roles
	for _, idToDelete := range req.Deletes {
		var role models.Role
		if err := tx.First(&role, uint(idToDelete)).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, gin.H{"message": "Rol no encontrado", "id": idToDelete})
			return
		}

		if err := tx.Delete(&role).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al eliminar rol", "error": err.Error()})
			return
		}
	}

	// Confirmar transacción
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al confirmar cambios", "error": err.Error()})
		return
	}

	// Responder con éxito
	c.JSON(http.StatusOK, response)
}
