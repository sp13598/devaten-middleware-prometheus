package storage_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/team7mysupermon/devaten_middlewareapp/storage"
)

var (
	RUN_TEST = []byte(`{
		"status":"200 OK",
		"responseCode":200,
		"data":{
			"idNum":1682,
			"usecaseIdentifier":"MIDDELWAREAPITEST",
			"applicationId":203,
			"applicationName":"Konakart_Application",
			"applicationIdentifier":"d9cdf882-6f1e-45d3-b8ca-2d1b19d9712e",
			"runSituationResult":[
				{
					"dataSourceId":232,
					"databaseType":"MySQL",
					"databaseName":"konakart",
					"schemaName":null,
					"hostUrl":"34.88.216.230",
					"data":{
						"SUM_ROWS_AFFECTED":0,
						"SUM_SELECT_RANGE":0,
						"SUM_LOCK_TIME":217000000.0000,
						"SUM_SORT_ROWS":0,
						"SUM_ERRORS":0,
						"SUM_ROWS_SENT":null,
						"SUM_SELECT_SCAN":1,
						"SUM_NO_GOOD_INDEX_USED":0,
						"EXEC_TIME_MAX":null,
						"SUM_SORT_SCAN":0,
						"SUM_SELECT_RANGE_CHECK":0,
						"SUM_TIMER_WAIT":769276000.0000,
						"USECASE_IDENTIFIER":"MIDDELWAREAPITEST",
						"STARTTIMESTMAP":"2022-05-09T12:38:40.000+0000",
						"SUM_ROWS_EXAMINED":null,
						"SUM_SELECT_FULL_JOIN":0,
						"SUM_NO_INDEX_USED":1,
						"COUNT_STAR":1,
						"SUM_SELECT_FULL_RANGE_JOIN":0,
						"SUM_SORT_MERGE_PASSES":0,
						"SUM_SORT_RANGE":0
					}
				}
			]
		},
		"errorMessage":null,
		"errorCode":null,
		"reportLink":null
	}`)
)

func setUpRunData() storage.RunAutoGenerated {
	var runAutoGenerated storage.RunAutoGenerated

	err := json.Unmarshal(RUN_TEST, &runAutoGenerated)
	if err != nil {
		log.Panicln(err)
	}

	return runAutoGenerated
}

func TestRunAutoGenerated(t *testing.T) {
	data := setUpRunData()

	assert.Equal(t, data.Status, "200 OK")
	assert.Equal(t, data.ResponseCode, 200)
	assert.Equal(t, data.ErrorMessage, nil)
	assert.Equal(t, data.ErrorCode, nil)
	assert.Equal(t, data.ReportLink, nil)
}

func TestRunMetaData(t *testing.T) {
	data := setUpRunData().RunMetaData

	assert.Equal(t, data.IDNum, 1682)
	assert.Equal(t, data.UsecaseIdentifier, "MIDDELWAREAPITEST")
	assert.Equal(t, data.ApplicationID, 203)
	assert.Equal(t, data.ApplicationName, "Konakart_Application")
	assert.Equal(t, data.ApplicationIdentifier, "d9cdf882-6f1e-45d3-b8ca-2d1b19d9712e")
}

func TestRunSituationResult(t *testing.T) {
	data := setUpRunData().RunMetaData.RunSituationResult[0]

	assert.Equal(t, data.DataSourceID, 232)
	assert.Equal(t, data.DatabaseType, "MySQL")
	assert.Equal(t, data.DatabaseName, "konakart")
	assert.Equal(t, data.SchemaName, nil)
	assert.Equal(t, data.HostURL, "34.88.216.230")
}

func TestRunData(t *testing.T) {
	data := setUpRunData().RunMetaData.RunSituationResult[0].RunData

	assert.Equal(t, data.SumRowsAffected, 0)
	assert.Equal(t, data.SumSelectRange, 0)
	assert.Equal(t, data.SumLockTime, 217000000.0000)
	assert.Equal(t, data.SumSortRows, 0)
	assert.Equal(t, data.SumErrors, 0)
	assert.Equal(t, data.SumRowsSent, nil)
	assert.Equal(t, data.SumSelectScan, 1)
	assert.Equal(t, data.SumNoGoodIndexUsed, 0)
	assert.Equal(t, data.ExecTimeMax, nil)
	assert.Equal(t, data.SumSortScan, 0)
	assert.Equal(t, data.SumSelectRangeCheck, 0)
	assert.Equal(t, data.SumTimerWait, 769276000.0000)
	assert.Equal(t, data.UsecaseIdentifier, "MIDDELWAREAPITEST")
	assert.Equal(t, data.Starttimestmap, "2022-05-09T12:38:40.000+0000")
	assert.Equal(t, data.SumRowsExamined, nil)
	assert.Equal(t, data.SumSelectFullJoin, 0)
	assert.Equal(t, data.SumNoIndexUsed, 1)
	assert.Equal(t, data.CountStar, 1)
	assert.Equal(t, data.SumSelectFullRangeJoin, 0)
	assert.Equal(t, data.SumSortMergePasses, 0)
	assert.Equal(t, data.SumSortRange, 0)
}
