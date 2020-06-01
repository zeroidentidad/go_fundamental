import {Injectable} from '@angular/core';
import {CustomerService} from '../../services/customer.service';
import {Actions, Effect, ofType} from '@ngrx/effects';
import * as customerActions from '../actions/customer.actions';
import {switchMap, map} from 'rxjs/operators';

@Injectable()
export class CustomerEffects {
    constructor(private actions$: Actions, private customerService: CustomerService) {

    }

    @Effect()
    getCustomers$=this.actions$.pipe(
        ofType<customerActions.LoadCustomers>(customerActions.CustomerActionTypes.LoadCustomers),
        switchMap(action => this.customerService.getCustomers(action.request)
            .pipe(
                map(data => new customerActions.LoadCustomersCompleted(data))
            ))
    );
}