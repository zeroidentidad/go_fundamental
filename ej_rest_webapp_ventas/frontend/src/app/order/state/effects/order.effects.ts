import {Injectable} from "@angular/core";
import {Actions, Effect, ofType} from "@ngrx/effects";
import {OrderService} from "../../services/order.service";
import {Store} from "@ngrx/store";
import * as fromRoot from './../reducers';
import * as orderActions from '../actions/order.actions';
import {withLatestFrom, switchMap, map} from "rxjs/operators";
import * as moment from 'moment';

@Injectable()
export class OrderEffects {
    constructor(private actions$: Actions, private orderService: OrderService, private store: Store<fromRoot.OrderState>,){

    }

    @Effect()
    getOrders$=this.actions$.pipe(
        ofType<orderActions.LoadOrders>(orderActions.OrderActionTypes.LoadOrders),
        withLatestFrom(this.store.select(fromRoot.getQuery)),
        switchMap(([action, query]) => {
            let orderRequest=action.request;
            /*orderRequest.status=(query.status ? query.status:null);
            orderRequest.dateFrom=query.dateFrom ? moment(query.dateFrom).format("YYYY/MM/DD"):null;
            orderRequest.dateTo=query.dateTo ? moment(query.dateTo).format("YYYY/MM/DD"):null;*/

            return this.orderService.getOrders(orderRequest)
                .pipe(
                    map(data => new orderActions.LoadOrdersCompleted(data))
                )
        })
    );
    
               
}