package resources

import (
	"context"
	"testing"

	interfaces2 "github.com/lyft/flyteadmin/pkg/manager/interfaces"
	"github.com/lyft/flyteadmin/pkg/repositories/interfaces"

	"github.com/golang/protobuf/proto"
	"github.com/lyft/flyteadmin/pkg/manager/impl/testutils"
	"github.com/lyft/flyteadmin/pkg/repositories/mocks"
	"github.com/lyft/flyteadmin/pkg/repositories/models"
	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"
	"github.com/stretchr/testify/assert"
)

const project = "project"
const domain = "domain"
const workflow = "workflow"

func TestUpdateWorkflowAttributes(t *testing.T) {
	request := admin.WorkflowAttributesUpdateRequest{
		Attributes: &admin.WorkflowAttributes{
			Project:            "project",
			Domain:             "domain",
			Workflow:           "workflow",
			MatchingAttributes: testutils.ExecutionQueueAttributes,
		},
	}
	db := mocks.NewMockRepository()
	expectedSerializedAttrs, _ := proto.Marshal(testutils.ExecutionQueueAttributes)
	var createOrUpdateCalled bool
	db.ResourceRepo().(*mocks.MockResourceRepo).CreateOrUpdateFunction = func(
		ctx context.Context, input models.Resource) error {
		assert.Equal(t, "project", input.Project)
		assert.Equal(t, "domain", input.Domain)
		assert.Equal(t, "workflow", input.Workflow)
		assert.Equal(t, admin.MatchableResource_EXECUTION_QUEUE.String(), input.ResourceType)
		assert.EqualValues(t, expectedSerializedAttrs, input.Attributes)
		createOrUpdateCalled = true
		return nil
	}
	manager := NewResourceManager(db)
	_, err := manager.UpdateWorkflowAttributes(context.Background(), request)
	assert.Nil(t, err)
	assert.True(t, createOrUpdateCalled)
}

func TestGetWorkflowAttributes(t *testing.T) {
	request := admin.WorkflowAttributesGetRequest{
		Project:      project,
		Domain:       domain,
		Workflow:     workflow,
		ResourceType: admin.MatchableResource_EXECUTION_QUEUE,
	}
	db := mocks.NewMockRepository()
	db.ResourceRepo().(*mocks.MockResourceRepo).GetFunction = func(
		ctx context.Context, ID interfaces.ResourceID) (models.Resource, error) {
		assert.Equal(t, "project", ID.Project)
		assert.Equal(t, "domain", ID.Domain)
		assert.Equal(t, "workflow", ID.Workflow)
		assert.Equal(t, admin.MatchableResource_EXECUTION_QUEUE.String(), ID.ResourceType)
		expectedSerializedAttrs, _ := proto.Marshal(testutils.ExecutionQueueAttributes)
		return models.Resource{
			Project:      project,
			Domain:       domain,
			Workflow:     workflow,
			ResourceType: "resource",
			Attributes:   expectedSerializedAttrs,
		}, nil
	}
	manager := NewResourceManager(db)
	response, err := manager.GetWorkflowAttributes(context.Background(), request)
	assert.Nil(t, err)
	assert.True(t, proto.Equal(&admin.WorkflowAttributesGetResponse{
		Attributes: &admin.WorkflowAttributes{
			Project:            "project",
			Domain:             "domain",
			Workflow:           "workflow",
			MatchingAttributes: testutils.ExecutionQueueAttributes,
		},
	}, response))
}

func TestDeleteWorkflowAttributes(t *testing.T) {
	request := admin.WorkflowAttributesDeleteRequest{
		Project:      "project",
		Domain:       "domain",
		Workflow:     "workflow",
		ResourceType: admin.MatchableResource_EXECUTION_QUEUE,
	}
	db := mocks.NewMockRepository()
	db.ResourceRepo().(*mocks.MockResourceRepo).DeleteFunction = func(
		ctx context.Context, ID interfaces.ResourceID) error {
		assert.Equal(t, "project", project)
		assert.Equal(t, "domain", domain)
		assert.Equal(t, "workflow", workflow)
		assert.Equal(t, admin.MatchableResource_EXECUTION_QUEUE.String(), ID.ResourceType)
		return nil
	}
	manager := NewResourceManager(db)
	_, err := manager.DeleteWorkflowAttributes(context.Background(), request)
	assert.Nil(t, err)
}

func TestUpdateProjectDomainAttributes(t *testing.T) {
	request := admin.ProjectDomainAttributesUpdateRequest{
		Attributes: &admin.ProjectDomainAttributes{
			Project:            "project",
			Domain:             "domain",
			MatchingAttributes: testutils.ExecutionQueueAttributes,
		},
	}
	db := mocks.NewMockRepository()
	expectedSerializedAttrs, _ := proto.Marshal(testutils.ExecutionQueueAttributes)
	var createOrUpdateCalled bool
	db.ResourceRepo().(*mocks.MockResourceRepo).CreateOrUpdateFunction = func(
		ctx context.Context, input models.Resource) error {
		assert.Equal(t, "project", input.Project)
		assert.Equal(t, "domain", input.Domain)
		assert.Equal(t, "", input.Workflow)
		assert.Equal(t, admin.MatchableResource_EXECUTION_QUEUE.String(), input.ResourceType)
		assert.EqualValues(t, expectedSerializedAttrs, input.Attributes)
		createOrUpdateCalled = true
		return nil
	}
	manager := NewResourceManager(db)
	_, err := manager.UpdateProjectDomainAttributes(context.Background(), request)
	assert.Nil(t, err)
	assert.True(t, createOrUpdateCalled)
}

func TestGetProjectDomainAttributes(t *testing.T) {
	request := admin.ProjectDomainAttributesGetRequest{
		Project:      "project",
		Domain:       "domain",
		ResourceType: admin.MatchableResource_EXECUTION_QUEUE,
	}
	db := mocks.NewMockRepository()
	db.ResourceRepo().(*mocks.MockResourceRepo).GetFunction = func(
		ctx context.Context, ID interfaces.ResourceID) (models.Resource, error) {
		assert.Equal(t, "project", ID.Project)
		assert.Equal(t, "domain", ID.Domain)
		assert.Equal(t, "", ID.Workflow)
		assert.Equal(t, admin.MatchableResource_EXECUTION_QUEUE.String(), ID.ResourceType)
		expectedSerializedAttrs, _ := proto.Marshal(testutils.ExecutionQueueAttributes)
		return models.Resource{
			Project:      project,
			Domain:       domain,
			ResourceType: "resource",
			Attributes:   expectedSerializedAttrs,
		}, nil
	}
	manager := NewResourceManager(db)
	response, err := manager.GetProjectDomainAttributes(context.Background(), request)
	assert.Nil(t, err)
	assert.True(t, proto.Equal(&admin.ProjectDomainAttributesGetResponse{
		Attributes: &admin.ProjectDomainAttributes{
			Project:            "project",
			Domain:             "domain",
			MatchingAttributes: testutils.ExecutionQueueAttributes,
		},
	}, response))
}

func TestDeleteProjectDomainAttributes(t *testing.T) {
	request := admin.ProjectDomainAttributesDeleteRequest{
		Project:      "project",
		Domain:       "domain",
		ResourceType: admin.MatchableResource_EXECUTION_QUEUE,
	}
	db := mocks.NewMockRepository()
	db.ResourceRepo().(*mocks.MockResourceRepo).DeleteFunction = func(
		ctx context.Context, ID interfaces.ResourceID) error {
		assert.Equal(t, "project", ID.Project)
		assert.Equal(t, "domain", ID.Domain)
		assert.Equal(t, admin.MatchableResource_EXECUTION_QUEUE.String(), ID.ResourceType)
		return nil
	}
	manager := NewResourceManager(db)
	_, err := manager.DeleteProjectDomainAttributes(context.Background(), request)
	assert.Nil(t, err)
}

func TestGetResource(t *testing.T) {
	request := interfaces2.ResourceRequest{
		Project:      "project",
		Domain:       "domain",
		Workflow:     "workflow",
		LaunchPlan:   "launch_plan",
		ResourceType: admin.MatchableResource_EXECUTION_QUEUE,
	}
	db := mocks.NewMockRepository()
	db.ResourceRepo().(*mocks.MockResourceRepo).GetFunction = func(
		ctx context.Context, ID interfaces.ResourceID) (models.Resource, error) {
		assert.Equal(t, "project", ID.Project)
		assert.Equal(t, "domain", ID.Domain)
		assert.Equal(t, "workflow", ID.Workflow)
		assert.Equal(t, "launch_plan", ID.LaunchPlan)
		assert.Equal(t, admin.MatchableResource_EXECUTION_QUEUE.String(), ID.ResourceType)
		expectedSerializedAttrs, _ := proto.Marshal(testutils.ExecutionQueueAttributes)
		return models.Resource{
			Project:      ID.Project,
			Domain:       ID.Domain,
			Workflow:     ID.Workflow,
			LaunchPlan:   ID.LaunchPlan,
			ResourceType: ID.ResourceType,
			Attributes:   expectedSerializedAttrs,
		}, nil
	}
	manager := NewResourceManager(db)
	response, err := manager.GetResource(context.Background(), request)
	assert.Nil(t, err)
	assert.Equal(t, request.Project, response.Project)
	assert.Equal(t, request.Domain, response.Domain)
	assert.Equal(t, request.Workflow, response.Workflow)
	assert.Equal(t, request.LaunchPlan, response.LaunchPlan)
	assert.Equal(t, request.ResourceType.String(), response.ResourceType)
	assert.True(t, proto.Equal(response.Attributes, testutils.ExecutionQueueAttributes))
}