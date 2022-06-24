/*
 * Zenoss CONFIDENTIAL
 * __________________
 *
 *  This software Copyright (c) Zenoss, Inc. 2021
 *  All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and remains the property of Zenoss Incorporated
 * and its suppliers, if any.  The intellectual and technical concepts contained herein are owned
 * and proprietary to Zenoss Incorporated and its suppliers and may be covered by U.S. and Foreign
 * Patents, patents in process, and are protected by U.S. and foreign trade secret or copyright law.
 * Dissemination of this information or reproduction of any this material herein is strictly forbidden
 * unless prior written permission by an authorized officer is obtained from Zenoss Incorporated.
 */
package filters

import (
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

// Format of clause to return from GetEntityClauses
type Format int8

const (
	// Elastic specific format
	Elastic Format = 1
)

// Data contains multiple Filters
type Data struct {
	Filter []*Filter `json:"filters"`
}

// Filter is a struct containing information about the user group and entity filters
// event/notification filters tbd
type Filter struct {
	Group        string        `json:"group"`
	EntityFilter *EntityFilter `json:"entityFilter"`
}

// EntityFilter contains the information specific to an entity filter
type EntityFilter struct {
	Clause string `json:"clause"`
}

// EntityClause is the default clause that is not specific to any format
type EntityClause struct {
	Field  string `json:"field"`
	Path   string `json:"path"`
	Source string `json:"source"`
	Value  string `json:"value"`
}

// GetEntityClauses returns the entity clauses associated with a restriction filter given a context
// callers responsibility to convert or handle the returned interfaces
func GetEntityClauses(ctx context.Context, format Format) (clauses []interface{}, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	var filters []string
	if val, ok := md["_restrictionfilters"]; ok {
		filters = val
	}
	switch format {
	case Elastic:
		clauses, err = parseEntityFiltersToElastic(filters)
	default:
		clauses, err = parseEntityFilters(filters)
	}
	return clauses, err
}

// TODO
// get event/notification filters once defined

func parseEntityFilters(filters []string) (clauses []interface{}, err error) {
	marshaled, err := json.Marshal(filters)
	if err != nil {
		logrus.WithError(err).Error("Error mashalling filters to json")
		return nil, err
	}
	var unmarshaledFilters []string
	if err := json.Unmarshal(marshaled, &unmarshaledFilters); err != nil {
		logrus.WithError(err).Error("Error unmarshalling filters")
		return nil, err
	}

	for _, val := range unmarshaledFilters {
		data := Data{}
		if err := json.Unmarshal([]byte(val), &data); err != nil {
			logrus.WithError(err).Error("Error unmarshalling each filter to struct")
			return nil, err
		}
		for _, filter := range data.Filter {
			clause := filter.EntityFilter.Clause
			var unmarshaledClauses []EntityClause
			if err := json.Unmarshal([]byte(clause), &unmarshaledClauses); err != nil {
				logrus.WithError(err).Error("Error unmarshalling each clause")
				return nil, err
			}
			for _, clause := range unmarshaledClauses {
				clauses = append(clauses, clause)
			}
		}
	}
	return clauses, nil
}

// clauses returned should be converted to []map[string]interface{}
func parseEntityFiltersToElastic(filters []string) (clauses []interface{}, err error) {
	marshaled, err := json.Marshal(filters)
	if err != nil {
		logrus.WithError(err).Error("Error mashalling filters to json")
		return nil, err
	}
	var unmarshaledFilters []string
	if err := json.Unmarshal(marshaled, &unmarshaledFilters); err != nil {
		logrus.WithError(err).Error("Error unmarshalling filters")
		return nil, err
	}

	for _, val := range unmarshaledFilters {
		data := Data{}
		if err := json.Unmarshal([]byte(val), &data); err != nil {
			logrus.WithError(err).Error("Error unmarshalling each filter to struct")
			return nil, err
		}
		for _, filter := range data.Filter {
			clause := filter.EntityFilter.Clause
			var unmarshaledClauses []map[string]string
			if err := json.Unmarshal([]byte(clause), &unmarshaledClauses); err != nil {
				logrus.WithError(err).Error("Error unmarshalling each clause")
				return nil, err
			}

			for _, clause := range unmarshaledClauses {
				sourceClause := map[string]interface{}{
					"nested": map[string]interface{}{
						"path": "fields",
						"query": map[string]interface{}{
							"bool": map[string]interface{}{
								"must": []interface{}{
									map[string]interface{}{"term": map[string]interface{}{"fields.key": "source"}},
									map[string]interface{}{"match": map[string]interface{}{"fields.string_values": clause["source"]}},
								},
							},
						},
						"score_mode": "none",
					},
				}
				clause := map[string]interface{}{
					"nested": map[string]interface{}{
						"path": "fields",
						"query": map[string]interface{}{
							"bool": map[string]interface{}{
								"minimum_should_match": 1,
								"must": []interface{}{
									map[string]interface{}{"term": map[string]interface{}{"fields.key": clause["field"]}},
								},
								"should": []interface{}{
									map[string]interface{}{"match": map[string]interface{}{"fields.string_values": clause["path"]}},
									map[string]interface{}{"prefix": map[string]interface{}{"fields.string_values": clause["path"] + "/"}},
								},
							},
						},
						"score_mode": "none",
					},
				}
				clauses = append(clauses, sourceClause)
				clauses = append(clauses, clause)
			}
		}
	}
	return clauses, nil
}
