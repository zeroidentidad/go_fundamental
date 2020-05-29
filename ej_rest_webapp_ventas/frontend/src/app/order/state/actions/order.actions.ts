import {Action} from "@ngrx/store";
import {SearchOrderCriteria} from "../../models/order/search-order-criteria";


export enum OrderActionTypes {
    UpdateOrderSearchCriteria='[Order] Update Order Search Criteria',
}

export class UpdateOrderSearchCriteria implements Action {
    readonly type=OrderActionTypes.UpdateOrderSearchCriteria;
    constructor(public payload: SearchOrderCriteria) {}
}

export type Actions = UpdateOrderSearchCriteria;