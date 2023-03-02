'use strict';

const { WorkloadModuleBase } = require('@hyperledger/caliper-core');



const emrName = ['zhangsan','lisi','wangwu','zhaoliu','chenqi', 'jimson','michael'] 
const patientID = ['111111111111111111','222222222222222','33333333333333','44444444444444444','5555555555555','6666666666666','77777777777777','8888888888888','9999999999999', '10101010101010'] 
const emrGeneHospitalID= ['123123123123','234234234234','345345345345','456456456456','567567567567','678678678678','789789789789','890890890890' ]
const emrGeneDoctorID= ['123123123123','234234234234','345345345345','456456456456','567567567567','678678678678','789789789789','890890890890']
const emrIPFSHash=['QmWinSCXWQDArYYrw9pH9b8RCtE1sKffvopgk1Q6kE1SKK','QmWinSCXWQDArYYrw9pH9b8RCtE1sKffvopgk1Q6kE1SK1','QmWinSCXWQDArYYrw9pH9b8RCtE1sKffvopgk1Q6kE1SK2','QmWinSCXWQDArYYrw9pH9b8RCtE1sKffvopgk1Q6kE1SK3', 'QmWinSCXWQDArYYrw9pH9b8RCtE1sKffvopgk1Q6kE1SK4', 'QmWinSCXWQDArYYrw9pH9b8RCtE1sKffvopgk1Q6kE1SK5', 'QmWinSCXWQDArYYrw9pH9b8RCtE1sKffvopgk1Q6kE1SK6', 'QmWinSCXWQDArYYrw9pH9b8RCtE1sKffvopgk1Q6kE1SK7', 'QmWinSCXWQDArYYrw9pH9b8RCtE1sKffvopgk1Q6kE1SK8', 'QmWinSCXWQDArYYrw9pH9b8RCtE1sKffvopgk1Q6kE1SK9']


class CreateEmrWorkload extends WorkloadModuleBase {
    constructor() {
        super();
        this.txIndex = 0;
    }

    async submitTransaction() {
        this.txIndex++;
        let emrID = 'patient' + this.workerIndex + '_Emr' + this.txIndex.toString();
        let agriEmrName = emrName[Math.floor(Math.random() * emrName.length)];
        let agriPatientID = patientID[Math.floor(Math.random() * patientID.length)];
        let agriEmrGeneHospitalID = emrGeneHospitalID[Math.floor(Math.random() * emrGeneHospitalID.length)];
        let agriGeneDoctorID = emrGeneDoctorID[Math.floor(Math.random() * emrGeneDoctorID.length)];
        let agriEmrIPFSHash = emrIPFSHash[Math.floor(Math.random() * emrIPFSHash.length)];

        let args = {
            contractId: 'emr',
            contractVersion: '1.0',
            contractFunction: 'CreateEmr',
            contractArguments: [emrID, agriEmrName, agriPatientID, agriEmrGeneHospitalID ,agriGeneDoctorID, agriEmrIPFSHash],
            timeout: 30
        };

        await this.sutAdapter.sendRequests(args);
    }
}


function createWorkloadModule() {
    return new CreateEmrWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;


