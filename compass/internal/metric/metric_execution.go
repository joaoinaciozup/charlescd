package metric

import (
	"compass/internal/util"
	"encoding/json"
	"io"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

const (
	MetricReached = "REACHED"
	MetricActive  = "ACTIVE"
	MetricError   = "ERROR"
	MetricUpdated = "UPDATED"
)

type MetricExecution struct {
	util.BaseModel `json:"-"`
	MetricID       uuid.UUID `json:"-"`
	LastValue      float64   `json:"lastValue"`
	Status         string    `json:"status"`
}

func (main Main) ParseMetricExecution(metricExecutionExecution io.ReadCloser) (MetricExecution, error) {
	var newMetricExecution *MetricExecution
	err := json.NewDecoder(metricExecutionExecution).Decode(&newMetricExecution)
	if err != nil {
		util.Error(util.GeneralParseError, "ParseMetricExecution", err, metricExecutionExecution)
		return MetricExecution{}, err
	}
	return *newMetricExecution, nil
}

func (main Main) FindAllMetricExecutions() ([]MetricExecution, error) {
	var metricExecutions []MetricExecution
	db := main.db.Find(&metricExecutions)
	if db.Error != nil {
		util.Error(util.FindAllMetricExecutionsError, "FindAllMetricExecutions", db.Error, metricExecutions)
		return []MetricExecution{}, db.Error
	}
	return metricExecutions, nil
}

func (main Main) FindAllActivesMetricExecutions() ([]MetricExecution, error) {
	var metricExecutions []MetricExecution
	db := main.db.Find(&metricExecutions)
	if db.Error != nil {
		util.Error(util.FindAllMetricExecutionsError, "FindAllMetricExecutions", db.Error, metricExecutions)
		return []MetricExecution{}, db.Error
	}
	return metricExecutions, nil
}

func (main Main) FindMetricExecutionById(id string) (MetricExecution, error) {
	metricExecution := MetricExecution{}
	db := main.db.Where("id = ?", id).First(&metricExecution)
	if db.Error != nil {
		util.Error(util.FindMetricExecutionByIdError, "FindMetricExecutionById", db.Error, "Id = "+id)
		return MetricExecution{}, db.Error
	}
	return metricExecution, nil
}

func (main Main) UpdateMetricExecution(metricExecution MetricExecution) (MetricExecution, error) {
	db := main.db.Model(&metricExecution).Updates(metricExecution)
	if db.Error != nil {
		util.Error(util.UpdateMetricExecutionError, "UpdateMetricExecution", db.Error, metricExecution)
		return MetricExecution{}, db.Error
	}
	return metricExecution, nil
}

func (main Main) saveMetricExecution(tx *gorm.DB, execution MetricExecution) (MetricExecution, error) {
	db := tx.Save(&execution)
	if db.Error != nil {
		util.Error(util.SaveMetricExecutionError, "SaveMetricExecution", db.Error, execution)
		return MetricExecution{}, db.Error
	}
	return execution, nil
}

func (main Main) removeMetricExecution(tx *gorm.DB, id string) error {
	db := tx.Where("id = ?", id).Delete(MetricExecution{})
	if db.Error != nil {
		util.Error(util.RemoveMetricExecutionError, "RemoveMetricExecution", db.Error, id)
		return db.Error
	}
	return nil
}
