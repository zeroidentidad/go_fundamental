export class ProductBestSeller {
    constructor(
        public id: number,
        public name: string,
        public vendidos: number,
        public y: number
    ) {}

    static mapFromResponse(data: any, totalVendidos: number) {
        return new ProductBestSeller(
            data.id, data.product_name,
            data.vendidos, 
            this.getPorcentaje(totalVendidos, data.vendidos)
        );
    }

    static getPorcentaje(totalVendidos: number, vendidos: number) {
        return (vendidos/totalVendidos)*100;
    }
}