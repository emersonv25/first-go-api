package controllers

import (
	"fmt"

	"first-go-api/dtos"
	"first-go-api/inits"
	"first-go-api/models"

	"github.com/gin-gonic/gin"
)

// CreatePost cria um novo post
// @Summary      Cria um post
// @Description  Cria um novo post a partir dos dados fornecidos no body
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        post body dtos.Post true "Post Data"
// @Success      200  {object}  dtos.SuccessResponse{data=dtos.PostResponse}
// @Failure      500  {object}  dtos.ErrorResponse
// @Router       /posts [post]
func CreatePost(ctx *gin.Context) {
	var body dtos.Post

	ctx.BindJSON(&body)

	post := models.Post{
		Title:  body.Title,
		Body:   body.Body,
		Likes:  body.Likes,
		Draft:  body.Draft,
		Author: body.Author,
	}

	fmt.Println(post)

	result := inits.DB.Create(&post)
	if result.Error != nil {
		ctx.JSON(500, dtos.ErrorResponse{Error: result.Error.Error()})

		return
	}

	response := dtos.NewPostResponse(post)

	ctx.JSON(200, dtos.SuccessResponse{Data: response})
}

// GetPosts retorna todos os posts
// @Summary      Retorna todos os posts
// @Description  Recupera todos os posts disponíveis
// @Tags         posts
// @Produce      json
// @Success      200  {object}  dtos.SuccessResponse{data=[]dtos.PostResponse}
// @Failure      500  {object}  dtos.ErrorResponse
// @Router       /posts [get]
func GetPosts(ctx *gin.Context) {
	var posts []models.Post

	result := inits.DB.Find(&posts)
	if result.Error != nil {
		ctx.JSON(500, dtos.ErrorResponse{Error: result.Error.Error()})
		return
	}

	response := dtos.NewPostResponseList(posts)

	ctx.JSON(200, dtos.SuccessResponse{Data: response})

}

// GetPost retorna um post específico
// @Summary      Retorna um post
// @Description  Recupera um post pelo ID fornecido
// @Tags         posts
// @Produce      json
// @Param        id   path      int  true  "ID do Post"
// @Success      200  {object}  dtos.SuccessResponse{data=dtos.PostResponse}
// @Failure      500  {object}  dtos.ErrorResponse
// @Router       /posts/{id} [get]
func GetPost(ctx *gin.Context) {
	id := ctx.Param("id")

	var post models.Post

	result := inits.DB.First(&post, id)
	if result.Error != nil {
		ctx.JSON(500, dtos.ErrorResponse{Error: result.Error.Error()})
		return
	}

	response := dtos.NewPostResponse(post)
	ctx.JSON(200, dtos.SuccessResponse{Data: response})

}

// UpdatePost atualiza um post
// @Summary      Atualiza um post
// @Description  Atualiza um post com base no ID e dados fornecidos
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path      int          true  "ID do Post"
// @Param        post body      dtos.Post true  "Post Data"
// @Success      200  {object}  dtos.SuccessResponse{data=dtos.PostResponse}
// @Failure      500  {object}  dtos.ErrorResponse
// @Router       /posts/{id} [put]
func UpdatePost(ctx *gin.Context) {
	var body dtos.Post

	ctx.BindJSON(&body)

	var post models.Post

	result := inits.DB.First(&post, ctx.Param("id"))

	if result.Error != nil {
		ctx.JSON(500, dtos.ErrorResponse{Error: result.Error.Error()})
		return
	}

	inits.DB.Model(&post).Updates(models.Post{
		Title:  body.Title,
		Body:   body.Body,
		Likes:  body.Likes,
		Draft:  body.Draft,
		Author: body.Author,
	})

	response := dtos.NewPostResponse(post)

	ctx.JSON(200, dtos.SuccessResponse{Data: response})
}

// DeletePost exclui um post
// @Summary      Exclui um post
// @Description  Remove um post baseado no ID fornecido
// @Tags         posts
// @Produce      json
// @Param        id   path      int  true  "ID do Post"
// @Success      200  {string}  string  "Post has been deleted successfully"
// @Failure      500  {object}  dtos.ErrorResponse
// @Router       /posts/{id} [delete]
func DeletePost(ctx *gin.Context) {

	id := ctx.Param("id")

	result := inits.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		ctx.JSON(500, dtos.ErrorResponse{Error: result.Error.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": "post has been deleted successfully"})

}
