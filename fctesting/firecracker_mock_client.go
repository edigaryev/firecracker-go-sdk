// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package fctesting

import (
	ops "github.com/firecracker-microvm/firecracker-go-sdk/client/operations"
)

type MockClient struct {
	GetMachineConfigFn             func(params *ops.GetMachineConfigParams) (*ops.GetMachineConfigOK, error)
	GetMmdsFn                      func(params *ops.GetMmdsParams) (*ops.GetMmdsOK, error)
	PatchMmdsFn                    func(params *ops.PatchMmdsParams) (*ops.PatchMmdsNoContent, error)
	PutMmdsFn                      func(params *ops.PutMmdsParams) (*ops.PutMmdsNoContent, error)
	CreateSyncActionFn             func(params *ops.CreateSyncActionParams) (*ops.CreateSyncActionNoContent, error)
	DescribeInstanceFn             func(params *ops.DescribeInstanceParams) (*ops.DescribeInstanceOK, error)
	PatchGuestDriveByIDFn          func(params *ops.PatchGuestDriveByIDParams) (*ops.PatchGuestDriveByIDNoContent, error)
	PutGuestBootSourceFn           func(params *ops.PutGuestBootSourceParams) (*ops.PutGuestBootSourceNoContent, error)
	PutGuestDriveByIDFn            func(params *ops.PutGuestDriveByIDParams) (*ops.PutGuestDriveByIDNoContent, error)
	PutGuestNetworkInterfaceByIDFn func(params *ops.PutGuestNetworkInterfaceByIDParams) (*ops.PutGuestNetworkInterfaceByIDNoContent, error)
	PutGuestVsockByIDFn            func(params *ops.PutGuestVsockByIDParams) (*ops.PutGuestVsockByIDCreated, *ops.PutGuestVsockByIDNoContent, error)
	PutLoggerFn                    func(params *ops.PutLoggerParams) (*ops.PutLoggerNoContent, error)
	PutMachineConfigurationFn      func(params *ops.PutMachineConfigurationParams) (*ops.PutMachineConfigurationNoContent, error)
}

func (c *MockClient) GetMachineConfig(params *ops.GetMachineConfigParams) (*ops.GetMachineConfigOK, error) {
	if c.GetMachineConfigFn != nil {
		return c.GetMachineConfigFn(params)
	}

	return nil, nil
}

func (c *MockClient) GetMmds(params *ops.GetMmdsParams) (*ops.GetMmdsOK, error) {
	if c.GetMmdsFn != nil {
		return c.GetMmdsFn(params)
	}

	return nil, nil
}

func (c *MockClient) PatchMmds(params *ops.PatchMmdsParams) (*ops.PatchMmdsNoContent, error) {
	if c.PatchMmdsFn != nil {
		return c.PatchMmdsFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutMmds(params *ops.PutMmdsParams) (*ops.PutMmdsNoContent, error) {
	if c.PutMmdsFn != nil {
		return c.PutMmdsFn(params)
	}

	return nil, nil
}

func (c *MockClient) CreateSyncAction(params *ops.CreateSyncActionParams) (*ops.CreateSyncActionNoContent, error) {
	if c.CreateSyncActionFn != nil {
		return c.CreateSyncActionFn(params)
	}

	return nil, nil
}

func (c *MockClient) DescribeInstance(params *ops.DescribeInstanceParams) (*ops.DescribeInstanceOK, error) {
	if c.DescribeInstanceFn != nil {
		return c.DescribeInstanceFn(params)
	}

	return nil, nil
}

func (c *MockClient) PatchGuestDriveByID(params *ops.PatchGuestDriveByIDParams) (*ops.PatchGuestDriveByIDNoContent, error) {
	if c.PatchGuestDriveByIDFn != nil {
		return c.PatchGuestDriveByIDFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutGuestBootSource(params *ops.PutGuestBootSourceParams) (*ops.PutGuestBootSourceNoContent, error) {
	if c.PutGuestBootSourceFn != nil {
		return c.PutGuestBootSourceFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutGuestDriveByID(params *ops.PutGuestDriveByIDParams) (*ops.PutGuestDriveByIDNoContent, error) {
	if c.PutGuestDriveByIDFn != nil {
		return c.PutGuestDriveByIDFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutGuestNetworkInterfaceByID(params *ops.PutGuestNetworkInterfaceByIDParams) (*ops.PutGuestNetworkInterfaceByIDNoContent, error) {
	if c.PutGuestNetworkInterfaceByIDFn != nil {
		return c.PutGuestNetworkInterfaceByIDFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutGuestVsockByID(params *ops.PutGuestVsockByIDParams) (*ops.PutGuestVsockByIDCreated, *ops.PutGuestVsockByIDNoContent, error) {
	if c.PutGuestVsockByIDFn != nil {
		return c.PutGuestVsockByIDFn(params)
	}

	return nil, nil, nil
}

func (c *MockClient) PutLogger(params *ops.PutLoggerParams) (*ops.PutLoggerNoContent, error) {
	if c.PutLoggerFn != nil {
		return c.PutLoggerFn(params)
	}

	return nil, nil
}

func (c *MockClient) PutMachineConfiguration(params *ops.PutMachineConfigurationParams) (*ops.PutMachineConfigurationNoContent, error) {
	if c.PutMachineConfigurationFn != nil {
		return c.PutMachineConfigurationFn(params)
	}

	return nil, nil
}