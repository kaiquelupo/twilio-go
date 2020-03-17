package twilio

import (
	"context"
	"net/url"
)

type ServerlessServiceCreator struct {
	client       *Client
	workspaceSid string
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
	err := r.client.GetResource(ctx, "Services", sid, service)
	return worker, err
}

func (r *ServerlessServiceCreator) Create(ctx context.Context, data url.Values) (*Service, error) {
	service := new(Service)
	err := r.client.CreateResource(ctx, "Services", data, service)
	return worker, err
}


func (r *ServerlessServiceCreator) Delete(ctx context.Context, sid string) error {
	return r.client.DeleteResource(ctx, "Services", sid)
}

func (r *ServerlessServiceCreator) Update(ctx context.Context, sid string, data url.Values) (*Service, error) {
	service := new(Service)
	err := r.client.UpdateResource(ctx, "Services", sid, data, service)
	return worker, err
}