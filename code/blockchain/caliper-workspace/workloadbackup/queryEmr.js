'use strict';

const  { WorkloadModuleBase } = require('@hyperledger/caliper-core');

const helper = require('./helper');

class QueryEmrWorkload extends WorkloadModuleBase {
    constructor() {
        super();
        this.txIndex = 0;
        this.limitIndex = 0;
    }



    async initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext) {
        await super.initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext);

        this.limitIndex = this.roundArguments.assets;
        await helper.CreateEmr(this.sutAdapter, this.workerIndex, this.roundArguments);
    }



    async submitTransaction(){
        this.txIndex++;

        let emrID = '123123124123123';

        let args = {
            contractId: 'emr',
            contractVersion: '1.0',
            contractFunction: 'ReadEmr',
            contractArguments: [emrID],
            timeout: 30,
            readOnly: true
        };

        if (this.txIndex === this.limitIndex) {
            this.txIndex = 0;
        }

        await this.sutAdapter.sendRequests(args);
    }

}


function createWorkloadModule() {
    return new QueryEmrWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;