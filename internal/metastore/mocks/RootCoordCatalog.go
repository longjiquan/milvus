// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	milvuspb "github.com/milvus-io/milvus-proto/go-api/milvuspb"
	metastore "github.com/milvus-io/milvus/internal/metastore"

	mock "github.com/stretchr/testify/mock"

	model "github.com/milvus-io/milvus/internal/metastore/model"
)

// RootCoordCatalog is an autogenerated mock type for the RootCoordCatalog type
type RootCoordCatalog struct {
	mock.Mock
}

// AlterAlias provides a mock function with given fields: ctx, alias, ts
func (_m *RootCoordCatalog) AlterAlias(ctx context.Context, alias *model.Alias, ts uint64) error {
	ret := _m.Called(ctx, alias, ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Alias, uint64) error); ok {
		r0 = rf(ctx, alias, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AlterCollection provides a mock function with given fields: ctx, oldColl, newColl, alterType, ts
func (_m *RootCoordCatalog) AlterCollection(ctx context.Context, oldColl *model.Collection, newColl *model.Collection, alterType metastore.AlterType, ts uint64) error {
	ret := _m.Called(ctx, oldColl, newColl, alterType, ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Collection, *model.Collection, metastore.AlterType, uint64) error); ok {
		r0 = rf(ctx, oldColl, newColl, alterType, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AlterCredential provides a mock function with given fields: ctx, credential
func (_m *RootCoordCatalog) AlterCredential(ctx context.Context, credential *model.Credential) error {
	ret := _m.Called(ctx, credential)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Credential) error); ok {
		r0 = rf(ctx, credential)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AlterGrant provides a mock function with given fields: ctx, tenant, entity, operateType
func (_m *RootCoordCatalog) AlterGrant(ctx context.Context, tenant string, entity *milvuspb.GrantEntity, operateType milvuspb.OperatePrivilegeType) error {
	ret := _m.Called(ctx, tenant, entity, operateType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *milvuspb.GrantEntity, milvuspb.OperatePrivilegeType) error); ok {
		r0 = rf(ctx, tenant, entity, operateType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AlterPartition provides a mock function with given fields: ctx, dbName, oldPart, newPart, alterType, ts
func (_m *RootCoordCatalog) AlterPartition(ctx context.Context, dbName string, oldPart *model.Partition, newPart *model.Partition, alterType metastore.AlterType, ts uint64) error {
	ret := _m.Called(ctx, dbName, oldPart, newPart, alterType, ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.Partition, *model.Partition, metastore.AlterType, uint64) error); ok {
		r0 = rf(ctx, dbName, oldPart, newPart, alterType, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AlterUserRole provides a mock function with given fields: ctx, tenant, userEntity, roleEntity, operateType
func (_m *RootCoordCatalog) AlterUserRole(ctx context.Context, tenant string, userEntity *milvuspb.UserEntity, roleEntity *milvuspb.RoleEntity, operateType milvuspb.OperateUserRoleType) error {
	ret := _m.Called(ctx, tenant, userEntity, roleEntity, operateType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *milvuspb.UserEntity, *milvuspb.RoleEntity, milvuspb.OperateUserRoleType) error); ok {
		r0 = rf(ctx, tenant, userEntity, roleEntity, operateType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *RootCoordCatalog) Close() {
	_m.Called()
}

// CollectionExists provides a mock function with given fields: ctx, dbName, collectionID, ts
func (_m *RootCoordCatalog) CollectionExists(ctx context.Context, dbName string, collectionID int64, ts uint64) bool {
	ret := _m.Called(ctx, dbName, collectionID, ts)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, uint64) bool); ok {
		r0 = rf(ctx, dbName, collectionID, ts)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CreateAlias provides a mock function with given fields: ctx, alias, ts
func (_m *RootCoordCatalog) CreateAlias(ctx context.Context, alias *model.Alias, ts uint64) error {
	ret := _m.Called(ctx, alias, ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Alias, uint64) error); ok {
		r0 = rf(ctx, alias, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateCollection provides a mock function with given fields: ctx, collectionInfo, ts
func (_m *RootCoordCatalog) CreateCollection(ctx context.Context, collectionInfo *model.Collection, ts uint64) error {
	ret := _m.Called(ctx, collectionInfo, ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Collection, uint64) error); ok {
		r0 = rf(ctx, collectionInfo, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateCredential provides a mock function with given fields: ctx, credential
func (_m *RootCoordCatalog) CreateCredential(ctx context.Context, credential *model.Credential) error {
	ret := _m.Called(ctx, credential)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Credential) error); ok {
		r0 = rf(ctx, credential)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDatabase provides a mock function with given fields: ctx, dbName, ts
func (_m *RootCoordCatalog) CreateDatabase(ctx context.Context, dbName string, ts uint64) error {
	ret := _m.Called(ctx, dbName, ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64) error); ok {
		r0 = rf(ctx, dbName, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePartition provides a mock function with given fields: ctx, dbName, partition, ts
func (_m *RootCoordCatalog) CreatePartition(ctx context.Context, dbName string, partition *model.Partition, ts uint64) error {
	ret := _m.Called(ctx, dbName, partition, ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.Partition, uint64) error); ok {
		r0 = rf(ctx, dbName, partition, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRole provides a mock function with given fields: ctx, tenant, entity
func (_m *RootCoordCatalog) CreateRole(ctx context.Context, tenant string, entity *milvuspb.RoleEntity) error {
	ret := _m.Called(ctx, tenant, entity)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *milvuspb.RoleEntity) error); ok {
		r0 = rf(ctx, tenant, entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteGrant provides a mock function with given fields: ctx, tenant, role
func (_m *RootCoordCatalog) DeleteGrant(ctx context.Context, tenant string, role *milvuspb.RoleEntity) error {
	ret := _m.Called(ctx, tenant, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *milvuspb.RoleEntity) error); ok {
		r0 = rf(ctx, tenant, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DropAlias provides a mock function with given fields: ctx, dbName, alias, ts
func (_m *RootCoordCatalog) DropAlias(ctx context.Context, dbName string, alias string, ts uint64) error {
	ret := _m.Called(ctx, dbName, alias, ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, uint64) error); ok {
		r0 = rf(ctx, dbName, alias, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DropCollection provides a mock function with given fields: ctx, collectionInfo, ts
func (_m *RootCoordCatalog) DropCollection(ctx context.Context, collectionInfo *model.Collection, ts uint64) error {
	ret := _m.Called(ctx, collectionInfo, ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Collection, uint64) error); ok {
		r0 = rf(ctx, collectionInfo, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DropCredential provides a mock function with given fields: ctx, username
func (_m *RootCoordCatalog) DropCredential(ctx context.Context, username string) error {
	ret := _m.Called(ctx, username)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DropDatabase provides a mock function with given fields: ctx, dbName, ts
func (_m *RootCoordCatalog) DropDatabase(ctx context.Context, dbName string, ts uint64) error {
	ret := _m.Called(ctx, dbName, ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64) error); ok {
		r0 = rf(ctx, dbName, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DropPartition provides a mock function with given fields: ctx, dbName, collectionID, partitionID, ts
func (_m *RootCoordCatalog) DropPartition(ctx context.Context, dbName string, collectionID int64, partitionID int64, ts uint64) error {
	ret := _m.Called(ctx, dbName, collectionID, partitionID, ts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64, uint64) error); ok {
		r0 = rf(ctx, dbName, collectionID, partitionID, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DropRole provides a mock function with given fields: ctx, tenant, roleName
func (_m *RootCoordCatalog) DropRole(ctx context.Context, tenant string, roleName string) error {
	ret := _m.Called(ctx, tenant, roleName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, tenant, roleName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCollectionByID provides a mock function with given fields: ctx, dbName, ts, collectionID
func (_m *RootCoordCatalog) GetCollectionByID(ctx context.Context, dbName string, ts uint64, collectionID int64) (*model.Collection, error) {
	ret := _m.Called(ctx, dbName, ts, collectionID)

	var r0 *model.Collection
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, int64) *model.Collection); ok {
		r0 = rf(ctx, dbName, ts, collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, uint64, int64) error); ok {
		r1 = rf(ctx, dbName, ts, collectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCollectionByName provides a mock function with given fields: ctx, dbName, collectionName, ts
func (_m *RootCoordCatalog) GetCollectionByName(ctx context.Context, dbName string, collectionName string, ts uint64) (*model.Collection, error) {
	ret := _m.Called(ctx, dbName, collectionName, ts)

	var r0 *model.Collection
	if rf, ok := ret.Get(0).(func(context.Context, string, string, uint64) *model.Collection); ok {
		r0 = rf(ctx, dbName, collectionName, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, uint64) error); ok {
		r1 = rf(ctx, dbName, collectionName, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCredential provides a mock function with given fields: ctx, username
func (_m *RootCoordCatalog) GetCredential(ctx context.Context, username string) (*model.Credential, error) {
	ret := _m.Called(ctx, username)

	var r0 *model.Credential
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Credential); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Credential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAliases provides a mock function with given fields: ctx, dbName, ts
func (_m *RootCoordCatalog) ListAliases(ctx context.Context, dbName string, ts uint64) ([]*model.Alias, error) {
	ret := _m.Called(ctx, dbName, ts)

	var r0 []*model.Alias
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64) []*model.Alias); ok {
		r0 = rf(ctx, dbName, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Alias)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, uint64) error); ok {
		r1 = rf(ctx, dbName, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCollections provides a mock function with given fields: ctx, dbName, ts
func (_m *RootCoordCatalog) ListCollections(ctx context.Context, dbName string, ts uint64) (map[string]*model.Collection, error) {
	ret := _m.Called(ctx, dbName, ts)

	var r0 map[string]*model.Collection
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64) map[string]*model.Collection); ok {
		r0 = rf(ctx, dbName, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*model.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, uint64) error); ok {
		r1 = rf(ctx, dbName, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCredentials provides a mock function with given fields: ctx
func (_m *RootCoordCatalog) ListCredentials(ctx context.Context) ([]string, error) {
	ret := _m.Called(ctx)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDatabases provides a mock function with given fields: ctx, ts
func (_m *RootCoordCatalog) ListDatabases(ctx context.Context, ts uint64) ([]string, error) {
	ret := _m.Called(ctx, ts)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []string); ok {
		r0 = rf(ctx, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListGrant provides a mock function with given fields: ctx, tenant, entity
func (_m *RootCoordCatalog) ListGrant(ctx context.Context, tenant string, entity *milvuspb.GrantEntity) ([]*milvuspb.GrantEntity, error) {
	ret := _m.Called(ctx, tenant, entity)

	var r0 []*milvuspb.GrantEntity
	if rf, ok := ret.Get(0).(func(context.Context, string, *milvuspb.GrantEntity) []*milvuspb.GrantEntity); ok {
		r0 = rf(ctx, tenant, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*milvuspb.GrantEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *milvuspb.GrantEntity) error); ok {
		r1 = rf(ctx, tenant, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPolicy provides a mock function with given fields: ctx, tenant
func (_m *RootCoordCatalog) ListPolicy(ctx context.Context, tenant string) ([]string, error) {
	ret := _m.Called(ctx, tenant)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRole provides a mock function with given fields: ctx, tenant, entity, includeUserInfo
func (_m *RootCoordCatalog) ListRole(ctx context.Context, tenant string, entity *milvuspb.RoleEntity, includeUserInfo bool) ([]*milvuspb.RoleResult, error) {
	ret := _m.Called(ctx, tenant, entity, includeUserInfo)

	var r0 []*milvuspb.RoleResult
	if rf, ok := ret.Get(0).(func(context.Context, string, *milvuspb.RoleEntity, bool) []*milvuspb.RoleResult); ok {
		r0 = rf(ctx, tenant, entity, includeUserInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*milvuspb.RoleResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *milvuspb.RoleEntity, bool) error); ok {
		r1 = rf(ctx, tenant, entity, includeUserInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUser provides a mock function with given fields: ctx, tenant, entity, includeRoleInfo
func (_m *RootCoordCatalog) ListUser(ctx context.Context, tenant string, entity *milvuspb.UserEntity, includeRoleInfo bool) ([]*milvuspb.UserResult, error) {
	ret := _m.Called(ctx, tenant, entity, includeRoleInfo)

	var r0 []*milvuspb.UserResult
	if rf, ok := ret.Get(0).(func(context.Context, string, *milvuspb.UserEntity, bool) []*milvuspb.UserResult); ok {
		r0 = rf(ctx, tenant, entity, includeRoleInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*milvuspb.UserResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *milvuspb.UserEntity, bool) error); ok {
		r1 = rf(ctx, tenant, entity, includeRoleInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUserRole provides a mock function with given fields: ctx, tenant
func (_m *RootCoordCatalog) ListUserRole(ctx context.Context, tenant string) ([]string, error) {
	ret := _m.Called(ctx, tenant)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRootCoordCatalog interface {
	mock.TestingT
	Cleanup(func())
}

// NewRootCoordCatalog creates a new instance of RootCoordCatalog. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRootCoordCatalog(t mockConstructorTestingTNewRootCoordCatalog) *RootCoordCatalog {
	mock := &RootCoordCatalog{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
