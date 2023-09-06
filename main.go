// Ortelius v11 Group Microservice that handles creating and retrieving Groups
package main

import (
	"context"
	"encoding/json"

	_ "cli/docs"

	"github.com/arangodb/go-driver"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/ortelius/scec-commons/database"
	"github.com/ortelius/scec-commons/model"
)

var logger = database.InitLogger()
var dbconn = database.InitializeDB("evidence")

// GetGroups godoc
// @Summary Get a List of Groups
// @Description Get a list of groups for the user.
// @Tags group
// @Accept */*
// @Produce json
// @Success 200
// @Router /msapi/group [get]
func GetGroups(c *fiber.Ctx) error {

	var cursor driver.Cursor       // db cursor for rows
	var err error                  // for error handling
	var ctx = context.Background() // use default database context

	// query all the groups in the collection
	aql := `FOR group in evidence
			FILTER (group.objtype == 'Group')
			RETURN group`

	// execute the query with no parameters
	if cursor, err = dbconn.Database.Query(ctx, aql, nil); err != nil {
		logger.Sugar().Errorf("Failed to run query: %v", err) // log error
	}

	defer cursor.Close() // close the cursor when returning from this function

<<<<<<< HEAD
	var groups []*model.Group // define a list of groups to be returned

	for cursor.HasMore() { // loop thru all of the documents

		group := model.NewGroup()    // fetched group
		var meta driver.DocumentMeta // data about the fetch
=======
	groups := model.NewGroups() // define a list of groups to be returned

	for cursor.HasMore() { // loop thru all of the documents

		group := model.NewGroup() // fetched group
		var meta driver.DocumentMeta           // data about the fetch
>>>>>>> 3c175a4 (update object names)

		// fetch a document from the cursor
		if meta, err = cursor.ReadDocument(ctx, group); err != nil {
			logger.Sugar().Errorf("Failed to read document: %v", err)
		}
<<<<<<< HEAD
		groups = append(groups, group)                                       // add the group to the list
=======
		groups.Groups = append(groups.Groups, group)       // add the group to the list
>>>>>>> 3c175a4 (update object names)
		logger.Sugar().Infof("Got doc with key '%s' from query\n", meta.Key) // log the key
	}

	return c.JSON(groups) // return the list of groups in JSON format
}

// GetGroup godoc
// @Summary Get a Group
// @Description Get a group based on the _key or name.
// @Tags group
// @Accept */*
// @Produce json
// @Success 200
// @Router /msapi/group/:key [get]
func GetGroup(c *fiber.Ctx) error {

	var cursor driver.Cursor       // db cursor for rows
	var err error                  // for error handling
	var ctx = context.Background() // use default database context

	key := c.Params("key")                // key from URL
	parameters := map[string]interface{}{ // parameters
		"key": key,
	}

	// query the groups that match the key or name
	aql := `FOR group in evidence
			FILTER (group.name == @key or group._key == @key)
			RETURN group`

	// run the query with patameters
	if cursor, err = dbconn.Database.Query(ctx, aql, parameters); err != nil {
		logger.Sugar().Errorf("Failed to run query: %v", err)
	}

	defer cursor.Close() // close the cursor when returning from this function

<<<<<<< HEAD
	group := model.NewGroup() // define a group to be returned
=======
	group := model.NewGroups() // define a group to be returned
>>>>>>> 3c175a4 (update object names)

	if cursor.HasMore() { // group found
		var meta driver.DocumentMeta // data about the fetch

		if meta, err = cursor.ReadDocument(ctx, group); err != nil { // fetch the document into the object
			logger.Sugar().Errorf("Failed to read document: %v", err)
		}
		logger.Sugar().Infof("Got doc with key '%s' from query\n", meta.Key)

	} else { // not found so get from NFT Storage
		if jsonStr, exists := database.MakeJSON(key); exists {
			if err := json.Unmarshal([]byte(jsonStr), group); err != nil { // convert the JSON string from LTF into the object
				logger.Sugar().Errorf("Failed to unmarshal from LTS: %v", err)
			}
		}
	}

	return c.JSON(group) // return the group in JSON format
}

// NewGroup godoc
// @Summary Create a Group
// @Description Create a new Group and persist it
// @Tags group
// @Accept application/json
// @Produce json
// @Success 200
// @Router /msapi/group [post]
func NewGroup(c *fiber.Ctx) error {

	var err error                  // for error handling
	var meta driver.DocumentMeta   // data about the document
	var ctx = context.Background() // use default database context
	group := new(model.Group)      // define a group to be returned

	if err = c.BodyParser(group); err != nil { // parse the JSON into the group object
		return c.Status(503).Send([]byte(err.Error()))
	}

	cid, dbStr := database.MakeNFT(group) // normalize the object into NFTs and JSON string for db persistence

	logger.Sugar().Infof("%s=%s\n", cid, dbStr) // log the new nft

	// add the group to the database.  Ignore if it already exists since it will be identical
	if meta, err = dbconn.Collection.CreateDocument(ctx, group); err != nil && !driver.IsConflict(err) {
		logger.Sugar().Errorf("Failed to create document: %v", err)
	}
	logger.Sugar().Infof("Created document in collection '%s' in db '%s' key='%s'\n", dbconn.Collection.Name(), dbconn.Database.Name(), meta.Key)

	return c.JSON(group) // return the group object in JSON format.  This includes the new _key
}

// setupRoutes defines maps the routes to the functions
func setupRoutes(app *fiber.App) {

	app.Get("/swagger/*", swagger.HandlerDefault) // handle displaying the swagger
	app.Get("/msapi/group", GetGroup)             // list of groups
	app.Get("/msapi/group/:key", GetGroup)        // single group based on name or key
	app.Post("/msapi/group", NewGroup)            // save a single group
}

// @title Ortelius v11 Group Microservice
// @version 11.0.0
// @description RestAPI for the Group Object
// @description ![Release](https://img.shields.io/github/v/release/ortelius/scec-group?sort=semver)
// @description ![license](https://img.shields.io/github/license/ortelius/scec-group)
// @description
// @description ![Build](https://img.shields.io/github/actions/workflow/status/ortelius/scec-group/build-push-chart.yml)
// @description [![MegaLinter](https://github.com/ortelius/scec-group/workflows/MegaLinter/badge.svg?branch=main)](https://github.com/ortelius/scec-group/actions?query=workflow%3AMegaLinter+branch%3Amain)
// @description ![CodeQL](https://github.com/ortelius/scec-group/workflows/CodeQL/badge.svg)
// @description [![OpenSSF-Scorecard](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-group/badge)](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-group)
// @description
// @description ![Discord](https://img.shields.io/discord/722468819091849316)

// @termsOfService http://swagger.io/terms/
// @contact.name Ortelius Google Group
// @contact.email ortelius-dev@googlegroups.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /msapi/group
func main() {
	port := ":" + database.GetEnvDefault("MS_PORT", "8080") // database port
	app := fiber.New()                                      // create a new fiber application
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowOrigins: "*",
	}))

	setupRoutes(app) // define the routes for this microservice

	if err := app.Listen(port); err != nil { // start listening for incoming connections
		logger.Sugar().Fatalf("Failed get the microservice running: %v", err)
	}
}
