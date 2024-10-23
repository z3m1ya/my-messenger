// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostChatsChatIDMessagesHandlerFunc turns a function with the right signature into a post chats chat ID messages handler
type PostChatsChatIDMessagesHandlerFunc func(PostChatsChatIDMessagesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostChatsChatIDMessagesHandlerFunc) Handle(params PostChatsChatIDMessagesParams) middleware.Responder {
	return fn(params)
}

// PostChatsChatIDMessagesHandler interface for that can handle valid post chats chat ID messages params
type PostChatsChatIDMessagesHandler interface {
	Handle(PostChatsChatIDMessagesParams) middleware.Responder
}

// NewPostChatsChatIDMessages creates a new http.Handler for the post chats chat ID messages operation
func NewPostChatsChatIDMessages(ctx *middleware.Context, handler PostChatsChatIDMessagesHandler) *PostChatsChatIDMessages {
	return &PostChatsChatIDMessages{Context: ctx, Handler: handler}
}

/*
	PostChatsChatIDMessages swagger:route POST /chats/{chatId}/messages postChatsChatIdMessages

Add a message to a chat
*/
type PostChatsChatIDMessages struct {
	Context *middleware.Context
	Handler PostChatsChatIDMessagesHandler
}

func (o *PostChatsChatIDMessages) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostChatsChatIDMessagesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostChatsChatIDMessagesBody post chats chat ID messages body
//
// swagger:model PostChatsChatIDMessagesBody
type PostChatsChatIDMessagesBody struct {

	// message
	Message string `json:"message,omitempty"`

	// sender Id
	SenderID string `json:"senderId,omitempty"`
}

// Validate validates this post chats chat ID messages body
func (o *PostChatsChatIDMessagesBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post chats chat ID messages body based on context it is used
func (o *PostChatsChatIDMessagesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostChatsChatIDMessagesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostChatsChatIDMessagesBody) UnmarshalBinary(b []byte) error {
	var res PostChatsChatIDMessagesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
