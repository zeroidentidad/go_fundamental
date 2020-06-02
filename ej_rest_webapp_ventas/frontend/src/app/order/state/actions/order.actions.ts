import {Action} from "@ngrx/store";
import {SearchOrderCriteria} from "../../models/order/search-order-criteria";
import {GetOrder} from "../../models/order/get-order";
import {OrderList} from "../../models/order/order-list";
import {PreOrder} from "../../models/pre-order/pre-order";
import {OrderListItem} from "../../models/order/order-list-item";


export enum OrderActionTypes {
    UpdateOrderSearchCriteria='[Order] Update Order Search Criteria',
    LoadOrders='[Order] Load Orders',
    LoadOrdersCompleted='[Order] Load Orders Completed',
    AddOrder='[Order] Add Order',
    AddOrderCompleted='[Order] Add Order Completed',
    LoadOrderById='[Order] Load Order By Id',
    LoadOrderByIdCompleted='[Order] Load Order By Id Completed',    
}

export class UpdateOrderSearchCriteria implements Action {
    readonly type=OrderActionTypes.UpdateOrderSearchCriteria;
    constructor(public payload: SearchOrderCriteria) {}
}

export class LoadOrders implements Action {
    readonly type=OrderActionTypes.LoadOrders;
    constructor(public request: GetOrder) {}
}

export class LoadOrdersCompleted implements Action {
    readonly type=OrderActionTypes.LoadOrdersCompleted;
    constructor(public payload: OrderList) {}
}

export class AddOrder implements Action {
    readonly type=OrderActionTypes.AddOrder;
    constructor(public request: PreOrder) {}
}

export class AddOrderCompleted implements Action {
    readonly type=OrderActionTypes.AddOrderCompleted;
    constructor(public orderId: number) {}
}

export class LoadOrderById implements Action {
    readonly type=OrderActionTypes.LoadOrderById;
    constructor(public orderId: number) {}
}

export class LoadOrderByIdCompleted implements Action {
    readonly type=OrderActionTypes.LoadOrderByIdCompleted;
    constructor(public payload: OrderListItem) {}
}

export type Actions=UpdateOrderSearchCriteria|LoadOrders|LoadOrdersCompleted|AddOrder|AddOrderCompleted|LoadOrderById|LoadOrderByIdCompleted;