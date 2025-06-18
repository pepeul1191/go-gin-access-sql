package controllers

import (
	"access/app/configs"
	"access/app/models"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GET: apis/v1/users
func UserFetchAll(c *gin.Context) {
	var users []models.User
	var total int64
	// Obtener parámetros opcionales
	name := c.Query("name")
	email := c.Query("email")
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
	query := configs.DB.Model(&models.User{})
	// Aplicar filtros opcionales
	if name != "" {
		query = query.Where("username LIKE ?", "%"+name+"%")
	}
	if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}
	// Paginación (si step y page están presentes)
	if step != "" && page != "" {
		// Contar total (antes de paginar)
		if err := query.Count(&total).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error al contar users",
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
			if err := query.Find(&users).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Error al listar la lista paginada de sistemas",
					"message": err.Error(),
				})
				return
			}
			// Respuesta
			pages := int(math.Ceil(float64(total) / float64(stepInt)))
			c.JSON(http.StatusOK, gin.H{
				"list":   users,
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
		if err := query.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error al consultar users",
				"message": err.Error(),
			})
			return
		}
		// Respuesta
		c.JSON(http.StatusOK, users)
	}
}

// POST: /apis/v1/users
func UserCreate(c *gin.Context) {
	var user models.User
	var input models.CreateUserInput
	// Parsear JSON recibido
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "message": err.Error()})
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
	// veficar si existe usuario y correo en uso
	var count int64
	if err := configs.DB.Model(&models.User{}).
		Where("username = ? OR email = ?", input.Username, input.Email).
		Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al verificar si usuario y contraseña ya existen", "message": err.Error()})
		return
	}
	if count == 0 {
		// grabar
		// Establecer fechas (si no se hace en frontend)
		user.Created = time.Now()
		user.Updated = time.Now()
		user.ResetKey = configs.HelperRandomString(30)
		user.ActivationKey = configs.HelperRandomString(30)
		user.Password = configs.HelperRandomString(30)
		user.Activated = false
		user.Email = input.Email
		user.Username = input.Username
		if err := configs.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el sistema", "message": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": user.ID})
	} else {
		// avisar que usuario y contraseña en uso
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Usuario y contraseña y en uso", "message": ""})
		return
	}
}

// PUT: /apis/v1/users/:id/password
func UserUpdatePassword(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	var input models.UpdatePasswordUserInput
	// Parsear JSON recibido
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "message": err.Error()})
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
	// Buscar el sistema existente
	if err := configs.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario para configurar contraseña no encontrado"})
		return
	}
	// actualizar
	user.Updated = time.Now()
	user.Password = input.Pasword
	if err := configs.DB.Model(&user).Select("password", "updated").Updates(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar", "message": err.Error()})
		return
	}
	// TODO: enviar correo
	c.JSON(http.StatusOK, "Contraseña actualizada")
}

// PUT: /apis/v1/users/:id/activation-key
func UserUpdateActivationKey(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	var input models.UpdateActivationKeyUserInput
	// Parsear JSON recibido
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "message": err.Error()})
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
	// Buscar el sistema existente
	if err := configs.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario para configurar cambio de clave de cambio de contraseña no encontrado"})
		return
	}
	// actualizar
	user.Updated = time.Now()
	user.ActivationKey = input.ActivationKey
	if err := configs.DB.Model(&user).Select("activation_key", "updated").Updates(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar", "message": err.Error()})
		return
	}
	// TODO: enviar correo
	c.JSON(http.StatusOK, "Clave de activación actualizada")
}

// PUT: /apis/v1/users/:id/reset-key
func UserUpdateResetKey(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	var input models.UpdateResetKeyUserInput
	// Parsear JSON recibido
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "message": err.Error()})
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
	// Buscar el sistema existente
	if err := configs.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario para configurar cambio de clave de cambio de contraseña no encontrado"})
		return
	}
	// actualizar
	user.Updated = time.Now()
	user.ResetKey = input.ResetKey
	if err := configs.DB.Model(&user).Select("reset_key", "updated").Updates(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar", "message": err.Error()})
		return
	}
	// TODO: enviar correo
	c.JSON(http.StatusOK, "Clave de reseto actualizada")
}

// PUT: /apis/v1/users/:id/activated
func UserUpdateActivated(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	var input models.UpdateActivatedUserInput
	// Parsear JSON recibido
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido", "message": err.Error()})
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
	// Buscar el sistema existente
	if err := configs.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario para configurar cambio estado de activdad de usuario"})
		return
	}
	// actualizar
	user.Updated = time.Now()
	user.Activated = input.Activated
	if err := configs.DB.Model(&user).Select("activated", "updated").Updates(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Cambio del esetado de activación actualizada")
}
