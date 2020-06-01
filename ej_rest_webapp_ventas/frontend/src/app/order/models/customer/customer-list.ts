import {Customer} from './customer';

export class CustomerList {
    data: Customer[];
    totalRecords: number;
    constructor(records: any, total: number) {
        this.data=records;
        this.totalRecords=total;
    }
}