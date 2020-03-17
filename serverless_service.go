package twilio

import (
	"context"
	"net/url"
)

type ServerlessServiceCreator struct {
	client       *Client
}

type Service struct {
	Sid          string `json:"sid"`
	AccountSid   string `json:"account_sid"`
	FriendlyName string `json:"friendly_name"`
	UniqueName	 string `json:"unique_name"`
	IncludeCredentials    bool       `json:"include_credentials"`
	UiEditable	 bool	`json:"ui_editable"`
	DateCreated  TwilioTime `json:"date_created"`
	DateUpdated  TwilioTime `json:"date_updated"`
	URL          string     `json:"url"`
}

func (r *ServerlessServiceCreator) Get(ctx context.Context, sid string) (*Service, error) {
	service := new(Service)
	err := r.client.Serverless.GetResource(ctx, "Services", sid, service)
	return service, err
}

func (r *ServerlessServiceCreator) Create(ctx context.Context, data url.Values) (*Service, error) {
	service := new(Service)
	err := r.client.Serverless.CreateResource(ctx, "Services", data, service)
	return service, err
}


func (r *ServerlessServiceCreator) Delete(ctx context.Context, sid string) error {
	return r.client.Serverless.DeleteResource(ctx, "Services", sid)
}

func (r *ServerlessServiceCreator) Update(ctx context.Context, sid string, data url.Values) (*Service, error) {
	service := new(Service)
	err := r.client.Serverless.UpdateResource(ctx, "Services", sid, data, service)
	return service, err
}