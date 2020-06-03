import { Component, OnInit, TemplateRef, ViewChild, AfterViewInit, ChangeDetectorRef } from '@angular/core';
import {Store, ActionsSubject} from '@ngrx/store';
import * as orderActions from '../../state/actions/order.actions';
import * as fromReducer from '../../state/reducers';
import {GetOrder} from "../../models/order/get-order";
import {Observable, Subject} from "rxjs";
import {OrderListItem} from "../../models/order/order-list-item";
import {TableViewComponent} from "src/app/shared/components/table-view/table-view.component";
import {FormGroup, FormBuilder} from "@angular/forms";
import {SearchOrderCriteria} from "../../models/order/search-order-criteria";
import {Status} from "../../models/order/status";
import {Router} from "@angular/router";
import {ConfirmData} from "src/app/shared/models/confirm-data";
import {AppConfirmService} from "src/app/shared/components/app-confirm/app-confirm.service";
import {takeUntil} from "rxjs/operators";
import {ofType} from "@ngrx/effects";

@Component({
  selector: 'app-order-main-container',
  templateUrl: './order-main-container.component.html',
  styleUrls: ['./order-main-container.component.scss']
})
export class OrderMainContainerComponent implements OnInit, AfterViewInit {

  orders$: Observable<OrderListItem[]>=this.store.select(fromReducer.getOrders);
  length$: Observable<number>=this.store.select(fromReducer.getTotalRecords);  

  request: GetOrder;
  pageSize=10;
  pageSizeOptions: number[]=[10, 20, 50, 100];

  columns: object[]=[];
  detailColumns: object[]=[];

  @ViewChild("orderDateCellTemplate", {static: false})
  private orderDateCellTemplate: TemplateRef<any>;

  @ViewChild("accionesCellTemplate", {static: false})
  private accionesCellTemplate: TemplateRef<any>;

  @ViewChild("orderIdCellTemplate", {static: false})
  private orderIdCellTemplate: TemplateRef<any>;

  @ViewChild("tableView", {static: false})
  private tableView: TableViewComponent;

  searchForm: FormGroup;
  statusList: Status[]=this.getStatus()

  private ngUnsubscribe: Subject<any>=new Subject<any>();

  constructor(private store: Store<fromReducer.OrderState>, private fb: FormBuilder, private router: Router, private confirmService: AppConfirmService, private ref: ChangeDetectorRef, private actionSubject$: ActionsSubject) { 
    this.buildSearchForm();
    this.refreshData();
    this.triggers();
  }

  ngOnInit() {
  }

  ngAfterViewInit(): void {
    this.columns=this.getColumns();
    this.detailColumns=this.getDetailColums();
    this.ref.detectChanges();
  }

  triggers(): void {
    this.actionSubject$
      .pipe(takeUntil(this.ngUnsubscribe))
      .pipe(ofType(orderActions.OrderActionTypes.DeleteOrderCompleted))
      .subscribe(_ => {
        this.refreshData();
      })
  }

  refreshData(): void {
    this.request=new GetOrder(this.pageSizeOptions[0], 0);
    this.store.dispatch(new orderActions.LoadOrders(this.request));
  }

  getColumns(): object[] {
    return [
      {
        name: "Id",
        cellTemplate: this.orderIdCellTemplate,
        flexGrow: 1
      },
      {
        name: "Cliente",
        prop: "customer",
        flexGrow: 1
      },
      {
        name: "Telefono",
        prop: "phone",
        flexGrow: 1
      },
      {
        name: "Dirección",
        prop: "address",
        flexGrow: 1
      },
      {
        name: "Ciudad",
        prop: "city",
        flexGrow: 1
      },
      {
        name: "Fecha",
        cellTemplate: this.orderDateCellTemplate,
        flexGrow: 1
      },
      {
        name: "Estado",
        prop: "statusName",
        flexGrow: 1
      },
      {
        name: "Acciones",
        cellTemplate: this.accionesCellTemplate,
        flexGrow: 1
      }
    ];
  }  

  toggleExpandRow(row: any) {
    this.tableView.toogleExpandRow(row);
  }

  getDetailColums(): object[] {
    return [
      {
        name: "Nombre producto",
        prop: "product_name",
        flexGrow: 1
      },
      {
        name: "Cantidad",
        prop: "quantity",
        flexGrow: 1
      },
      {
        name: "Precio unitario",
        prop: "unit_price",
        flexGrow: 0.5
      }
    ];
  }

  changePage(event: any): void {
    const offSet=event.pageIndex*event.pageSize;
    this.pageSize=event.pageSize;
    this.request=new GetOrder(event.pageSize, offSet);
    this.store.dispatch(new orderActions.LoadOrders(this.request));
  }
  
  buildSearchForm(): void {
    this.searchForm=this.fb.group({
      dateFrom: ['', []],
      dateTo: ['', []],
      status: ['', []],
    });
  }

  search() {
    if(this.searchForm.valid) {
      if(this.searchForm.dirty) {
        let searchCriteria: SearchOrderCriteria;
        const newSearchCriteria={...searchCriteria, ...this.searchForm.value};
        this.store.dispatch(new orderActions.UpdateOrderSearchCriteria(newSearchCriteria));
        this.refreshData();
      }
    } else {
      console.log("nada ingreso algo...")
    }
  } 
  
  getStatus(): Status[] {
    return [
      Status.CreateInstance(0, "Nuevo"),
      Status.CreateInstance(1, "Pagado"),
      Status.CreateInstance(2, "Enviado"),
      Status.CreateInstance(3, "Cerrado")
    ];
  }  

  onViewDetail(orderId: number) {
    this.router.navigate(['order/detail/', orderId])
  }

  onRemoveOrder(orderId: number) {
    const confirmData=new ConfirmData('Eliminar orden', 'Estas seguro de eliminar la orden?', true);
    this.confirmService.confirm(confirmData)
      .subscribe(result => {
        if(result) {
          this.store.dispatch(new orderActions.DeleteOrder(orderId));
        }
      })
  }  

}
