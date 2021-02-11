package zenkit

/*
 * Zenoss CONFIDENTIAL
 * __________________
 *
 *  This software Copyright (c) Zenoss, Inc. 2018
 *  All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and remains the property of Zenoss Incorporated
 * and its suppliers, if any.  The intellectual and technical concepts contained herein are owned
 * and proprietary to Zenoss Incorporated and its suppliers and may be covered by U.S. and Foreign
 * Patents, patents in process, and are protected by U.S. and foreign trade secret or copyright law.
 * Dissemination of this information or reproduction of any this material herein is strictly forbidden
 * unless prior written permission by an authorized officer is obtained from Zenoss Incorporated.
 */

import (
	"context"
	"github.com/pkg/errors"

	"github.com/mitchellh/go-testing-interface"
	"github.com/sirupsen/logrus"
)

const (

	// AuditActionAdd represents an action to create an audited resource a when logging.
	AuditActionCreate = "create"

	// AuditActionDelete represents an action to delete an audited resource a when logging.
	AuditActionDelete = "delete"

	// AuditActionUpdate represents an action to update an audited resource a when logging.
	AuditActionUpdate = "update"

	// EntryFieldAction is the name of the action field in an audit log entry
	EntryFieldAction = "zing.audit.action"

	// EntryFieldAuditedResourceType is the name of audited resource type field in an audit log entry
	EntryFieldAuditedResourceType = "zing.audit.resource.type"

	// EntryFieldAuditedResourceID is the name of audited resource ID in an audit log entry.
	EntryFieldAuditedResourceID = "zing.audit.resource.id"
)

// Entity is the interface for encapsulating an object's name
// and type for use with modules such as a logger.
type Entity interface {
	GetID() string
	GetType() string
}

// AuditLogger is the interface for audit logging.  Any implementations for
// audit logging should implement this interface.
type AuditLogger interface {

	// Set the action that we are auditing.
	Action(action string) AuditLogger

	// Set the message that we are writing to the audit log.
	Message(message string) AuditLogger

	// Set the type of entity being modified.
	Type(theType string) AuditLogger

	// Set the id of the entity being modified.
	ID(id string) AuditLogger

	// Set the type of entity being modified.
	Entity(entity Entity) AuditLogger

	// Add an additional field to the entry.
	WithField(name string, value interface{}) AuditLogger

	// Add additional fields to the entry.
	WithFields(fields logrus.Fields) AuditLogger

	// Log that the action succeeded.
	Succeeded()

	// Log that the action failed.
	Failed()

	// Log whether the action succeeded or failed based on the value passed in.
	SucceededIf(value bool)

	// Log whether the action succeeded or failed based on the error passed in.
	Error(err error) error
}

type logger struct {
	entry   *logrus.Entry
	message string
}

var _ AuditLogger = &logger{}

// NewAuditLogger returns a default implementation of the audit logger.
func NewAuditLogger(ctx context.Context) AuditLogger {
	auditLogEntry := ContextAuditLogger(ctx).WithFields(
		logrus.Fields{
			LogTypeField: LogTypeAudit,
		},
	)
	return &logger{entry: auditLogEntry}
}

func (l *logger) addFields(fields logrus.Fields) AuditLogger {
	l.entry = l.entry.WithFields(fields)
	return l
}

func (l *logger) addField(name string, value interface{}) AuditLogger {
	return l.addFields(logrus.Fields{name: value})
}

func (l *logger) newLoggerWithFields(fields logrus.Fields) *logger {
	result := &logger{
		entry:   l.entry,
		message: l.message,
	}
	result.addFields(fields)
	return result
}

func (l *logger) newLoggerWith(name string, value interface{}) *logger {
	return l.newLoggerWithFields(logrus.Fields{name: value})
}

func (l *logger) Action(action string) AuditLogger {
	return l.newLoggerWith(EntryFieldAction, action)
}

func (l *logger) Message(message string) AuditLogger {
	l.message = message
	return l
}

func (l *logger) Type(theType string) AuditLogger {
	return l.newLoggerWith(EntryFieldAuditedResourceType, theType)
}

func (l *logger) ID(id string) AuditLogger {
	return l.newLoggerWith(EntryFieldAuditedResourceID, id)
}

func (l *logger) Entity(entity Entity) AuditLogger {
	return l.newLoggerWithFields(logrus.Fields{
		EntryFieldAuditedResourceID:   entity.GetID(),
		EntryFieldAuditedResourceType: entity.GetType(),
	})
}

func (l *logger) WithField(name string, value interface{}) AuditLogger {
	return l.newLoggerWith(name, value)
}

func (l *logger) WithFields(fields logrus.Fields) AuditLogger {
	return l.newLoggerWithFields(fields)
}

func (l *logger) log(success bool) {
	l.addField("success", success)
	if l.message == "" {
		//pkgLogger := Logger("audit-pkg-logger").WithFields(l.entry.Data)
		l.entry.Error("attempting to audit log empty message")
	}
	if success {
		l.entry.Info(l.message)
	} else {
		l.entry.Warn(l.message)
	}
}

func (l *logger) Succeeded() {
	l.log(true)
}

func (l *logger) Failed() {
	l.log(false)
}

func (l *logger) SucceededIf(value bool) {
	l.log(value)
}

func (l *logger) Error(err error) error {
	l.log(err == nil)
	return err
}

func GetLogrusLogger(t testing.T, l AuditLogger) (*logrus.Logger, error) {
	if aLog, ok := l.(*logger); ok {
		return aLog.entry.Logger, nil
	}
	return nil, errors.New("invalid type")
}
