/*
 * Zenoss CONFIDENTIAL
 * __________________
 *
 *  This software Copyright (c) Zenoss, Inc. 2019
 *  All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and remains the property of Zenoss Incorporated
 * and its suppliers, if any.  The intellectual and technical concepts contained herein are owned
 * and proprietary to Zenoss Incorporated and its suppliers and may be covered by U.S. and Foreign
 * Patents, patents in process, and are protected by U.S. and foreign trade secret or copyright law.
 * Dissemination of this information or reproduction of any this material herein is strictly forbidden
 * unless prior written permission by an authorized officer is obtained from Zenoss Incorporated.
 */
package mocks

import "github.com/zenoss/zenkit/v5"

type MockTenantIdentity struct {
	tenant     string
	id         string
	clientID   string
	connection string
	scopes     []string
	email      string
	subject    string
}

var _ zenkit.TenantIdentity = &MockTenantIdentity{} // verify the interface is implemented correctly

func NewMockTenantIdentity(tenant, id string) zenkit.TenantIdentity {
	return &MockTenantIdentity{tenant: tenant, id: id}
}

func (ti *MockTenantIdentity) Tenant() string {
	return ti.tenant
}

func (ti *MockTenantIdentity) ID() string {
	return ti.id
}

func (ti *MockTenantIdentity) ClientID() string {
	return ti.clientID
}

func (ti *MockTenantIdentity) Connection() (connection string) {
	return ti.connection
}

func (ti *MockTenantIdentity) Scopes() (scopes []string) {
	return ti.scopes
}

func (ti *MockTenantIdentity) HasScope(scope string) bool {
	for _, s := range ti.scopes {
		if s == scope {
			return true
		}
	}
	return false
}

func (ti *MockTenantIdentity) Email() (email string) {
	return ti.email
}

func (ti *MockTenantIdentity) GetSubject() (subject string) {
	return ti.subject
}

func (ti *MockTenantIdentity) SetClientID(value string) {
	ti.clientID = value
}

func (ti *MockTenantIdentity) SetConnection(value string) {
	ti.connection = value
}

func (ti *MockTenantIdentity) SetScopes(value []string) {
	ti.scopes = []string{}
	ti.scopes = append(ti.scopes, value...)
}

func (ti *MockTenantIdentity) SetEmail(value string) {
	ti.email = value
}

func (ti *MockTenantIdentity) SetSubject(value string) {
	ti.subject = value
}
