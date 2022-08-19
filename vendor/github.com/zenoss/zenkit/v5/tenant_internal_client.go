/*
 * Zenoss CONFIDENTIAL
 * __________________
 *
 * This software Copyright (c) Zenoss, Inc. 2022
 * All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and remains the property of Zenoss incorporated
 * and its suppliers, if any.  The intellectual and technical concepts contained herein are owned
 * and proprietary to Zenoss Incorporated and its suppliers and may be covered by U.S. and Foreign
 * Patents, patents in process, and are protected by U.S. and foreign trade secret or copyright law.
 * Dissemination of this information or reproduction of any this material herein is strictly forbidden
 * unless prior written permission by an authorized officer is obtained from Zenoss Incorporated.
 *
 */

package zenkit

import (
	"context"
	"sync"

	"github.com/pkg/errors"

	proto "github.com/zenoss/zing-proto/v11/go/cloud/tenant"
)

var (
	tenantDataIdCache = TenantCache{cache: make(map[string]string), mutex: sync.Mutex{}}
	tenantIdCache     = TenantCache{cache: make(map[string]string), mutex: sync.Mutex{}}
	ErrMissingRequest = errors.New("Request can not be empty")
)

type TenantCache struct {
	cache map[string]string
	mutex sync.Mutex
}

type TenantInternalServiceClient struct {
	client proto.TenantInternalServiceClient
}

func CreateTenantInternalServiceProtoClient(ctx context.Context) (proto.TenantInternalServiceClient, error) {
	conn, err := NewClientConn(ctx, "tenant-svc-internal")
	if err != nil {
		return nil, err
	}

	protoClient := proto.NewTenantInternalServiceClient(conn)
	return protoClient, nil
}

func CreateNewTenantInternalServiceClient(ctx context.Context) (tcl *TenantInternalServiceClient) {
	protoClient, err := CreateTenantInternalServiceProtoClient(ctx)
	if err != nil {
		return nil
	}
	return &TenantInternalServiceClient{
		client: protoClient,
	}
}

func NewTenantInternalServiceClient(protoClient proto.TenantInternalServiceClient) (tcl *TenantInternalServiceClient) {
	if protoClient == nil {
		return nil
	}

	return &TenantInternalServiceClient{
		client: protoClient,
	}
}

func (tcl *TenantInternalServiceClient) GetTenantId(ctx context.Context, req *proto.GetTenantIdRequest) (string, error) {
	log := ContextLogger(ctx)
	if req == nil {
		return "", ErrMissingRequest
	}
	arg, _ := GetTenantValue(req.GetTenantInfo())

	tenantIdCache.mutex.Lock()
	if val, ok := tenantIdCache.cache[arg]; ok {
		tenantIdCache.mutex.Unlock()
		log.Debugf("Found %s matching Id in Id Cache", arg)
		return val, nil
	}
	tenantIdCache.mutex.Unlock()
	log.Debugf("Missed %s in tenant Id Cache", arg)

	response, err := tcl.client.GetTenantId(ctx, req)
	if err != nil {
		return "", err
	}

	tenantIdCache.mutex.Lock()
	tenantIdCache.cache[arg] = response.TenantId
	tenantIdCache.mutex.Unlock()
	log.Debugf("Saved %s in tenant Id Cache", response.TenantId)

	return response.TenantId, nil
}

func (tcl *TenantInternalServiceClient) GetTenantIds(ctx context.Context, req *proto.GetTenantIdsRequest) ([]*proto.TenantResult, error) {
	log := ContextLogger(ctx)
	var response []*proto.TenantResult = []*proto.TenantResult{}
	var filtered []*proto.TenantInfo
	if req == nil {
		return response, ErrMissingRequest
	}
	infos := req.GetTenantInfo()

	for _, info := range infos {
		arg, res := GetTenantValue(info)

		tenantIdCache.mutex.Lock()
		if val, ok := tenantIdCache.cache[arg]; ok {
			log.Debugf("Found %s matching id in Id Cache", arg)
			tenantIdCache.mutex.Unlock()
			res.TenantId = val
			response = append(response, res)
			continue
		}
		tenantIdCache.mutex.Unlock()
		log.Debugf("Missed %s in tenant Id Cache", arg)

		filtered = append(filtered, info)

	}

	if len(filtered) == 0 {
		return response, nil
	}

	req.TenantInfo = filtered

	result, err := tcl.client.GetTenantIds(ctx, req)
	if err != nil {
		return response, err
	}

	for _, tRes := range result.TenantIds {
		response = append(response, tRes)
		UpdateCache(ctx, tRes, tRes.TenantId, &tenantIdCache)

	}

	return response, nil
}

func (tcl *TenantInternalServiceClient) GetTenantDataIds(ctx context.Context, req *proto.GetTenantDataIdsRequest) ([]*proto.TenantResult, error) {
	log := ContextLogger(ctx)
	var response []*proto.TenantResult = []*proto.TenantResult{}
	var filtered []*proto.TenantInfo

	if req == nil {
		return response, ErrMissingRequest
	}

	infos := req.GetTenantInfo()
	for _, info := range infos {
		arg, res := GetTenantValue(info)

		tenantDataIdCache.mutex.Lock()
		if val, ok := tenantDataIdCache.cache[arg]; ok {
			tenantDataIdCache.mutex.Unlock()
			log.Debugf("Found %s matching dataId in dataId Cache", arg)
			res.TenantDataId = val
			response = append(response, res)
			continue
		}
		tenantDataIdCache.mutex.Unlock()
		log.Debugf("Missed %s in tenant dataId Cache", arg)

		filtered = append(filtered, info)

	}

	if len(filtered) == 0 {
		return response, nil
	}

	req.TenantInfo = filtered

	result, err := tcl.client.GetTenantDataIds(ctx, req)
	if err != nil {
		return response, err
	}

	for _, tRes := range result.TenantDataIds {
		response = append(response, tRes)
		UpdateCache(ctx, tRes, tRes.TenantDataId, &tenantDataIdCache)
	}

	return response, nil
}

func (tcl *TenantInternalServiceClient) GetTenantDataId(ctx context.Context, req *proto.GetTenantDataIdRequest) (string, error) {
	log := ContextLogger(ctx)
	if req == nil {
		return "", ErrMissingRequest
	}
	arg, _ := GetTenantValue(req.GetTenantInfo())

	tenantDataIdCache.mutex.Lock()
	if val, ok := tenantDataIdCache.cache[arg]; ok {
		log.Debugf("Found %s matching dataId in dataId Cache", arg)
		tenantDataIdCache.mutex.Unlock()
		return val, nil
	}
	tenantDataIdCache.mutex.Unlock()
	log.Debugf("Missed %s in tenant dataId Cache", arg)

	response, err := tcl.client.GetTenantDataId(ctx, req)
	if err != nil {
		return "", err
	}
	tenantDataIdCache.mutex.Lock()
	tenantDataIdCache.cache[arg] = response.TenantDataId
	tenantDataIdCache.mutex.Unlock()
	log.Debugf("Saved %s in tenant dataId Cache", response.TenantDataId)

	return response.TenantDataId, nil
}

func GetTenantValue(req *proto.TenantInfo) (string, *proto.TenantResult) {
	res := proto.TenantResult{}

	switch req.Field.(type) {
	case *proto.TenantInfo_TenantDataId:
		res.TenantDataId = req.GetTenantDataId()
		return req.GetTenantDataId(), &res
	case *proto.TenantInfo_TenantName:
		res.TenantName = req.GetTenantName()
		return req.GetTenantName(), &res
	case *proto.TenantInfo_TenantId:
		res.TenantId = req.GetTenantId()
		return req.GetTenantId(), &res
	default:
		res.TenantReference = req.GetTenantReference()
		return req.GetTenantReference(), &res
	}

}

func UpdateCache(ctx context.Context, tenantRes *proto.TenantResult, val string, cache *TenantCache) {
	log := ContextLogger(ctx)
	if tenantRes.GetTenantId() != "" {
		cache.mutex.Lock()
		cache.cache[tenantRes.GetTenantId()] = val
		cache.mutex.Unlock()
	}
	if tenantRes.GetTenantReference() != "" {
		cache.mutex.Lock()
		cache.cache[tenantRes.GetTenantReference()] = val
		cache.mutex.Unlock()
	}
	if tenantRes.GetTenantName() != "" {
		cache.mutex.Lock()
		cache.cache[tenantRes.GetTenantName()] = val
		cache.mutex.Unlock()
	}
	if tenantRes.GetTenantDataId() != "" {
		cache.mutex.Lock()
		cache.cache[tenantRes.GetTenantDataId()] = val
		cache.mutex.Unlock()
	}
	log.Debugf("Saved %s in tenant Cache", val)

}
