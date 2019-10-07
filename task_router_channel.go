package twilio

import (
	"context"
	"net/url"
)

type TaskChannelService struct {
	client       *Client
	workspaceSid string
}

type TaskChannel struct {
	AccountSid              string            `json:"account_sid"`
	DateCreated             TwilioTime        `json:"date_created"`
	DateUpdated             TwilioTime        `json:"date_updated"`
	FriendlyName            string            `json:"friendly_name"`
	Sid                     string            `json:"sid"`
	UniqueName              string            `json: "unique_name"`
	WorkspaceSid            string            `json:"workspace_sid"`
	ChannelOptimizedRouting bool              `json:"channel_optimized_routing"`
	URL                     string            `json:"url"`
	Links                   map[string]string `json:"links"`
}

func taskChannelUrlPath(workspaceSid string) string {
	urlPath := "Workspaces/" + workspaceSid + "/" + "TaskChannels"
	return urlPath
}

func (svc *TaskChannelService) Get(ctx context.Context, sid string) (*TaskChannel, error) {
	taskChannel := new(TaskChannel)
	err := svc.client.GetResource(ctx, taskChannelUrlPath(svc.workspaceSid), sid, taskChannel)
	return taskChannel, err
}

func (svc *TaskChannelService) Create(ctx context.Context, data url.Values) (*TaskChannel, error) {
	taskChannel := new(TaskChannel)
	err := svc.client.CreateResource(ctx, taskChannelUrlPath(svc.workspaceSid), data, taskChannel)
	return taskChannel, err
}

func (svc *TaskChannelService) Update(ctx context.Context, sid string, data url.Values) (*TaskChannel, error) {
	taskChannel := new(TaskChannel)
	err := svc.client.UpdateResource(ctx, taskChannelUrlPath(svc.workspaceSid), sid, data, taskChannel)
	return taskChannel, err
}

func (svc *TaskChannelService) Delete(ctx context.Context, sid string) error {
	return svc.client.DeleteResource(ctx, taskChannelUrlPath(svc.workspaceSid), sid)
}
