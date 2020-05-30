import { Component, OnInit, Input, OnChanges, SimpleChanges, TemplateRef, ViewChild, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-table-view',
  templateUrl: './table-view.component.html',
  styleUrls: ['./table-view.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class TableViewComponent implements OnInit, OnChanges {

  @Input() items: any[]=[];
  @Input() columns: any[]=[];
  @Input() minTableHeight: number=500;
  @Input() detailTemplate: TemplateRef<any>;

  @ViewChild("myTable", {static: false})
  table: any;

  tableColumns: any[]=[];
  constructor() { }

  ngOnChanges(changes: SimpleChanges): void {
    this.tableColumns=this.columns;
  }

  ngOnInit() {
    this.tableColumns=this.columns;
  }

  toogleExpandRow(row: any) {
    this.table.rowDetail.toggleExpandRow(row);
  }

}
