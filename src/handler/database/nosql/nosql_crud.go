package nosql

import (
	"context"
	"net/http"
	"reflect"

	"serviceX/src/handler/log"
	"serviceX/src/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (s db) Create(c *gin.Context, data model.ModelI) {
	name := reflect.TypeOf(data).Elem().Name() // using the struct name as a collection name
	//flatting the  struct
	b, _ := bson.Marshal(data)
	println(string(b))
	_, err := s.conn.Collection(name).InsertOne(context.Background(), b, nil)
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, data)
}

// FindOne
func (s db) FindOne(c *gin.Context, data model.ModelI) {

}

// FindAll
func (s db) FindAll(c *gin.Context, request model.FindAllRequest, modelType model.ModelI, data any) {

}
func (s db) Update(c *gin.Context, modelType model.ModelI, data map[string]any) {

}
func (s db) Delete(c *gin.Context, data model.ModelI) {

}
