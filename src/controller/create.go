package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wolkenheim.cloud/thumbnail-generator/app"
	"wolkenheim.cloud/thumbnail-generator/service"
)

type CreateController struct {
	app *app.Application
	process service.ProcessFacade
	validate CustomValidator
}

// CreateRequest : describes JSON post body
type CreateRequest struct {
	FileName  string `json:"fileName" validate:"required,allowed-extensions"`
}

func(c *CreateController) Create(w http.ResponseWriter, r *http.Request){

	var createRequest CreateRequest
	err := json.NewDecoder(r.Body).Decode(&createRequest)
	if err != nil {
		c.app.JSON(w, http.StatusBadRequest, &app.MessageResponse{Message: "Invalid JSON"})
		return
	}

	// validate struct
	validationErrors := c.validate.ValidateStruct(createRequest)
	if validationErrors != nil {
		c.app.JSON(w, http.StatusBadRequest, validationErrors)
		return
	}

	go func(){
		c.process.ProcessImage(createRequest.FileName)
	}()

	c.app.JSON(w, http.StatusOK, &app.MessageResponse{Message: fmt.Sprintf("received %s", createRequest.FileName)})
	return
}

func(c *CreateController) SetApp(a *app.Application){
	c.app = a
}

func(c *CreateController) SetProcess(p service.ProcessFacade){
	c.process = p
}

func(c *CreateController) SetValidator(v CustomValidator){
	c.validate = v
}
