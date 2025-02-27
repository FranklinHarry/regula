// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fugue/regula/pkg/swagger/models"
)

// DeleteNotificationReader is a Reader for the DeleteNotification structure.
type DeleteNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNotificationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteNotificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteNotificationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteNotificationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteNotificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteNotificationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteNotificationNoContent creates a DeleteNotificationNoContent with default headers values
func NewDeleteNotificationNoContent() *DeleteNotificationNoContent {
	return &DeleteNotificationNoContent{}
}

/*DeleteNotificationNoContent handles this case with default header values.

Notification deleted.
*/
type DeleteNotificationNoContent struct {
}

func (o *DeleteNotificationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /notifications/{notification_id}][%d] deleteNotificationNoContent ", 204)
}

func (o *DeleteNotificationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNotificationBadRequest creates a DeleteNotificationBadRequest with default headers values
func NewDeleteNotificationBadRequest() *DeleteNotificationBadRequest {
	return &DeleteNotificationBadRequest{}
}

/*DeleteNotificationBadRequest handles this case with default header values.

BadRequestError
*/
type DeleteNotificationBadRequest struct {
	Payload *models.BadRequestError
}

func (o *DeleteNotificationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /notifications/{notification_id}][%d] deleteNotificationBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteNotificationBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *DeleteNotificationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNotificationUnauthorized creates a DeleteNotificationUnauthorized with default headers values
func NewDeleteNotificationUnauthorized() *DeleteNotificationUnauthorized {
	return &DeleteNotificationUnauthorized{}
}

/*DeleteNotificationUnauthorized handles this case with default header values.

AuthenticationError
*/
type DeleteNotificationUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *DeleteNotificationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /notifications/{notification_id}][%d] deleteNotificationUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteNotificationUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *DeleteNotificationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNotificationForbidden creates a DeleteNotificationForbidden with default headers values
func NewDeleteNotificationForbidden() *DeleteNotificationForbidden {
	return &DeleteNotificationForbidden{}
}

/*DeleteNotificationForbidden handles this case with default header values.

AuthorizationError
*/
type DeleteNotificationForbidden struct {
	Payload *models.AuthorizationError
}

func (o *DeleteNotificationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /notifications/{notification_id}][%d] deleteNotificationForbidden  %+v", 403, o.Payload)
}

func (o *DeleteNotificationForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *DeleteNotificationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNotificationNotFound creates a DeleteNotificationNotFound with default headers values
func NewDeleteNotificationNotFound() *DeleteNotificationNotFound {
	return &DeleteNotificationNotFound{}
}

/*DeleteNotificationNotFound handles this case with default header values.

NotFoundError
*/
type DeleteNotificationNotFound struct {
	Payload *models.NotFoundError
}

func (o *DeleteNotificationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /notifications/{notification_id}][%d] deleteNotificationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNotificationNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *DeleteNotificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNotificationInternalServerError creates a DeleteNotificationInternalServerError with default headers values
func NewDeleteNotificationInternalServerError() *DeleteNotificationInternalServerError {
	return &DeleteNotificationInternalServerError{}
}

/*DeleteNotificationInternalServerError handles this case with default header values.

InternalServerError
*/
type DeleteNotificationInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *DeleteNotificationInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /notifications/{notification_id}][%d] deleteNotificationInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteNotificationInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *DeleteNotificationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
