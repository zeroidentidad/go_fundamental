import {PreOrderProduct} from './pre-order-product';

export class PreOrder {
    Id: Number;
    CustomerId: number;
    OrderDate: String;
    OrderDetails: PreOrderDetail[]=[];

    static createEmptyInstance() {
        return new PreOrder();
    }
}

export class PreOrderDetail {
    Id: Number;
    OrderId: Number;
    ProductId: Number;
    Quantity: number;
    UnitPrice: Number;

    static mapOrderDetail(products: PreOrderProduct[]) {
        let newProducts: PreOrderDetail[]=[];
        products.map((item: PreOrderProduct) => {
            const product: PreOrderDetail=new PreOrderDetail();
            product.Id=item.Id;
            product.ProductId=item.ProductId;
            product.Quantity=item.Quantity;
            product.UnitPrice=item.UnitPrice;
            newProducts.push(product);
        });
        return newProducts;
    }
}