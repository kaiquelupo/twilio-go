package twilio

import (
	"context"
	"net/url"
	"os"
	"testing"
	"time"
)

func TestWorkspaceCreation(t *testing.T) {
	c := NewClient(os.Getenv("TWILIO_ACCOUNT_SID"), os.Getenv("TWILIO_AUTH_TOKEN"), nil)

	friendlyName := "TestFriendly " + time.Now().String()
	workspace := createWorkspaceForTestingAndCheckErrors(t, c, friendlyName)
	defer c.WorkspaceCreator.Delete(context.TODO(), workspace.Sid)
}

func TestWorkspaceRead(t *testing.T) {
	c := NewClient(os.Getenv("TWILIO_ACCOUNT_SID"), os.Getenv("TWILIO_AUTH_TOKEN"), nil)

	// First Create
	friendlyName := "TestFriendly " + time.Now().String()
	workspace := createWorkspaceForTestingAndCheckErrors(t, c, friendlyName)
	defer c.WorkspaceCreator.Delete(context.TODO(), workspace.Sid)

	// Next Read
	readWorkspace, readErr := c.WorkspaceCreator.Get(context.TODO(), workspace.Sid)

	if readErr != nil {
		t.Fatal("Error occured read:", readErr)
	}
	if readWorkspace.Sid != workspace.Sid {
		t.Error("Read Workspace SID was incorrect, expected :", workspace.Sid, ", got instead:", readWorkspace.Sid)
	}
	if readWorkspace.FriendlyName != friendlyName {
		t.Error("Read Friendly Name was incorrect, expected :", friendlyName, ", got instead:", readWorkspace.FriendlyName)
	}
}
func TestWorkspaceUpdate(t *testing.T) {
	c := NewClient(os.Getenv("TWILIO_ACCOUNT_SID"), os.Getenv("TWILIO_AUTH_TOKEN"), nil)

	// First Create
	friendlyName := "TestFriendly " + time.Now().String()
	workspace := createWorkspaceForTestingAndCheckErrors(t, c, friendlyName)
	defer c.WorkspaceCreator.Delete(context.TODO(), workspace.Sid)

	// Next update
	updatedFriendlyName := "UpdateFriendly " + time.Now().String()
	updateData := url.Values{"FriendlyName": []string{updatedFriendlyName}}
	updatedWorkspace, updateErr := c.WorkspaceCreator.Update(context.TODO(), workspace.Sid, updateData)

	if updateErr != nil {
		t.Fatal("Error occured during update:", updateErr)
	}

	if updatedWorkspace.FriendlyName != updatedFriendlyName {
		t.Error("Updated Friendly name was incorrect, expected :", updatedFriendlyName, ", got instead:", updatedWorkspace.FriendlyName)
	}

	// Read back update to confirm
	readWorkspace, readErr := c.WorkspaceCreator.Get(context.TODO(), workspace.Sid)

	if readErr != nil {
		t.Fatal("Error occured read:", readErr)
	}
	if readWorkspace.Sid != workspace.Sid {
		t.Error("Read Workspace SID was incorrect, expected :", workspace.Sid, ", got instead:", readWorkspace.Sid)
	}
	if readWorkspace.FriendlyName != updatedWorkspace.FriendlyName {
		t.Error("Read Friendly Name was incorrect, expected :", updatedWorkspace.FriendlyName, ", got instead:", readWorkspace.FriendlyName)
	}
}

func TestWorkspaceDelete(t *testing.T) {
	c := NewClient(os.Getenv("TWILIO_ACCOUNT_SID"), os.Getenv("TWILIO_AUTH_TOKEN"), nil)
	friendlyName := "TestFriendly " + time.Now().String()
	workspace := createWorkspaceForTestingAndCheckErrors(t, c, friendlyName)

	deleteErr := c.WorkspaceCreator.Delete(context.TODO(), workspace.Sid)

	if deleteErr != nil {
		t.Fatal("Error occured during delete:", deleteErr)
	}
}

func createWorkspaceForTestingAndCheckErrors(t *testing.T, c *Client, friendlyName string) *Workspace {
	data := url.Values{"FriendlyName": []string{friendlyName}}

	workspace, err := c.WorkspaceCreator.Create(context.TODO(), data)
	if err != nil {
		t.Fatal("Error occurred:", err.Error())
	}
	if workspace.FriendlyName != friendlyName {
		t.Fatal("Friendly name was incorrect, expected :", friendlyName, ", got instead:", workspace.FriendlyName)
	}
	return workspace
}
