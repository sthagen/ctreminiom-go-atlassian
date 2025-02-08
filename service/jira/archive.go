package jira

import (
	"context"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type ArchiveService interface {
	Preserve(ctx context.Context, issueIdsOrKeys []string) (result *models.IssueArchivalSyncResponseScheme, response *models.ResponseScheme, err error)
	PreserveByJQL(ctx context.Context, jql string) (taskID string, response *models.ResponseScheme, err error)
	Restore(ctx context.Context, issueIdsOrKeys []string) (result *models.IssueArchivalSyncResponseScheme, response *models.ResponseScheme, err error)
	Export(ctx context.Context, payload *models.IssueArchivalExportPayloadScheme) (taskID string, response *models.ResponseScheme, err error)
}
