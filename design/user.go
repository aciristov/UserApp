package design

import (
        . "github.com/goadesign/goa/design" // Use . imports to enable the DSL
        . "github.com/goadesign/goa/design/apidsl"
)

var _ = API("user", func(){
        Title("Goa user example title")
        Description("Goa user example description")
        Scheme("http")
        Host("localhost:8080")
})

var _ = Resource("user", func(){
        BasePath("/user")
        DefaultMedia(UserMediaType)

        Action("GetUser", func(){
                Description("Returns user by ID")
                Routing(GET("/:userID"))
                Params(func(){
                        Param("userID", Integer, "user ID")
                })
                Response(OK)
                Response(NotFound)
        })

        Action("CreateUser", func(){
                Description("Create user with POST method")
                Routing(POST(""))
                Payload(func(){
                        Member("username", String)
                        Member("password", String)
                        Member("email", String)
                })
                Response(OK)
                Response(BadRequest)
        })
})

var UserMediaType = MediaType("application/example.user+json", func ()  {
        Description("User")

        Attributes(func(){
                Attribute("id", Integer, "User ID")
                Attribute("username", String, "User name")
                Attribute("password", String, "User password")
                Attribute("email", String, "User email")
        })

        View("default", func(){
                Attribute("id", Integer, "User ID")
                Attribute("username", String, "User name")
                Attribute("email", String, "User email")
        })

})
