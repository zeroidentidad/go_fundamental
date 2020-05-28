export class GetOrder {
    dateFrom: string;
    dateTo: string;
    status: number;
    limit: number;
    offset: number;

    constructor(limit: number, offset: number) {
        this.limit=limit;
        this.offset=offset;
    }
}