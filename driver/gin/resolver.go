package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

const(
	param = "param"
)

type resolve func(c context.Context, gc *gin.Context, field reflect.StructField) (reflect.Value, error)

var resolves = map[reflect.Type]resolve{
	reflect.TypeOf(context.Background()): contextResolver,
	reflect.TypeOf(gin.Context{}): ginContextResolver,
	reflect.TypeOf(reflect.Type(int8(0))): int8Resolver,
}

func contextResolver(c context.Context, gc *gin.Context, field reflect.StructField) (reflect.Value, error){
	return reflect.ValueOf(c), nil
}

func ginContextResolver(c context.Context, gc *gin.Context, field reflect.StructField) (reflect.Value, error){
	return reflect.ValueOf(*gc), nil
}

func int8Resolver(c context.Context, gc *gin.Context, field reflect.StructField) (reflect.Value, error){
	paramTag := field.Tag.Get(param)
	if paramTag == ""{
		return reflect.New(field.Type), nil
	}
	infos := strings.Split(paramTag, ",")
	source := infos[0]
	name := field.Name
	if len(infos) > 1{
		name = infos[1]
	}
	switch source {
	case "query":
		gc.Query(name)
	}
	return reflect.ValueOf(int8(0)), nil
}

