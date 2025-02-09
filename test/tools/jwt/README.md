# jwt-beego
A simple implementation of dgrijalva/jwt-go for beego.

Steps to implement:

1) Create a RSA key pair with commands used in file ***generar_key.log***

2) Generate a REST "path"/route with POST action in specific controller to return token, in this case the route was /user/getToken:

```go
// ./controllers/user.go
package controllers

import (
	...

	"jf_tools/jwt"
)

...

// @Title getToken
// @Description Get token from user session
// @Param	username		query 	string	true		"The username for get token"
// @Param	password		query 	string	true		"The password for get token"
// @Success 200 {string} Obtain Token
// @router /getToken [post]
func (u *UserController) GetToken() {
	username := u.Ctx.Input.Query("username")
	password := u.Ctx.Input.Query("password")

	tokenString := ""
	if username == "admin" && password == "mipassword" {
		et := jwt.EasyToken{
			Username: username,
			Expires:  time.Now().Unix() + 3600, //Segundos
		}
		tokenString, _ = et.GetToken()
	}

	u.Data["json"] = "{'tokenString': '" + tokenString + "'}"
	u.ServeJSON()
	return
}

...
```

3) Add a validation token in each controller that needs it. This is execute through ***Prepare*** function.

```go
// ./controllers/my_restricted_controller.go
package controllers

import (
	...

	"jf_tools/jwt"
)

func (c *TipoCancelacionSemestreController) Prepare() {
	tokenString := c.Ctx.Input.Query("tokenString")
	// O puede ser leído de una cabecera HEADER!!
	// tokenString := c.Ctx.Request.Header.Get("X-JWTtoken")

	et := jwt.EasyToken{}
	valido, _, _ := et.ValidateToken(tokenString)
	if !valido {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = "Permission Deny"
		c.ServeJSON()
	}
	return
}

...
```

With that this is finished, but if you want to all controllers have the same validation can you do:

1) Configure a new package.

```go
//./myBeego/controller.go

//Se crea un espacio de nombres llamado myBeego
package myBeego

//Se agrega la biblioteca de beego
import (
	...
	"github.com/astaxie/beego"
)

//Se genera un tipo Controller que hereda de beego.Controller
type Controller struct {
	beego.Controller
}

//Se redefine lo que hace la función Prepare
//* es un apuntador al igual que en C
//& hace referencia a la dirección de memoria
//La iniciación de una variable o funcion con * se traduce en que almacena
//u := 10 //var z *int  //z = &u //fmt.Println(z)//0x1040e0f8
//var s *string //var r **string = &s //fmt.Println(r)//0x1040a120
func (c *Controller) Prepare() {
	//Lo que quieras hacer en todos los controladores
	tokenString := c.Ctx.Input.Query("tokenString")
	// O puede ser leído de una cabecera HEADER!!
	// tokenString := c.Ctx.Request.Header.Get("X-JWTtoken")

	et := jwtbeego.EasyToken{}
	valido, _, _ := et.ValidateToken(tokenString)
	if !valido {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = "Permission Deny"
		c.ServeJSON()
	}
	return
}

```