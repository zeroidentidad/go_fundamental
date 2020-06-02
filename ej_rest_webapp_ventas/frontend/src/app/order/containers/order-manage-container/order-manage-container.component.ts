import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {MatDialog} from "@angular/material";
import {CustomerPopupContainerComponent} from "../customer-popup-container/customer-popup-container.component";
import {Customer} from "../../models/customer/customer";
import {ProductPopupContainerComponent} from "../product-popup-container/product-popup-container.component";
import {Product} from "../../models/product/product";
import {PreOrderProduct} from "../../models/pre-order/pre-order-product";
import {PreOrderFooter} from "../../models/pre-order/pre-order-footer";
import {PreOrder, PreOrderDetail} from "../../models/pre-order/pre-order";
import * as moment from "moment";
import {Store, ActionsSubject} from "@ngrx/store";
import * as orderActions from '../../state/actions/order.actions';
import * as fromReducer from '../../state/reducers';
import {Router, ActivatedRoute} from "@angular/router";
import {ofType} from "@ngrx/effects";
import {OrderListItem} from "../../models/order/order-list-item";

@Component({
  selector: 'app-order-manage-container',
  templateUrl: './order-manage-container.component.html',
  styleUrls: ['./order-manage-container.component.scss']
})
export class OrderManageContainerComponent implements OnInit {

  orderId: number=0;
  orderForm: FormGroup;
  orderProductList: PreOrderProduct[]=[];
  preOrderFooter: PreOrderFooter=PreOrderFooter.createEmptyInstance();
  preOrder: PreOrder = PreOrder.createEmptyInstance();
  constructor(private fb: FormBuilder, private dialog: MatDialog, private store: Store<fromReducer.OrderState>, private router: Router, private activateRouter: ActivatedRoute, private actionsSubject$: ActionsSubject,) { 
    this.buildNewForm();
    this.triggers();
  }

  ngOnInit() {
    this.activateRouter.params
      .subscribe((params: any) => {
        if(params.id) {
          this.store.dispatch(new orderActions.LoadOrderById(params.id));
        }
      })
  }

  triggers() {
    this.actionsSubject$
      .pipe(ofType(orderActions.OrderActionTypes.LoadOrderByIdCompleted))
      .subscribe((response: any) => {
        this.buildEditForm(response.payload)
      });
  }


  buildNewForm() {
    this.orderForm=this.fb.group({
      id: ['', [Validators.required]],
      name: ['', [Validators.required]],
      company: ['', [Validators.required]],
      fecha: [new Date(), [Validators.required]],
      address: ['', [Validators.required]],
      phone: ['', [Validators.required]],
      city: ['', [Validators.required]]
    });
  }

  buildEditForm(orderItem: OrderListItem): void {
    this.orderId=orderItem.orderId;
    this.orderForm=this.fb.group({
      id: [orderItem.customerId, [Validators.required]],
      name: [orderItem.customer, [Validators.required]],
      company: [orderItem.company, [Validators.required]],
      fecha: [new Date(orderItem.orderDate), [Validators.required]],
      address: [orderItem.address, [Validators.required]],
      phone: [orderItem.phone, [Validators.required]],
      city: [orderItem.city, [Validators.required]]
    });

    const orderDetailEdit=orderItem.data.data;
    if(orderDetailEdit) {
      for(let index=0;index<orderDetailEdit.length;index++) {
        const element=orderDetailEdit[index];
        const product=new PreOrderProduct(element.id, element.product_id, element.description,
          element.product_name, element.unit_price, element.quantity);
        this.orderProductList.push(product);
        this.orderProductList=[...this.orderProductList];
        this.preOrderFooter=new PreOrderFooter(this.orderProductList);
      }
    }
  }

  openCustomerPopup(){
    const dialogRef = this.dialog.open(CustomerPopupContainerComponent, {
      panelClass: ''
    })
    
    dialogRef.afterClosed().subscribe((response: Customer)=>{
      if (response) {
        this.orderForm.patchValue({
          id: response.id,
          name: response.name,
          company: response.company,
          address: response.address,
          phone: response.businessphone,
          city: response.city,
        });        
      }
    })
  }

  openProductPopup(){
    const dialogRef=this.dialog.open(ProductPopupContainerComponent, {
      panelClass: ''
    });
    dialogRef.afterClosed()
      .subscribe((response: Product) => {
        this.AddProductToList(response);
      });
  }

  AddProductToList(item: Product): void {
    const product=new PreOrderProduct(0, item.id, item.description, item.productName, item.standardCost, 1);
    this.orderProductList.push(product);
    this.orderProductList=[...this.orderProductList];

    this.preOrderFooter=new PreOrderFooter(this.orderProductList);
  }  

  UpdateQuantity(event: any): any {
    const productToUpdate={...this.orderProductList[event.index]};
    productToUpdate.Quantity=Number(event.newValue);
    productToUpdate.TotalValue=Number((productToUpdate.Quantity*productToUpdate.UnitPrice).toFixed(2));

    this.orderProductList[event.index]=productToUpdate;
    this.orderProductList=[...this.orderProductList];

    this.preOrderFooter=new PreOrderFooter(this.orderProductList);
  }

  onDeleteProductOrder(orderDetailId: any) {

  }  

  onSave(){
    const customer=this.orderForm.value; // props: id, fecha
    this.preOrder.Id=this.orderId;
    this.preOrder.CustomerId=customer.id;
    this.preOrder.OrderDate=moment(customer.fecha).format("YYYY/MM/DD");
    this.preOrder.OrderDetails=PreOrderDetail.mapOrderDetail(this.orderProductList);

    if(this.orderId===0) {
      this.store.dispatch(new orderActions.AddOrder(this.preOrder));
    } else {
      //Actualizar
    }

  }

  onCancel(): void {
    this.router.navigate(['order']);
  }

}
