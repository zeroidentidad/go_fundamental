import { Component, OnInit, Input, OnChanges, SimpleChanges } from '@angular/core';

@Component({
  selector: 'app-table-view',
  templateUrl: './table-view.component.html',
  styleUrls: ['./table-view.component.scss']
})
export class TableViewComponent implements OnInit, OnChanges {

  @Input() items: any[]=[];
  @Input() columns: any[]=[];
  @Input() minTableHeight: number=500;

  tableColumns: any[]=[];
  constructor() { }

  ngOnChanges(changes: SimpleChanges): void {
    this.tableColumns=this.columns;
  }

  ngOnInit() {
    this.tableColumns=this.columns;
  }

}
