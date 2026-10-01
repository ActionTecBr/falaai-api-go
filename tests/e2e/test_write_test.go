package e2e

import (
	"context"
	"fmt"
	"testing"

	falaai "github.com/actiontecbr/falaai-api-go"
)

const (
	wName  = "E2E Test Webhook"
	wURL   = "https://e2e-falaai.invalid/hook"
	aName  = "E2E Test Alert"
	aEmail = "e2e-test@falaai.invalid"
)

func sptr(s string) *string { return &s }
func bptr(b bool) *bool     { return &b }

func assertWebhookFull(t *testing.T, w falaai.WebhookItem, wid, name string, active bool) {
	t.Helper()
	if w.Id != wid {
		t.Fatalf("id %q != %q", w.Id, wid)
	}
	mustStr(t, "user_id", w.UserId)
	if w.Name != name {
		t.Fatalf("name %q != %q", w.Name, name)
	}
	if w.Url != wURL {
		t.Fatalf("url %q != %q", w.Url, wURL)
	}
	mustStr(t, "secret", w.Secret)
	if len(w.Events) != 2 {
		t.Fatalf("events len %d != 2", len(w.Events))
	}
	if w.Active != active {
		t.Fatalf("active %v != %v", w.Active, active)
	}
	mustStr(t, "created_at", w.CreatedAt)
	mustStr(t, "updated_at", w.UpdatedAt)
}

func assertAlertFull(t *testing.T, a falaai.EmailAlertItem, aid, name string, active bool) {
	t.Helper()
	if a.Id != aid {
		t.Fatalf("id %q != %q", a.Id, aid)
	}
	mustStr(t, "user_id", a.UserId)
	if a.Name != name {
		t.Fatalf("name %q != %q", a.Name, name)
	}
	if a.Email != aEmail {
		t.Fatalf("email %q != %q", a.Email, aEmail)
	}
	if len(a.Events) != 2 {
		t.Fatalf("events len %d != 2", len(a.Events))
	}
	if a.Active != active {
		t.Fatalf("active %v != %v", a.Active, active)
	}
	mustStr(t, "created_at", a.CreatedAt)
	mustStr(t, "updated_at", a.UpdatedAt)
}

func TestWebhooksCrud(t *testing.T) {
	c, _ := NewClient(BaseURL(), TestKey())
	ctx := context.Background()
	lst, _, _ := c.WebhooksAPI.ListWebhooksV1WebhooksGet(ctx).Page(1).Limit(100).Execute()
	if lst != nil {
		for _, w := range lst.Data {
			if w.Url == wURL {
				_, _, _ = c.WebhooksAPI.DeleteWebhookV1WebhooksWebhookIdDelete(ctx, w.Id).Execute()
			}
		}
	}

	body := falaai.CreateWebhookRequest{Name: wName, Url: wURL, Events: []falaai.WebhookEvent{falaai.WEBHOOKEVENT_CREDITS_LOW, falaai.WEBHOOKEVENT_PAYMENT_FAILED}}
	cr, httpResp, err := c.WebhooksAPI.CreateWebhookV1WebhooksPost(ctx).CreateWebhookRequest(body).Execute()
	if err != nil {
		t.Fatal(err)
	}
	status := httpResp.StatusCode
	Log("webhooks_create", "POST", "/v1/webhooks", body, cr, fmt.Sprintf("HTTP %d", status), status)
	if status != 200 {
		t.Fatalf("create status %d", status)
	}
	wid := cr.Id
	assertWebhookFull(t, *cr, wid, wName, true)

	ub := falaai.UpdateWebhookRequest{Name: sptr(wName + " (updated)"), Active: bptr(false)}
	ur, httpResp, err := c.WebhooksAPI.UpdateWebhookV1WebhooksWebhookIdPut(ctx, wid).UpdateWebhookRequest(ub).Execute()
	if err != nil {
		t.Fatal(err)
	}
	status = httpResp.StatusCode
	Log("webhooks_update", "PUT", "/v1/webhooks/"+wid, ub, ur, fmt.Sprintf("HTTP %d", status), status)
	if status != 200 {
		t.Fatalf("update status %d", status)
	}
	if ur.Message != "updated" {
		t.Fatalf("message %q != updated", ur.Message)
	}

	lst2, _, err := c.WebhooksAPI.ListWebhooksV1WebhooksGet(ctx).Page(1).Limit(100).Execute()
	if err != nil {
		t.Fatal(err)
	}
	var row *falaai.WebhookItem
	for i := range lst2.Data {
		if lst2.Data[i].Id == wid {
			row = &lst2.Data[i]
		}
	}
	if row == nil {
		t.Fatalf("webhook %s nao encontrado na lista", wid)
	}
	assertWebhookFull(t, *row, wid, wName+" (updated)", false)

	dr, httpResp, err := c.WebhooksAPI.DeleteWebhookV1WebhooksWebhookIdDelete(ctx, wid).Execute()
	if err != nil {
		t.Fatal(err)
	}
	status = httpResp.StatusCode
	Log("webhooks_delete", "DELETE", "/v1/webhooks/"+wid, nil, dr, fmt.Sprintf("HTTP %d", status), status)
	if status != 200 {
		t.Fatalf("delete status %d", status)
	}
	if dr.Message != "deleted" {
		t.Fatalf("message %q != deleted", dr.Message)
	}

	lst3, _, err := c.WebhooksAPI.ListWebhooksV1WebhooksGet(ctx).Page(1).Limit(100).Execute()
	if err != nil {
		t.Fatal(err)
	}
	for _, w := range lst3.Data {
		if w.Id == wid {
			t.Fatalf("webhook %s ainda existe apos delete", wid)
		}
	}
}

func TestEmailAlertsCrud(t *testing.T) {
	c, _ := NewClient(BaseURL(), TestKey())
	ctx := context.Background()
	lst, _, _ := c.EmailAlertsAPI.ListEmailAlertsV1EmailAlertsGet(ctx).Page(1).Limit(100).Execute()
	if lst != nil {
		for _, a := range lst.Data {
			if a.Email == aEmail {
				_, _, _ = c.EmailAlertsAPI.DeleteEmailAlertV1EmailAlertsAlertIdDelete(ctx, a.Id).Execute()
			}
		}
	}

	body := falaai.CreateEmailAlertRequest{Name: aName, Email: aEmail, Events: []falaai.EmailEvent{falaai.EMAILEVENT_CREDITS_LOW, falaai.EMAILEVENT_PAYMENT_FAILED}}
	cr, httpResp, err := c.EmailAlertsAPI.CreateEmailAlertV1EmailAlertsPost(ctx).CreateEmailAlertRequest(body).Execute()
	if err != nil {
		t.Fatal(err)
	}
	status := httpResp.StatusCode
	Log("email_alerts_create", "POST", "/v1/email-alerts", body, cr, fmt.Sprintf("HTTP %d", status), status)
	if status != 200 {
		t.Fatalf("create status %d", status)
	}
	aid := cr.Id
	assertAlertFull(t, *cr, aid, aName, true)

	ub := falaai.UpdateEmailAlertRequest{Name: sptr(aName + " (updated)"), Active: bptr(false)}
	ur, httpResp, err := c.EmailAlertsAPI.UpdateEmailAlertV1EmailAlertsAlertIdPut(ctx, aid).UpdateEmailAlertRequest(ub).Execute()
	if err != nil {
		t.Fatal(err)
	}
	status = httpResp.StatusCode
	Log("email_alerts_update", "PUT", "/v1/email-alerts/"+aid, ub, ur, fmt.Sprintf("HTTP %d", status), status)
	if status != 200 {
		t.Fatalf("update status %d", status)
	}
	if ur.Message != "updated" {
		t.Fatalf("message %q != updated", ur.Message)
	}

	lst2, _, err := c.EmailAlertsAPI.ListEmailAlertsV1EmailAlertsGet(ctx).Page(1).Limit(100).Execute()
	if err != nil {
		t.Fatal(err)
	}
	var row *falaai.EmailAlertItem
	for i := range lst2.Data {
		if lst2.Data[i].Id == aid {
			row = &lst2.Data[i]
		}
	}
	if row == nil {
		t.Fatalf("alert %s nao encontrado na lista", aid)
	}
	assertAlertFull(t, *row, aid, aName+" (updated)", false)

	dr, httpResp, err := c.EmailAlertsAPI.DeleteEmailAlertV1EmailAlertsAlertIdDelete(ctx, aid).Execute()
	if err != nil {
		t.Fatal(err)
	}
	status = httpResp.StatusCode
	Log("email_alerts_delete", "DELETE", "/v1/email-alerts/"+aid, nil, dr, fmt.Sprintf("HTTP %d", status), status)
	if status != 200 {
		t.Fatalf("delete status %d", status)
	}
	if dr.Message != "deleted" {
		t.Fatalf("message %q != deleted", dr.Message)
	}

	lst3, _, err := c.EmailAlertsAPI.ListEmailAlertsV1EmailAlertsGet(ctx).Page(1).Limit(100).Execute()
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range lst3.Data {
		if a.Id == aid {
			t.Fatalf("alert %s ainda existe apos delete", aid)
		}
	}
}