import {OrderListDetail} from './order-list-detail';

export class OrderListItem {
    orderId: number;
    address: string;
    city: string;
    company: string;
    customer: string;
    customerId: number;
    data: OrderListDetail; // resumen venta (Total, Impuestos, Importe General)
    orderDate: string;
    phone: string;
    statusId: string;
    statusName: string;

    static mapFromResponse(data: any): OrderListItem {
        const orderListItem: OrderListItem=new OrderListItem();
        orderListItem.orderId=data.orderId;
        orderListItem.customer=data.customer;
        orderListItem.customerId=data.customerId;
        orderListItem.customer=data.customer;
        orderListItem.company=data.company;
        orderListItem.address=data.address;
        orderListItem.city=data.city;
        orderListItem.phone=data.phone;
        orderListItem.orderDate=data.orderDate;
        orderListItem.statusId=data.statusId;
        orderListItem.statusName=data.statusName;
        orderListItem.data=OrderListDetail.mapFromResponse(data.data);
        return orderListItem;
    }
}