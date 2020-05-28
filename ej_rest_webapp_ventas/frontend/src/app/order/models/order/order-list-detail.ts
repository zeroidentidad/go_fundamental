import {OrderListDetailItem} from './order-list-detail-item';
import {OrderListItem} from './order-list-item';

export class OrderListDetail {
    data: OrderListDetailItem[];
    totalValorVenta: number;
    impuestoGeneralVentas: number;
    importeTotal: number;

    static mapFromResponse(data: any): OrderListDetail {
        if(!data) {
            return new OrderListDetail();
        }

        const orderDetail=new OrderListDetail();
        const products: OrderListDetailItem[]=[];

        for(let index=0;index<data.length;index++) {
            const product=new OrderListDetailItem();
            const element=data[index];
            product.id=element.id;
            product.product_id=element.product_id;
            product.product_name=element.product_name;
            product.quantity=element.quantity;
            product.unit_price=element.unit_price;
            products.push(product);
        }
        
        orderDetail.data=products;
        orderDetail.totalValorVenta=Number(products.reduce((sum, item) => sum+(item.quantity*item.unit_price), 0).toFixed(2));
        orderDetail.impuestoGeneralVentas=Number((0.18*orderDetail.totalValorVenta).toFixed(2));
        orderDetail.importeTotal=orderDetail.totalValorVenta+orderDetail.impuestoGeneralVentas;
        return orderDetail;
    }
}