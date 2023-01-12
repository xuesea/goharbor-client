package quota

import (
	"context"
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/go-openapi/runtime"

	v2client "github.com/xuesea/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/xuesea/goharbor-client/v5/apiv2/internal/api/client/quota"
	"github.com/xuesea/goharbor-client/v5/apiv2/model"
	"github.com/xuesea/goharbor-client/v5/apiv2/pkg/config"
	"github.com/xuesea/goharbor-client/v5/apiv2/pkg/errors"
)

// RESTClient is a subclient for handling project related actions.
type RESTClient struct {
	// Options contains optional configuration when making API calls.
	Options *config.Options

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(v2Client *v2client.Harbor, opts *config.Options, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Options:  opts,
		V2Client: v2Client,
		AuthInfo: authInfo,
	}
}

type Client interface {
	ListQuotas(ctx context.Context, referenceType, referenceID *string) ([]*model.Quota, error)
	GetQuotaByProjectID(ctx context.Context, projectID int64) (*model.Quota, error)
	UpdateStorageQuotaByProjectID(ctx context.Context, projectID int64, storageLimit int64) error
}

func (c *RESTClient) ListQuotas(ctx context.Context, referenceType, referenceID *string) ([]*model.Quota, error) {
	var quotas []*model.Quota
	page := c.Options.Page

	params := &quota.ListQuotasParams{
		Page:        &page,
		PageSize:    &c.Options.PageSize,
		Reference:   referenceType,
		ReferenceID: referenceID,
		Sort:        &c.Options.Sort,
		Context:     ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	for {
		resp, err := c.V2Client.Quota.ListQuotas(params, c.AuthInfo)
		if err != nil {
			return nil, handleSwaggerQuotaErrors(err)
		}

		if len(resp.Payload) == 0 {
			break
		}

		totalCount := resp.XTotalCount

		quotas = append(quotas, resp.Payload...)

		if int64(len(quotas)) >= totalCount {
			break
		}

		page++
	}

	return quotas, nil
}

// GetQuotaByProjectID returns a quota object containing all configured quotas for a project.
func (c *RESTClient) GetQuotaByProjectID(ctx context.Context, projectID int64) (*model.Quota, error) {
	projectIDStr := strconv.Itoa(int(projectID))

	quotas, err := c.ListQuotas(ctx, nil, &projectIDStr)
	if err != nil {
		return nil, handleSwaggerQuotaErrors(err)
	}

	// Assert that quota.Ref implements a map[string]interface{} type, as it holds json data.
	for _, q := range quotas {
		if values, ok := q.Ref.(map[string]interface{}); ok {
			if reflect.DeepEqual(values["id"], json.Number(projectIDStr)) {
				return q, nil
			}
		}
	}

	return nil, &errors.ErrQuotaRefNotFound{}
}

// UpdateStorageQuotaByProjectID updates the storageLimit quota of a project.
// A storageLimit value smaller than '0' will implicitly be set to '-1', equalling the 'unlimited' setting.
func (c *RESTClient) UpdateStorageQuotaByProjectID(ctx context.Context, projectID int64, storageLimit int64) error {
	if storageLimit <= 0 {
		storageLimit = -1
	}

	params := &quota.UpdateQuotaParams{
		Hard: &model.QuotaUpdateReq{
			Hard: model.ResourceList{
				"storage": storageLimit,
			},
		},
		ID:      projectID,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Quota.UpdateQuota(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerQuotaErrors(err)
	}

	return nil
}
