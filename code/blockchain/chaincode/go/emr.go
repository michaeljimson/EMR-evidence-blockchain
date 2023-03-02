package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type EmrContract struct {
	contractapi.Contract
}
type emrAsset struct {
	EmrID             string `json:"EmrID"`
	EmrName           string `json:"EmrName"`
	PatientID         string `json:"PatientID`
	EmrGeneHospitalID string `json:"HospitalID"`
	EmrGeneDoctorID   string `json:"DoctorID"`
	EmrIPFSHash       string `json:"EmrIPFSHash"`
}

func (s *EmrContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	emrs := []emrAsset{
		{EmrID: "123123124123123", EmrName: "zhangsansansan", PatientID: "111111111111111111", EmrGeneHospitalID: "123123123123", EmrGeneDoctorID: "1231231232123", EmrIPFSHash: "QmWinSCXWQDArYYrw9pH9b8RCtE1sKffvopgk1Q6kE1SKK"},
	}
	for _, emr := range emrs {
		emrBlockJson, err := json.Marshal(emr)
		if err != nil {
			return err
		}
		err = ctx.GetStub().PutState(emr.EmrID, emrBlockJson)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}
	return nil
}

// 判断病历是否存在
func (s *EmrContract) EmrExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	emrBlockJson, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state %v", err)
	}

	return emrBlockJson != nil, nil
}

// 写入新病历存证

func (s *EmrContract) CreateEmr(ctx contractapi.TransactionContextInterface, emrid string, emrname string, patient string, emrgenehospitalid string, emrgenedoctorid string, emripfshash string) error {
	exists, err := s.EmrExists(ctx, emrid)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the project %s already exists", emrid)
	}

	emr := emrAsset{
		EmrID:             emrid,
		EmrName:           emrname,
		PatientID:         patient,
		EmrGeneHospitalID: emrgenehospitalid,
		EmrGeneDoctorID:   emrgenedoctorid,
		EmrIPFSHash:       emripfshash,
	}
	emrBlockJson, err := json.Marshal(emr)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(emrid, emrBlockJson)
}

// 读取指定ID的病历信息
func (s *EmrContract) ReadEmr(ctx contractapi.TransactionContextInterface, emrid string) (*emrAsset, error) {
	emrBlockJson, err := ctx.GetStub().GetState(emrid)
	if err != nil {
		return nil, fmt.Errorf("failed to read form world state: %v", err)
	}
	if emrBlockJson == nil {
		return nil, fmt.Errorf("the emrBlock %s does not exist", emrid)
	}

	var emr emrAsset
	err = json.Unmarshal(emrBlockJson, &emr)
	if err != nil {
		return nil, err
	}

	return &emr, nil
}

// 更新指定ID的病历存证信息

func (s *EmrContract) UpdateEmr(ctx contractapi.TransactionContextInterface, emrid string, emrname string, patient string, emrgenehospitalid string, emrgenedoctorid string, emripfshash string) error {
	exists, err := s.EmrExists(ctx, emrid)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the project %s does not exist", emrid)
	}
	emr := emrAsset{
		EmrID:             emrid,
		EmrName:           emrname,
		PatientID:         patient,
		EmrGeneHospitalID: emrgenehospitalid,
		EmrGeneDoctorID:   emrgenedoctorid,
		EmrIPFSHash:       emripfshash,
	}
	emrBlockJson, err := json.Marshal(emr)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(emrid, emrBlockJson)
}

// 删除指定ID的病历存证信息

func (s *EmrContract) DeleteEmr(ctx contractapi.TransactionContextInterface, emrid string) error {
	exists, err := s.EmrExists(ctx, emrid)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the emrprob %s does not exists", emrid)
	}

	return ctx.GetStub().DelState(emrid)
}

func (s *EmrContract) GetAllEmrs(ctx contractapi.TransactionContextInterface) ([]*emrAsset, error) {
	//查询参数为两个空字符串时即查找所有数据
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var emrs []*emrAsset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var emr emrAsset
		err = json.Unmarshal(queryResponse.Value, &emr)
		if err != nil {
			return nil, err
		}
		emrs = append(emrs, &emr)
	}
	return emrs, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&EmrContract{})
	if err != nil {
		log.Panicf("Error creating project-manage chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting project-manage chaincode: %v", err)
	}
}
