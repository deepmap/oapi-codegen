package openapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindReferences(t *testing.T) {
	t.Run("unfiltered", func(t *testing.T) {
		swagger, err := LoadFromData([]byte(pruneSpecTestFixture))
		assert.NoError(t, err)

		refs := findComponentRefs(swagger)
		assert.Len(t, refs, 14)
	})
	t.Run("only cat", func(t *testing.T) {
		swagger, err := LoadFromData([]byte(pruneSpecTestFixture))
		assert.NoError(t, err)
		opts := Configuration{
			OutputOptions: OutputOptions{
				IncludeTags: []string{"cat"},
			},
		}

		filterOperationsByTag(swagger, opts)

		refs := findComponentRefs(swagger)
		assert.Len(t, refs, 7)
	})
	t.Run("only dog", func(t *testing.T) {
		swagger, err := LoadFromData([]byte(pruneSpecTestFixture))
		assert.NoError(t, err)

		opts := Configuration{
			OutputOptions: OutputOptions{
				IncludeTags: []string{"dog"},
			},
		}

		filterOperationsByTag(swagger, opts)

		refs := findComponentRefs(swagger)
		assert.Len(t, refs, 7)
	})
}

func TestFilterOnlyCat(t *testing.T) {
	// Get a spec from the test definition in this file:
	swagger, err := LoadFromData([]byte(pruneSpecTestFixture))
	assert.NoError(t, err)

	opts := Configuration{
		OutputOptions: OutputOptions{
			IncludeTags: []string{"cat"},
		},
	}

	refs := findComponentRefs(swagger)
	assert.Len(t, refs, 14)

	assert.Equal(t, 5, swagger.Model.Components.Schemas.Len())

	filterOperationsByTag(swagger, opts)

	refs = findComponentRefs(swagger)
	assert.Len(t, refs, 7)

	assert.NotEmpty(t, swagger.Model.Paths.PathItems.Value("/cat"), "/cat path should still be in spec")
	assert.NotEmpty(t, swagger.Model.Paths.PathItems.Value("/cat").Get, "GET /cat operation should still be in spec")
	assert.Empty(t, swagger.Model.Paths.PathItems.Value("/dog").Get, "GET /dog should have been removed from spec")

	pruneUnusedComponents(swagger)

	assert.Equal(t, 3, swagger.Model.Components.Schemas.Len())
}

func TestFilterOnlyDog(t *testing.T) {
	// Get a spec from the test definition in this file:
	swagger, err := LoadFromData([]byte(pruneSpecTestFixture))
	assert.NoError(t, err)

	opts := Configuration{
		OutputOptions: OutputOptions{
			IncludeTags: []string{"dog"},
		},
	}

	refs := findComponentRefs(swagger)
	assert.Len(t, refs, 14)

	filterOperationsByTag(swagger, opts)

	refs = findComponentRefs(swagger)
	assert.Len(t, refs, 7)

	assert.Equal(t, 5, swagger.Model.Components.Schemas.Len())

	assert.NotEmpty(t, swagger.Model.Paths.PathItems.Value("/dog"))
	assert.NotEmpty(t, swagger.Model.Paths.PathItems.Value("/dog").Get)
	assert.Empty(t, swagger.Model.Paths.PathItems.Value("/cat").Get)

	pruneUnusedComponents(swagger)

	assert.Equal(t, 3, swagger.Model.Components.Schemas.Len())
}

func TestPruningUnusedComponents(t *testing.T) {
	// Get a spec from the test definition in this file:
	swagger, err := LoadFromData([]byte(pruneComprehensiveTestFixture))
	assert.NoError(t, err)

	assert.Equal(t, 8, swagger.Model.Components.Schemas.Len())
	assert.Equal(t, 1, swagger.Model.Components.Parameters.Len())
	assert.Equal(t, 2, swagger.Model.Components.SecuritySchemes.Len())
	assert.Equal(t, 1, swagger.Model.Components.RequestBodies.Len())
	assert.Equal(t, 2, swagger.Model.Components.Responses.Len())
	assert.Equal(t, 3, swagger.Model.Components.Headers.Len())
	assert.Equal(t, 1, swagger.Model.Components.Examples.Len())
	assert.Equal(t, 1, swagger.Model.Components.Links.Len())
	assert.Equal(t, 1, swagger.Model.Components.Callbacks.Len())

	pruneUnusedComponents(swagger)

	assert.Equal(t, 0, swagger.Model.Components.Schemas.Len())
	assert.Equal(t, 0, swagger.Model.Components.Parameters.Len())
	// securitySchemes are an exception. definitions in securitySchemes
	// are referenced directly by name. and not by $ref
	assert.Equal(t, 2, swagger.Model.Components.SecuritySchemes.Len())
	assert.Equal(t, 0, swagger.Model.Components.RequestBodies.Len())
	assert.Equal(t, 0, swagger.Model.Components.Responses.Len())
	assert.Equal(t, 0, swagger.Model.Components.Headers.Len())
	assert.Equal(t, 0, swagger.Model.Components.Examples.Len())
	assert.Equal(t, 0, swagger.Model.Components.Links.Len())
	assert.Equal(t, 0, swagger.Model.Components.Callbacks.Len())
}

const pruneComprehensiveTestFixture = `
openapi: 3.0.1

info:
  title: OpenAPI-CodeGen Test
  description: 'This is a test OpenAPI Spec'
  version: 1.0.0

servers:
- url: https://test.oapi-codegen.com/v2
- url: http://test.oapi-codegen.com/v2

paths:
  /test:
    get:
      operationId: doesNothing
      summary: does nothing
      tags: [nothing]
      responses:
        default:
          description: returns nothing
          content:
            application/json:
              schema:
                type: object
components:
  schemas:
    Object1:
      type: object
      properties:
        object:
          $ref: "#/components/schemas/Object2"
    Object2:
      type: object
      properties:
        object:
          $ref: "#/components/schemas/Object3"
    Object3:
      type: object
      properties:
        object:
          $ref: "#/components/schemas/Object4"
    Object4:
      type: object
      properties:
        object:
          $ref: "#/components/schemas/Object5"
    Object5:
      type: object
      properties:
        object:
          $ref: "#/components/schemas/Object6"
    Object6:
      type: object
    Pet:
      type: object
      required:
        - id
        - name
      properties:
        id:
          type: integer
          format: int64
        name:
          type: string
        tag:
          type: string
    Error:
      required:
        - code
        - message
      properties:
        code:
          type: integer
          format: int32
          description: Error code
        message:
          type: string
          description: Error message
  parameters:
    offsetParam:
      name: offset
      in: query
      description: Number of items to skip before returning the results.
      required: false
      schema:
        type: integer
        format: int32
        minimum: 0
        default: 0
  securitySchemes:
    BasicAuth:
      type: http
      scheme: basic
    BearerAuth:
      type: http
      scheme: bearer
  requestBodies:
    PetBody:
      description: A JSON object containing pet information
      required: true
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/Pet'
  responses:
    NotFound:
      description: The specified resource was not found
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/Error'
    Unauthorized:
      description: Unauthorized
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/Error'
  headers:
    X-RateLimit-Limit:
      schema:
        type: integer
      description: Request limit per hour.
    X-RateLimit-Remaining:
      schema:
        type: integer
      description: The number of requests left for the time window.
    X-RateLimit-Reset:
      schema:
        type: string
        format: date-time
      description: The UTC date/time at which the current rate limit window resets
  examples:
    objectExample:
      value:
        id: 1
        name: new object
      summary: A sample object
  links:
    GetUserByUserId:
      description: >
        The id value returned in the response can be used as
        the userId parameter in GET /users/{userId}.
      operationId: getUser
      parameters:
        userId: '$response.body#/id'
  callbacks:
    MyCallback:
      '{$request.body#/callbackUrl}':
        post:
          requestBody:
            required: true
            content:
              application/json:
                schema:
                  type: object
                  properties:
                    message:
                      type: string
                      example: Some event happened
                  required:
                    - message
          responses:
            '200':
              description: Your server returns this code if it accepts the callback
`

const pruneSpecTestFixture = `
openapi: 3.0.1

info:
  title: OpenAPI-CodeGen Test
  description: 'This is a test OpenAPI Spec'
  version: 1.0.0

servers:
- url: https://test.oapi-codegen.com/v2
- url: http://test.oapi-codegen.com/v2

paths:
  /cat:
    get:
      tags:
        - cat
      summary: Get cat status
      operationId: getCatStatus
      responses:
        200:
          description: Success
          content:
            application/json:
              schema:
                oneOf:
                  - $ref: '#/components/schemas/CatAlive'
                  - $ref: '#/components/schemas/CatDead'
            application/xml:
              schema:
                anyOf:
                  - $ref: '#/components/schemas/CatAlive'
                  - $ref: '#/components/schemas/CatDead'
            application/yaml:
              schema:
                allOf:
                  - $ref: '#/components/schemas/CatAlive'
                  - $ref: '#/components/schemas/CatDead'
        default:
          description: Error
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'

  /dog:
    get:
      tags:
        - dog
      summary: Get dog status
      operationId: getDogStatus
      responses:
        200:
          description: Success
          content:
            application/json:
              schema:
                oneOf:
                  - $ref: '#/components/schemas/DogAlive'
                  - $ref: '#/components/schemas/DogDead'
            application/xml:
              schema:
                anyOf:
                  - $ref: '#/components/schemas/DogAlive'
                  - $ref: '#/components/schemas/DogDead'
            application/yaml:
              schema:
                allOf:
                  - $ref: '#/components/schemas/DogAlive'
                  - $ref: '#/components/schemas/DogDead'
        default:
          description: Error
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'

components:
  schemas:

    Error:
      properties:
        code:
          type: integer
          format: int32
        message:
          type: string

    CatAlive:
      properties:
        name:
          type: string
        alive_since:
          type: string
          format: date-time

    CatDead:
      properties:
        name:
          type: string
        dead_since:
          type: string
          format: date-time
        cause:
          type: string
          enum: [car, dog, oldage]

    DogAlive:
      properties:
        name:
          type: string
        alive_since:
          type: string
          format: date-time

    DogDead:
      properties:
        name:
          type: string
        dead_since:
          type: string
          format: date-time
        cause:
          type: string
          enum: [car, cat, oldage]

`
