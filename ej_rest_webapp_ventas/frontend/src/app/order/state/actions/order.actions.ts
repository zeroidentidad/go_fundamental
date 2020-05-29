import {Action} from "@ngrx/store";
import {SearchOrderCriteria} from "../../models/order/search-order-criteria";
import {GetOrder} from "../../models/order/get-order";
import {OrderList} from "../../models/order/order-list";


export enum OrderActionTypes {
    UpdateOrderSearchCriteria='[Order] Update Order Search Criteria',
    LoadOrders='[Order] Load Orders',
    LoadOrdersCompleted='[Order] Load Orders Completed',
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

export type Actions=UpdateOrderSearchCriteria|LoadOrders|LoadOrdersCompleted;