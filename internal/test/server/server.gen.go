// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/discord-gophers/goapi-gen version (devel) DO NOT EDIT.
package server

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/discord-gophers/goapi-gen/pkg/runtime"
	openapi_types "github.com/discord-gophers/goapi-gen/pkg/types"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// EveryTypeOptional defines model for EveryTypeOptional.
type EveryTypeOptional struct {
	ArrayInlineField     *[]int              `json:"array_inline_field,omitempty"`
	ArrayReferencedField *[]SomeObject       `json:"array_referenced_field,omitempty"`
	BoolField            *bool               `json:"bool_field,omitempty"`
	ByteField            *[]byte             `json:"byte_field,omitempty"`
	DateField            *openapi_types.Date `json:"date_field,omitempty"`
	DateTimeField        *time.Time          `json:"date_time_field,omitempty"`
	DoubleField          *float64            `json:"double_field,omitempty"`
	FloatField           *float32            `json:"float_field,omitempty"`
	InlineObjectField    *struct {
		Name   string `json:"name"`
		Number int    `json:"number"`
	} `json:"inline_object_field,omitempty"`
	Int32Field      *int32      `json:"int32_field,omitempty"`
	Int64Field      *int64      `json:"int64_field,omitempty"`
	IntField        *int        `json:"int_field,omitempty"`
	NumberField     *float32    `json:"number_field,omitempty"`
	ReferencedField *SomeObject `json:"referenced_field,omitempty"`
	StringField     *string     `json:"string_field,omitempty"`
}

// EveryTypeRequired defines model for EveryTypeRequired.
type EveryTypeRequired struct {
	ArrayInlineField     []int                `json:"array_inline_field"`
	ArrayReferencedField []SomeObject         `json:"array_referenced_field"`
	BoolField            bool                 `json:"bool_field"`
	ByteField            []byte               `json:"byte_field"`
	DateField            openapi_types.Date   `json:"date_field"`
	DateTimeField        time.Time            `json:"date_time_field"`
	DoubleField          float64              `json:"double_field"`
	EmailField           *openapi_types.Email `json:"email_field,omitempty"`
	FloatField           float32              `json:"float_field"`
	InlineObjectField    struct {
		Name   string `json:"name"`
		Number int    `json:"number"`
	} `json:"inline_object_field"`
	Int32Field      int32      `json:"int32_field"`
	Int64Field      int64      `json:"int64_field"`
	IntField        int        `json:"int_field"`
	NumberField     float32    `json:"number_field"`
	ReferencedField SomeObject `json:"referenced_field"`
	StringField     string     `json:"string_field"`
}

// ReservedKeyword defines model for ReservedKeyword.
type ReservedKeyword struct {
	Channel *string `json:"channel,omitempty"`
}

// Resource defines model for Resource.
type Resource struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}

// SomeObject defines model for some_object.
type SomeObject struct {
	Name string `json:"name"`
}

// Argument defines model for argument.
type Argument string

// ResponseWithReference defines model for ResponseWithReference.
type ResponseWithReference SomeObject

// SimpleResponse defines model for SimpleResponse.
type SimpleResponse struct {
	Name string `json:"name"`
}

// GetWithArgsParams defines parameters for GetWithArgs.
type GetWithArgsParams struct {
	// An optional query argument
	OptionalArgument *int64 `json:"optional_argument,omitempty"`

	// A required query argument
	RequiredArgument int64 `json:"required_argument"`

	// An optional query argument
	HeaderArgument *int32 `json:"header_argument,omitempty"`
}

// GetWithContentTypeParamsContentType defines parameters for GetWithContentType.
type GetWithContentTypeParamsContentType string

// CreateResourceJSONBody defines parameters for CreateResource.
type CreateResourceJSONBody EveryTypeRequired

// CreateResource2JSONBody defines parameters for CreateResource2.
type CreateResource2JSONBody Resource

// CreateResource2Params defines parameters for CreateResource2.
type CreateResource2Params struct {
	// Some query argument
	InlineQueryArgument *int `json:"inline_query_argument,omitempty"`
}

// UpdateResource3JSONBody defines parameters for UpdateResource3.
type UpdateResource3JSONBody struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CreateResourceJSONRequestBody defines body for CreateResource for application/json ContentType.
type CreateResourceJSONRequestBody CreateResourceJSONBody

// Bind implements render.Binder.
func (CreateResourceJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// CreateResource2JSONRequestBody defines body for CreateResource2 for application/json ContentType.
type CreateResource2JSONRequestBody CreateResource2JSONBody

// Bind implements render.Binder.
func (CreateResource2JSONRequestBody) Bind(*http.Request) error {
	return nil
}

// UpdateResource3JSONRequestBody defines body for UpdateResource3 for application/json ContentType.
type UpdateResource3JSONRequestBody UpdateResource3JSONBody

// Bind implements render.Binder.
func (UpdateResource3JSONRequestBody) Bind(*http.Request) error {
	return nil
}

// Response is a common response struct for all the API calls.
// A Response object may be instantiated via functions for specific operation responses.
type Response struct {
	body        interface{}
	statusCode  int
	contentType string
}

// Render implements the render.Renderer interface. It sets the Content-Type header
// and status code based on the response definition.
func (resp *Response) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", resp.contentType)
	render.Status(r, resp.statusCode)
	return nil
}

// Status is a builder method to override the default status code for a response.
func (resp *Response) Status(statusCode int) *Response {
	resp.statusCode = statusCode
	return resp
}

// ContentType is a builder method to override the default content type for a response.
func (resp *Response) ContentType(contentType string) *Response {
	resp.contentType = contentType
	return resp
}

// MarshalJSON implements the json.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(resp.body)
}

// MarshalXML implements the xml.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(resp.body)
}

// GetEveryTypeOptionalJSON200Response is a constructor method for a GetEveryTypeOptional response.
// A *Response is returned with the configured status code and content type from the spec.
func GetEveryTypeOptionalJSON200Response(body EveryTypeOptional) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// GetSimpleJSON200Response is a constructor method for a GetSimple response.
// A *Response is returned with the configured status code and content type from the spec.
func GetSimpleJSON200Response(body SomeObject) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// GetWithArgsJSON200Response is a constructor method for a GetWithArgs response.
// A *Response is returned with the configured status code and content type from the spec.
func GetWithArgsJSON200Response(body struct {
	Name string `json:"name"`
}) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// GetWithReferencesJSON200Response is a constructor method for a GetWithReferences response.
// A *Response is returned with the configured status code and content type from the spec.
func GetWithReferencesJSON200Response(body struct {
	Name string `json:"name"`
}) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// GetWithContentTypeJSON200Response is a constructor method for a GetWithContentType response.
// A *Response is returned with the configured status code and content type from the spec.
func GetWithContentTypeJSON200Response(body SomeObject) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// GetReservedKeywordJSON200Response is a constructor method for a GetReservedKeyword response.
// A *Response is returned with the configured status code and content type from the spec.
func GetReservedKeywordJSON200Response(body ReservedKeyword) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// CreateResourceJSON200Response is a constructor method for a CreateResource response.
// A *Response is returned with the configured status code and content type from the spec.
func CreateResourceJSON200Response(body struct {
	Name string `json:"name"`
}) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// CreateResource2JSON200Response is a constructor method for a CreateResource2 response.
// A *Response is returned with the configured status code and content type from the spec.
func CreateResource2JSON200Response(body struct {
	Name string `json:"name"`
}) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// UpdateResource3JSON200Response is a constructor method for a UpdateResource3 response.
// A *Response is returned with the configured status code and content type from the spec.
func UpdateResource3JSON200Response(body struct {
	Name string `json:"name"`
}) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// GetResponseWithReferenceJSON200Response is a constructor method for a GetResponseWithReference response.
// A *Response is returned with the configured status code and content type from the spec.
func GetResponseWithReferenceJSON200Response(body SomeObject) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// GetWithTaggedMiddlewareJSON200Response is a constructor method for a GetWithTaggedMiddleware response.
// A *Response is returned with the configured status code and content type from the spec.
func GetWithTaggedMiddlewareJSON200Response(body struct {
	Name string `json:"name"`
}) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// PostWithTaggedMiddlewareJSON200Response is a constructor method for a PostWithTaggedMiddleware response.
// A *Response is returned with the configured status code and content type from the spec.
func PostWithTaggedMiddlewareJSON200Response(body struct {
	Name string `json:"name"`
}) *Response {
	return &Response{
		body:        body,
		statusCode:  200,
		contentType: "application/json",
	}
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// get every type optional
	// (GET /every-type-optional)
	GetEveryTypeOptional(w http.ResponseWriter, r *http.Request)
	// Get resource via simple path
	// (GET /get-simple)
	GetSimple(w http.ResponseWriter, r *http.Request)
	// Getter with referenced parameter and referenced response
	// (GET /get-with-args)
	GetWithArgs(w http.ResponseWriter, r *http.Request, params GetWithArgsParams)
	// Getter with referenced parameter and referenced response
	// (GET /get-with-references/{global_argument}/{argument})
	GetWithReferences(w http.ResponseWriter, r *http.Request, globalArgument int64, argument Argument)
	// Get an object by ID
	// (GET /get-with-type/{content_type})
	GetWithContentType(w http.ResponseWriter, r *http.Request, contentType GetWithContentTypeParamsContentType)
	// get with reserved keyword
	// (GET /reserved-keyword)
	GetReservedKeyword(w http.ResponseWriter, r *http.Request)
	// Create a resource
	// (POST /resource/{argument})
	CreateResource(w http.ResponseWriter, r *http.Request, argument Argument)
	// Create a resource with inline parameter
	// (POST /resource2/{inline_argument})
	CreateResource2(w http.ResponseWriter, r *http.Request, inlineArgument int, params CreateResource2Params)
	// Update a resource with inline body. The parameter name is a reserved
	// keyword, so make sure that gets prefixed to avoid syntax errors
	// (PUT /resource3/{fallthrough})
	UpdateResource3(w http.ResponseWriter, r *http.Request, pFallthrough int)
	// get response with reference
	// (GET /response-with-reference)
	GetResponseWithReference(w http.ResponseWriter, r *http.Request)

	// (GET /with-tagged-middleware)
	GetWithTaggedMiddleware(w http.ResponseWriter, r *http.Request)

	// (POST /with-tagged-middleware)
	PostWithTaggedMiddleware(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler          ServerInterface
	Middlewares      map[string]func(http.Handler) http.Handler
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// GetEveryTypeOptional operation middleware
func (siw *ServerInterfaceWrapper) GetEveryTypeOptional(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetEveryTypeOptional(w, r)
	})

	handler(w, r.WithContext(ctx))
}

// GetSimple operation middleware
func (siw *ServerInterfaceWrapper) GetSimple(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSimple(w, r)
	})

	handler(w, r.WithContext(ctx))
}

// GetWithArgs operation middleware
func (siw *ServerInterfaceWrapper) GetWithArgs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parameter object where we will unmarshal all parameters from the context
	var params GetWithArgsParams

	// ------------- Optional query parameter "optional_argument" -------------

	if err := runtime.BindQueryParameter("form", true, false, "optional_argument", r.URL.Query(), &params.OptionalArgument); err != nil {
		err = fmt.Errorf("invalid format for parameter optional_argument: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err})
		return
	}

	// ------------- Required query parameter "required_argument" -------------

	if err := runtime.BindQueryParameter("form", true, true, "required_argument", r.URL.Query(), &params.RequiredArgument); err != nil {
		err = fmt.Errorf("invalid format for parameter required_argument: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err})
		return
	}

	headers := r.Header

	// ------------- Optional header parameter "header_argument" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("header_argument")]; found {
		var HeaderArgument int32
		n := len(valueList)
		if n != 1 {
			err := fmt.Errorf("expected one value for header_argument, got %d", n)
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{err})
			return
		}

		if err := runtime.BindStyledParameterWithLocation("simple", false, "header_argument", runtime.ParamLocationHeader, valueList[0], &HeaderArgument); err != nil {
			err = fmt.Errorf("invalid format for parameter header_argument: %w", err)
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err})
			return
		}

		params.HeaderArgument = &HeaderArgument

	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetWithArgs(w, r, params)
	})

	handler(w, r.WithContext(ctx))
}

// GetWithReferences operation middleware
func (siw *ServerInterfaceWrapper) GetWithReferences(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "global_argument" -------------
	var globalArgument int64

	if err := runtime.BindStyledParameter("simple", false, "global_argument", chi.URLParam(r, "global_argument"), &globalArgument); err != nil {
		err = fmt.Errorf("invalid format for parameter global_argument: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err})
		return
	}

	// ------------- Path parameter "argument" -------------
	var argument Argument

	if err := runtime.BindStyledParameter("simple", false, "argument", chi.URLParam(r, "argument"), &argument); err != nil {
		err = fmt.Errorf("invalid format for parameter argument: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetWithReferences(w, r, globalArgument, argument)
	})

	handler(w, r.WithContext(ctx))
}

// GetWithContentType operation middleware
func (siw *ServerInterfaceWrapper) GetWithContentType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "content_type" -------------
	var contentType GetWithContentTypeParamsContentType

	if err := runtime.BindStyledParameter("simple", false, "content_type", chi.URLParam(r, "content_type"), &contentType); err != nil {
		err = fmt.Errorf("invalid format for parameter content_type: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetWithContentType(w, r, contentType)
	})

	handler(w, r.WithContext(ctx))
}

// GetReservedKeyword operation middleware
func (siw *ServerInterfaceWrapper) GetReservedKeyword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetReservedKeyword(w, r)
	})

	handler(w, r.WithContext(ctx))
}

// CreateResource operation middleware
func (siw *ServerInterfaceWrapper) CreateResource(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "argument" -------------
	var argument Argument

	if err := runtime.BindStyledParameter("simple", false, "argument", chi.URLParam(r, "argument"), &argument); err != nil {
		err = fmt.Errorf("invalid format for parameter argument: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateResource(w, r, argument)
	})

	handler(w, r.WithContext(ctx))
}

// CreateResource2 operation middleware
func (siw *ServerInterfaceWrapper) CreateResource2(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "inline_argument" -------------
	var inlineArgument int

	if err := runtime.BindStyledParameter("simple", false, "inline_argument", chi.URLParam(r, "inline_argument"), &inlineArgument); err != nil {
		err = fmt.Errorf("invalid format for parameter inline_argument: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params CreateResource2Params

	// ------------- Optional query parameter "inline_query_argument" -------------

	if err := runtime.BindQueryParameter("form", true, false, "inline_query_argument", r.URL.Query(), &params.InlineQueryArgument); err != nil {
		err = fmt.Errorf("invalid format for parameter inline_query_argument: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateResource2(w, r, inlineArgument, params)
	})

	handler(w, r.WithContext(ctx))
}

// UpdateResource3 operation middleware
func (siw *ServerInterfaceWrapper) UpdateResource3(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "fallthrough" -------------
	var pFallthrough int

	if err := runtime.BindStyledParameter("simple", false, "fallthrough", chi.URLParam(r, "fallthrough"), &pFallthrough); err != nil {
		err = fmt.Errorf("invalid format for parameter fallthrough: %w", err)
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateResource3(w, r, pFallthrough)
	})

	handler(w, r.WithContext(ctx))
}

// GetResponseWithReference operation middleware
func (siw *ServerInterfaceWrapper) GetResponseWithReference(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetResponseWithReference(w, r)
	})

	handler(w, r.WithContext(ctx))
}

// GetWithTaggedMiddleware operation middleware
func (siw *ServerInterfaceWrapper) GetWithTaggedMiddleware(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetWithTaggedMiddleware(w, r)
	})

	// Operation specific middleware
	handler = siw.Middlewares["pathMiddleware"](handler).ServeHTTP

	handler(w, r.WithContext(ctx))
}

// PostWithTaggedMiddleware operation middleware
func (siw *ServerInterfaceWrapper) PostWithTaggedMiddleware(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostWithTaggedMiddleware(w, r)
	})

	// Operation specific middleware
	handler = siw.Middlewares["pathMiddleware"](handler).ServeHTTP
	handler = siw.Middlewares["operationMiddleware"](handler).ServeHTTP

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	error
}
type UnmarshalingParamError struct {
	error
}
type RequiredParamError struct {
	error
}
type RequiredHeaderError struct {
	error
}
type InvalidParamFormatError struct {
	error
}
type TooManyValuesForParamError struct {
	error
}

type ServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      map[string]func(http.Handler) http.Handler
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

type ServerOption func(*ServerOptions)

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface, opts ...ServerOption) http.Handler {
	options := &ServerOptions{
		BaseURL:     "/",
		BaseRouter:  chi.NewRouter(),
		Middlewares: make(map[string]func(http.Handler) http.Handler),
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
	}

	for _, f := range opts {
		f(options)
	}

	r := options.BaseRouter
	wrapper := ServerInterfaceWrapper{
		Handler:          si,
		Middlewares:      options.Middlewares,
		ErrorHandlerFunc: options.ErrorHandlerFunc,
	}

	middlewares := []string{"operationMiddleware", "pathMiddleware"}
	for _, m := range middlewares {
		if _, ok := wrapper.Middlewares[m]; !ok {
			panic("goapi-gen: could not find tagged middleware " + m)
		}
	}

	r.Route(options.BaseURL, func(r chi.Router) {
		r.Get("/every-type-optional", wrapper.GetEveryTypeOptional)
		r.Get("/get-simple", wrapper.GetSimple)
		r.Get("/get-with-args", wrapper.GetWithArgs)
		r.Get("/get-with-references/{global_argument}/{argument}", wrapper.GetWithReferences)
		r.Get("/get-with-type/{content_type}", wrapper.GetWithContentType)
		r.Get("/reserved-keyword", wrapper.GetReservedKeyword)
		r.Post("/resource/{argument}", wrapper.CreateResource)
		r.Post("/resource2/{inline_argument}", wrapper.CreateResource2)
		r.Put("/resource3/{fallthrough}", wrapper.UpdateResource3)
		r.Get("/response-with-reference", wrapper.GetResponseWithReference)
		r.Get("/with-tagged-middleware", wrapper.GetWithTaggedMiddleware)
		r.Post("/with-tagged-middleware", wrapper.PostWithTaggedMiddleware)

	})
	return r
}

func WithRouter(r chi.Router) ServerOption {
	return func(s *ServerOptions) {
		s.BaseRouter = r
	}
}

func WithServerBaseURL(url string) ServerOption {
	return func(s *ServerOptions) {
		s.BaseURL = url
	}
}

func WithMiddleware(key string, middleware func(http.Handler) http.Handler) ServerOption {
	return func(s *ServerOptions) {
		s.Middlewares[key] = middleware
	}
}

func WithMiddlewares(middlewares map[string]func(http.Handler) http.Handler) ServerOption {
	return func(s *ServerOptions) {
		s.Middlewares = middlewares
	}
}

func WithErrorHandler(handler func(w http.ResponseWriter, r *http.Request, err error)) ServerOption {
	return func(s *ServerOptions) {
		s.ErrorHandlerFunc = handler
	}
}
