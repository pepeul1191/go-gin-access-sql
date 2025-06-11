// controllers/foot_controller.go
package controllers

import (
	"access/app/configs"
	"access/app/models"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GET: apis/v1/systems
func SystemFetchAll(c *gin.Context) {
	var systems []models.System
	var total int64
	// Obtener parámetros opcionales
	name := c.Query("name")
	description := c.Query("description")
	step := c.Query("step")
	page := c.Query("page")
	// Conexión a la base de datos
	if err := configs.ConnectToDB(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "No se pudo conectar a la base de datos",
			"message": err.Error(),
		})
		return
	}
	// Comenzar la consulta
	query := configs.DB.Model(&models.System{})
	// Aplicar filtros opcionales
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if description != "" {
		query = query.Where("description LIKE ?", "%"+description+"%")
	}
	// Paginación (si step y page están presentes)
	if step != "" && page != "" {
		// Contar total (antes de paginar)
		if err := query.Count(&total).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error al contar systems",
				"message": err.Error(),
			})
			return
		}
		// traer pagina
		stepInt, err1 := strconv.Atoi(step)
		pageInt, err2 := strconv.Atoi(page)
		if err1 == nil && err2 == nil && stepInt > 0 && pageInt > 0 {
			offset := (pageInt - 1) * stepInt
			query = query.Offset(offset).Limit(stepInt)
			// Ejecutar la consulta
			if err := query.Find(&systems).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Error al listar la lista paginada de sistemas",
					"message": err.Error(),
				})
				return
			}
			// Respuesta
			pages := int(math.Ceil(float64(total) / float64(stepInt)))
			c.JSON(http.StatusOK, gin.H{
				"list":   systems,
				"pages":  pages,
				"total":  total,
				"offset": offset,
			})
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al leer los parametros de la paginación",
			"message": err1.Error(),
		})
		return
	} else {
		// Ejecutar la consulta
		if err := query.Find(&systems).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error al consultar systems",
				"message": err.Error(),
			})
			return
		}
		// Respuesta
		c.JSON(http.StatusOK, systems)
	}
}

// POST: /apis/v1/systems
func SystemCreate(c *gin.Context) {
	var system models.System
	// Parsear JSON recibido
	if err := c.ShouldBindJSON(&system); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "message": err.Error()})
		return
	}
	// Establecer fechas (si no se hace en frontend)
	system.Created = time.Now()
	system.Updated = time.Now()
	if err := configs.ConnectToDB(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Conexión fallida", "message": err.Error()})
		return
	}
	if err := configs.DB.Create(&system).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el sistema", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, system)
}

// PUT: /apis/v1/systems/:id
func SystemUpdate(c *gin.Context) {
	id := c.Param("id")
	var system models.System
	if err := configs.ConnectToDB(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Conexión fallida", "message": err.Error()})
		return
	}
	// Buscar el sistema existente
	if err := configs.DB.First(&system, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sistema no encontrado"})
		return
	}
	// Bind JSON sobre el sistema existente
	if err := c.ShouldBindJSON(&system); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "message": err.Error()})
		return
	}
	system.Updated = time.Now()
	if err := configs.DB.Model(&system).Select("name", "description", "repository", "updated").Updates(system).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, system)
}

// DELETE: /apis/v1/systems/:id
func SystemDelete(c *gin.Context) {
	id := c.Param("id")
	var system models.System
	// Conexión a la base de datos
	if err := configs.ConnectToDB(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Conexión fallida",
			"message": err.Error(),
		})
		return
	}
	// Buscar el sistema por ID
	if err := configs.DB.First(&system, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Sistema no encontrado",
		})
		return
	}
	// Eliminar el sistema
	if err := configs.DB.Delete(&system).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "No se pudo eliminar el sistema",
			"message": err.Error(),
		})
		return
	}
	// Éxito
	c.JSON(http.StatusOK, gin.H{
		"message": "Sistema eliminado correctamente",
		"id":      id,
	})
}

// GET: apis/v1/systems/:id/roles
func SystemFetchRoles(c *gin.Context) {
	var roles []models.Role
	idStr := c.Param("id")
	// Convertir el ID a uint (puedes usar strconv.ParseUint si prefieres)
	var systemID uint
	if _, err := fmt.Sscanf(idStr, "%d", &systemID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de sistema inválido",
		})
		return
	}
	// Conexión a la base de datos
	if err := configs.DB.Where("system_id = ?", systemID).Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al obtener los roles",
			"message": err.Error(),
		})
		return
	}
	// Respuesta
	// Si no se encontraron roles, devolver arreglo vacío
	if len(roles) == 0 {
		c.JSON(http.StatusOK, []struct{}{})
		return
	}
	// Devolver respuesta exitosa
	c.JSON(http.StatusOK, roles)
}
