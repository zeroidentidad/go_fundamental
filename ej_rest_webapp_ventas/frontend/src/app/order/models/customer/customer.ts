export class Customer {
    constructor(public id: number,
        public name: string,
        public address: string,
        public businessphone: string,
        public city: string,
        public company: string) {

    }

    static EmptyInstance() {
        return new Customer(null, null, null, null, null, null);
    }

    static mapFromResponse(data: any) {
        return new Customer(data.id,
            `${data.firstname},${data.lastname}`,
            data.address,
            data.businessphone,
            data.city,
            data.company);
    }
}