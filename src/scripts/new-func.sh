#!/bin/bash

# Read the file name from the command line
# check if the file name is provided if not show error message
if [ -z "$1" ]
then
    echo "Please provide the file name"
    exit 1
fi
file_name=$1

# if the file src/internal/api/handlers/${file_name}_handler.go is already exists then show error message
if [ -f src/internal/api/handlers/${file_name}_handler.go ]
then
    echo "File already exists"
    exit 1
fi

# Create the handler file
name_cap=$(echo $file_name | awk '{print toupper(substr($0,1,1)) substr($0,2)}')
dir_name=$(pwd) 
cat > src/internal/api/handlers/${file_name}_handler.go << EOF
package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/validation"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/services"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
)

/*
	"""
	${name_cap}Handler can provide the following services.
	"""
*/

type ${name_cap}Handler interface {
	Create${name_cap}(c *gin.Context)
}

type ${file_name}Handler struct {
  validator            validation.Validator
	${file_name}Mapper   mappers.${name_cap}Mapper
	${file_name}Services services.${name_cap}Services
}

/*
	"""
	This will create a new instance of the ${name_cap}Handler, we will use this as a constructor
	"""
*/

func New${name_cap}Handler(service services.${name_cap}Services, mapper mappers.${name_cap}Mapper,v validation.Validator) ${name_cap}Handler {
	return &${file_name}Handler{
		${file_name}Mapper:   mapper,
		${file_name}Services: service,
		validator:            v,
	}
}

func (s *${file_name}Handler) Create${name_cap}(c *gin.Context) {}
EOF

# Create a router file
cat > src/internal/api/routers/${file_name}_router.go << EOF
package routers
import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/api/handlers"
)
func ${name_cap}Router(r *gin.RouterGroup, h handlers.${name_cap}Handler) {

	g := r.Group("/${file_name}")

	g.POST("/create", h.Create${name_cap})
}

EOF

# # Create the services file
cat > src/internal/core/domain/services/${file_name}_services.go << EOF
package services

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/repositories"
)

type ${name_cap}Services interface {
	Create${name_cap}(dtos.${name_cap}CreateDTO) error
}

type ${file_name}Services struct {
	${file_name}Mapper     mappers.${name_cap}Mapper
	${file_name}Repository repositories.${name_cap}Repository
}

func New${name_cap}Services(repo repositories.${name_cap}Repository, mapper mappers.${name_cap}Mapper) ${name_cap}Services {
	return &${file_name}Services{
		${file_name}Repository: repo,
		${file_name}Mapper:     mapper,
	}
}

func (s *${file_name}Services) Create${name_cap}(${file_name} dtos.${name_cap}CreateDTO) error {

}
EOF

# # Create the mapper file
cat > src/internal/core/infrastructure/postgres/mappers/${file_name}_mapper.go << EOF
package mappers

type ${name_cap}Mapper interface {
    ${name_cap}CreateMapper()
}

type ${file_name}Mapper struct {
}

func New${name_cap}Mapper() ${name_cap}Mapper {
    return &${file_name}Mapper{}
}

func (m *${file_name}Mapper) ${name_cap}CreateMapper() {
}

EOF

# # Create the repository file

cat > src/internal/core/infrastructure/postgres/repositories/${file_name}_repository.go << EOF
package repositories

import "gorm.io/gorm"

type ${name_cap}Repository interface {
    Insert${name_cap}(entities.${name_cap}Entity) error
}

type ${file_name}Connection struct {
    conn *gorm.DB
}

func New${name_cap}Repository(db *gorm.DB) ${name_cap}Repository {
    return &${file_name}Connection{
        conn: db,
    }
}

func (r *${file_name}Connection) Insert${name_cap}(${file_name entities.${name_cap}Entity}) error {

}

EOF