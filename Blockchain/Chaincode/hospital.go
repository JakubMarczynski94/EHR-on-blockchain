package main

import (
	"encoding/json"
	. "fmt"
	"time"

	"github.com/google/uuid"
)

func (c *Chaincode) CreateNewReport(ctx CustomTransactionContextInterface, patientID, refDoctor string) (string, error) {
	id := uuid.New().String()
	report := Report{
		DocTyp:      REPORT,
		ID:          id,
		PatientID:   patientID,
		Status:      "0",
		RefDoctorID: refDoctor,
		Comments:    make(map[int64]string),
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	}
	reportAsByte, _ := json.Marshal(report)
	return report.ID, ctx.GetStub().PutState(id, reportAsByte)
}

func (c *Chaincode) StartTreatment(ctx CustomTransactionContextInterface, treatmentID, supervisor string) error {
	if ctx.GetData() == nil {
		return Errorf("Treatment with ID %v doesn't exists", treatmentID)
	}
	var treatment Treatment
	json.Unmarshal(ctx.GetData(), &treatment)
	treatment.Supervisor = supervisor
	treatment.Status = 1
	treatment.UpdateTime = time.Now().Unix()

	treatmentAsByte, _ := json.Marshal(treatment)

	return ctx.GetStub().PutState(treatment.ID, treatmentAsByte)
}