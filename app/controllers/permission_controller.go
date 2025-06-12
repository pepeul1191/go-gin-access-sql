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

// POST: apis/v1/permissions/:role-id
func SavePermissions(c *gin.Context) {
	roleIdStr := c.Param("role-id")

	// Convertir el ID del sistema
	var roleID uint
	if _, err := fmt.Sscanf(roleIdStr, "%d", &roleID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID de rol inválido", "error": err.Error()})
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

	response := make([]models.CreatedPermissionResponse, 0)

	// 1. Crear nuevos permissions
	for _, incoming := range req.News {
		newPermission := models.Permission{
			Name:    incoming.Name,
			Created: time.Now(),
			Updated: time.Now(),
			RoleId:  roleID,
		}

		if err := tx.Create(&newPermission).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al crear rol", "error": err.Error()})
			return
		}

		response = append(response, models.CreatedPermissionResponse{
			Tmp: incoming.ID,
			ID:  strconv.Itoa(int(newPermission.ID)),
		})
	}

	// 2. Actualizar permissions existentes
	for _, incoming := range req.Edits {
		var permission models.Permission
		if err := tx.First(&permission, uint(incoming.ID)).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, gin.H{"message": "Rol no encontrado", "id": incoming.ID})
			return
		}

		permission.Name = incoming.Name
		permission.Updated = time.Now()

		if err := tx.Model(&permission).Select("Name", "Updated").Updates(permission).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al actualizar permiso", "id": incoming.ID, "error": err.Error()})
			return
		}
	}

	// 3. Eliminar permissions
	for _, idToDelete := range req.Deletes {
		var permission models.Permission
		if err := tx.First(&permission, uint(idToDelete)).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, gin.H{"message": "Permiso no encontrado", "id": idToDelete})
			return
		}

		if err := tx.Delete(&permission).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al eliminar permiso", "error": err.Error()})
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
