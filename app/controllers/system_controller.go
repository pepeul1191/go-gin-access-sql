// controllers/foot_controller.go
package controllers

import (
	"access/app/configs"
	"access/app/models"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GET: apis/v1/systems
func SystemFetchAll(c *gin.Context) {
	var systems []models.System
	var total int64

	name := c.Query("name")
	description := c.Query("description")
	step := c.Query("step")
	page := c.Query("page")

	if configs.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Conexión DB no inicializada"})
		return
	}

	query := configs.DB.Model(&models.System{})

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if description != "" {
		query = query.Where("description LIKE ?", "%"+description+"%")
	}

	if step != "" && page != "" {
		stepInt, err1 := strconv.Atoi(step)
		pageInt, err2 := strconv.Atoi(page)

		if err1 != nil || err2 != nil || stepInt <= 0 || pageInt <= 0 {
			var msg string
			if err1 != nil {
				msg = "step inválido: " + err1.Error()
			} else if err2 != nil {
				msg = "page inválido: " + err2.Error()
			} else {
				msg = "Los valores de paginación deben ser mayores a cero"
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Error al leer los parámetros de la paginación",
				"message": msg,
			})
			return
		}

		if err := query.Count(&total).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error al contar systems",
				"message": err.Error(),
			})
			return
		}

		offset := (pageInt - 1) * stepInt
		query = query.Offset(offset).Limit(stepInt)

		if err := query.Find(&systems).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error al listar la lista paginada de sistemas",
				"message": err.Error(),
			})
			return
		}

		pages := int(math.Ceil(float64(total) / float64(stepInt)))
		c.JSON(http.StatusOK, gin.H{
			"list":   systems,
			"pages":  pages,
			"total":  total,
			"offset": offset,
		})
		return
	}

	if err := query.Find(&systems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al consultar systems",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, systems)
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
	if configs.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Conexión DB no inicializada"})
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
	if configs.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Conexión DB no inicializada"})
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

// GET: /apis/v1/systems/:id/users
func SystemFetchUsers(c *gin.Context) {
	var users []models.UserWithRegistrationStatus
	var total int64

	idStr := c.Param("id")
	var systemID uint
	if _, err := fmt.Sscanf(idStr, "%d", &systemID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de sistema inválido"})
		return
	}

	username := c.Query("username")
	email := c.Query("email")
	step := c.Query("step")
	page := c.Query("page")
	activated := c.Query("status")

	// Conexión a DB (asegúrate de que esto devuelva un *gorm.DB válido)
	if err := configs.ConnectToDB(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Conexión fallida",
			"message": err.Error(),
		})
		return
	}

	// Construimos la parte base de la consulta
	baseQuery := `
			SELECT 
					U.id,
					U.username,
					U.email,
					CASE WHEN SU.system_id IS NOT NULL THEN TRUE ELSE FALSE END AS registered
			FROM users U
			LEFT JOIN systems_users SU ON U.id = SU.user_id AND SU.system_id = ?
			WHERE 1=1
	`

	// Aplicamos filtros opcionales
	args := []interface{}{systemID}
	countArgs := []interface{}{systemID}

	if username != "" {
		baseQuery += " AND U.username LIKE ?"
		args = append(args, "%"+username+"%")
		countArgs = append(countArgs, "%"+username+"%")
	}

	if activated == "1" || activated == "0" {
		var registeredFilter bool
		if activated == "1" {
			registeredFilter = true // usuario registrado
		} else {
			registeredFilter = false // usuario no registrado
		}
		baseQuery += " AND registered = ?"
		args = append(args, registeredFilter)
		countArgs = append(countArgs, registeredFilter)
	}

	if email != "" {
		baseQuery += " AND U.email LIKE ?"
		args = append(args, "%"+email+"%")
		countArgs = append(countArgs, "%"+email+"%")
	}

	// Contar total
	countQuery := "SELECT COUNT(*) FROM (" + baseQuery + ") AS tmp"

	if err := configs.DB.Raw(countQuery, countArgs...).Scan(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al contar usuarios",
			"message": err.Error(),
		})
		return
	}

	// Paginación
	if step != "" && page != "" {
		stepInt, err1 := strconv.Atoi(step)
		pageInt, err2 := strconv.Atoi(page)

		if err1 == nil && err2 != nil || stepInt <= 0 || pageInt <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros de paginación inválidos"})
			return
		}

		offset := (pageInt - 1) * stepInt
		args = append(args, stepInt, offset)
		baseQuery += " LIMIT ? OFFSET ?"

	} else if step != "" || page != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ambos parámetros 'step' y 'page' deben estar presentes"})
		return
	}

	// Ejecutar la consulta final
	if err := configs.DB.Raw(baseQuery, args...).Scan(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al obtener usuarios del sistema",
			"message": err.Error(),
		})
		return
	}

	// Respuesta
	if step != "" && page != "" {
		stepInt, _ := strconv.Atoi(step)
		pages := int(math.Ceil(float64(total) / float64(stepInt)))
		offset := (stepInt * (stepInt - 1))

		c.JSON(http.StatusOK, gin.H{
			"list":   users,
			"pages":  pages,
			"total":  total,
			"offset": offset,
		})
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// POST: /apis/v1/systems/:id/users
func SystemSaveUsers(c *gin.Context) {
	systemIdStr := c.Param("id")

	// Convertir el ID del sistema
	var systemID uint
	if _, err := fmt.Sscanf(systemIdStr, "%d", &systemID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID de sistemas no inválido", "error": err.Error()})
		return
	}

	// Leer JSON del cuerpo
	var req models.SystemUsersCreateRequest
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

	// 1. Crear nuevos permissions???

	// 2. Actualizar permissions existentes
	for _, incoming := range req.Edits {
		// Buscar si YA EXISTE la relación user <-> system
		var existing models.SystemUser
		err := tx.Where("user_id = ? AND system_id = ?", incoming.ID, systemID).
			First(&existing).Error

		if incoming.Registered {
			// Queremos que el usuario esté registrado
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					// No existe → crear
					newRelation := models.SystemUser{
						UserID:   incoming.ID,
						SystemID: systemID,
						Created:  time.Now(),
					}
					if err := tx.Create(&newRelation).Error; err != nil {
						tx.Rollback()
						c.JSON(http.StatusInternalServerError, gin.H{
							"message": "Error al crear la relación",
							"id":      incoming.ID,
						})
						return
					}
				} else {
					tx.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": "Error al buscar relación existente",
						"id":      incoming.ID,
					})
					return
				}
			}
		} else {
			// Queremos que el usuario NO esté registrado
			if err == nil {
				// Existe → eliminar
				if err := tx.Delete(&existing).Error; err != nil {
					tx.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": "Error al eliminar la relación",
						"id":      incoming.ID,
					})
					return
				}
			} else if errors.Is(err, gorm.ErrRecordNotFound) {
				// No existe → ignorar
			} else {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error al buscar relación para eliminar",
					"id":      incoming.ID,
				})
				return
			}
		}
	}

	// 3. Eliminar permissions???

	// Confirmar transacción
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al confirmar cambios", "error": err.Error()})
		return
	}

	// Responder con éxito
	c.JSON(http.StatusOK, response)
}

// POST: /apis/v1/systems/:system_id/users/:user_id
func SystemSavePermissionsUsers(c *gin.Context) {
	// 1. Validar y convertir IDs
	systemID, err := strconv.ParseUint(c.Param("system_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de sistema inválido",
			"error":   err.Error(),
		})
		return
	}

	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID de usuario inválido",
			"error":   err.Error(),
		})
		return
	}

	// 2. Parsear el cuerpo JSON
	var req models.UserPermissionSystemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error al parsear los datos",
			"error":   err.Error(),
		})
		return
	}

	// 3. Conectar a la base de datos
	if err := configs.ConnectToDB(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error de conexión a la base de datos",
			"error":   err.Error(),
		})
		return
	}

	// 4. Iniciar transacción
	tx := configs.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 5. Procesar cada edición
	response := make([]models.CreatedPermissionResponse, 0)

	for _, incoming := range req.Edits {
		var existing models.SystemUserPermission
		err := tx.Where("permission_id = ? AND user_id = ? AND system_id = ?",
			incoming.PermissionID, uint(userID), uint(systemID)).
			First(&existing).Error

		if incoming.Registered {
			// Caso: Debe existir la relación
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					// Crear nueva relación
					newRelation := models.SystemUserPermission{
						UserID:       uint(userID),
						SystemID:     uint(systemID),
						PermissionID: incoming.PermissionID,
						Created:      time.Now(),
					}

					if err := tx.Create(&newRelation).Error; err != nil {
						tx.Rollback()
						c.JSON(http.StatusInternalServerError, gin.H{
							"message": "Error al crear relación de permiso",
							"error":   err.Error(),
							"details": fmt.Sprintf("system_id: %d, user_id: %d, permission_id: %d",
								systemID, userID, incoming.PermissionID),
						})
						return
					}
				} else {
					tx.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": "Error al verificar permiso existente",
						"error":   err.Error(),
					})
					return
				}
			}
		} else {
			// Caso: NO debe existir la relación
			if err == nil {
				// Eliminar relación existente
				deleteResult := tx.Where("user_id = ? AND system_id = ? AND permission_id = ?",
					userID, systemID, incoming.PermissionID).
					Delete(&models.SystemUserPermission{}).Error
				if deleteResult != nil {
					tx.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": "Error al eliminar relación de permiso",
						"error":   deleteResult.Error(),
						"details": fmt.Sprintf("system_id: %d, user_id: %d, permission_id: %d",
							systemID, userID, incoming.PermissionID),
					})
					return
				}
			} else if !errors.Is(err, gorm.ErrRecordNotFound) {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error al verificar permiso para eliminar",
					"error":   err.Error(),
				})
				return
			}
			// Si no existe, no hacemos nada
		}
	}

	// 6. Confirmar transacción
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al confirmar cambios",
			"error":   err.Error(),
		})
		return
	}

	// 7. Responder con éxito
	c.JSON(http.StatusOK, response)
}

// GET: /apis/v1/systems-permissions/:system_id/users/:user_id
func SystemPermissionFetchUsers(c *gin.Context) {
	var permissions []models.UserPermissionWithRegistrationStatus

	systemIdStr := c.Param("system_id")
	var systemID uint
	if _, err := fmt.Sscanf(systemIdStr, "%d", &systemID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de sistema inválido"})
		return
	}

	idUserStr := c.Param("user_id")
	var userID uint
	if _, err := fmt.Sscanf(idUserStr, "%d", &userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de sistema inválido"})
		return
	}

	idRoleStr := c.Param("role_id")
	var roleID uint
	if _, err := fmt.Sscanf(idRoleStr, "%d", &roleID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de sistema inválido"})
		return
	}

	// Conexión a DB (asegúrate de que esto devuelva un *gorm.DB válido)
	if err := configs.ConnectToDB(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Conexión fallida",
			"message": err.Error(),
		})
		return
	}

	// Construimos la parte base de la consulta
	baseQuery := `
		SELECT 
			P.id,
			P.name,
			CASE WHEN SU.system_id IS NOT NULL THEN TRUE ELSE FALSE END AS registered
		FROM permissions P
		LEFT JOIN systems_users_permissions SU ON P.id = SU.permission_id AND SU.system_id = ? AND SU.user_id = ?
		INNER JOIN roles R ON R.id = P.role_id 
		WHERE R.id = ?`

	// 4. Ejecutar consulta
	if err := configs.DB.Raw(baseQuery, systemID, userID, roleID).Scan(&permissions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al obtener permisos",
			"message": err.Error(),
		})
		return
	}

	// Respuesta
	c.JSON(http.StatusOK, permissions)
}
