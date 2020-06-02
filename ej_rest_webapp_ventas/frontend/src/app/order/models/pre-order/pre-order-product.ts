export class PreOrderProduct {
    Id: number;
    ProductId: number;
    Description: string;
    Quantity: number;
    UnitPrice: number;
    TotalValue: number;
    Name: string;

    constructor(id: number,
        productId: number,
        description: string,
        name: string,
        unitPrice: number,
        quantity: number=1) {
        this.Id=id;
        this.ProductId=productId;
        this.Description=description;
        this.Name=name;
        this.Quantity=quantity;
        this.UnitPrice=Number(unitPrice.toFixed(2));
        this.TotalValue=Number((this.Quantity*unitPrice).toFixed(2));
    }
}