import {Injectable} from "@angular/core";
import {Actions, Effect, ofType} from "@ngrx/effects";
import {OrderService} from "../../services/order.service";
import {Store} from "@ngrx/store";
import * as fromRoot from './../reducers';
import * as orderActions from '../actions/order.actions';
import {withLatestFrom, switchMap, map} from "rxjs/operators";
import * as moment from 'moment';
import {Router} from "@angular/router";

@Injectable()
export class OrderEffects {
    constructor(private actions$: Actions, private orderService: OrderService, private store: Store<fromRoot.OrderState>, private router: Router){

    }

    @Effect()
    getOrders$=this.actions$.pipe(
        ofType<orderActions.LoadOrders>(orderActions.OrderActionTypes.LoadOrders),
        withLatestFrom(this.store.select(fromRoot.getQuery)),
        switchMap(([action, query]) => {            
            let orderRequest={
                ...action.request,
                status: (query.status? query.status:null),
                dateFrom: query.dateFrom? moment(query.dateFrom).format("YYYY/MM/DD"):null,
                dateTo: query.dateTo? moment(query.dateTo).format("YYYY/MM/DD"):null
            }            

            return this.orderService.getOrders(orderRequest)
                .pipe(
                    map(data => new orderActions.LoadOrdersCompleted(data))
                )
        })
    );
    

    @Effect()
    addOrder$=this.actions$.pipe(
        ofType<orderActions.AddOrder>(orderActions.OrderActionTypes.AddOrder),
        switchMap(action => this.orderService.addOrder(action.request)
            .pipe(
                map(response => {
                    this.router.navigate(['/order']);
                    return new orderActions.AddOrderCompleted(response);
                })
            ))
    );
   
               
}