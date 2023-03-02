'use strict';

const  { WorkloadModuleBase } = require('@hyperledger/caliper-core');

const helper = require('./helper');

class DeleteEmrWorkload extends WorkloadModuleBase {
    constructor() {
        super();
        this.txIndex = 0;
        this.limitIndex = 0;
    }



    async initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext) {
        await super.initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext);

        this.limitIndex = this.roundArguments.assets;
        await helper.DeleteEmr(this.sutAdapter, this.workerIndex, this.roundArguments);
    }



    async submitTransaction(){
        this.txIndex++;

        let emrID = 'patient' + this.workerIndex + '_Emr' + this.txIndex.toString();

        let args = {
            contractId: 'emr',
            contractVersion: '1.0',
            contractFunction: 'DeleteEmr',
            contractArguments: [emrID],
            timeout: 30,
            readOnly: false
        };

        if (this.txIndex === this.limitIndex) {
            this.txIndex = 0;
        }

        await this.sutAdapter.sendRequests(args);
    }

   /*  async cleanupWorkloadModule() {
        for (let i=0; i<this.roundArguments.assets; i++) {
            const emrID = 'patient' + this.workerIndex + '_Emr' + this.txIndex.toString();
            console.log(`Worker ${this.workerIndex}: Deleting emrs ${emrID}`);
            const request = {
                contractId: 'emr',
                contractFunction: 'DeleteEmr',
                contractArguments: [emrID],
                readOnly: false
            };
            this.txIndex++;
            await this.sutAdapter.sendRequests(request);
        }
    } */
}


function createWorkloadModule() {
    return new DeleteEmrWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;