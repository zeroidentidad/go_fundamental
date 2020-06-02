import {PreOrderProduct} from './pre-order-product';

export class PreOrderFooter {
    TotalValorVenta: number;
    ImpuestoGeneralVentas: number;
    ImporteTotal: number;

    constructor(orderDetailList: PreOrderProduct[]) {
        this.TotalValorVenta=Number(orderDetailList.reduce((sum, item) => sum+item.TotalValue, 0).toFixed(2));
        this.ImpuestoGeneralVentas=Number((0.16*this.TotalValorVenta).toFixed(2));
        this.ImporteTotal=this.TotalValorVenta+this.ImpuestoGeneralVentas;
    }

    static createEmptyInstance(): PreOrderFooter {
        return new PreOrderFooter([]);
    }
}