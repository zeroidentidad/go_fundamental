import { Component, OnInit, Input, SimpleChanges, OnChanges } from '@angular/core';
import {ProductBestSeller} from "../../models/product/best-seller";
import * as Highcharts from "highcharts";

@Component({
  selector: 'app-best-seller-product-charts',
  templateUrl: './best-seller-product-charts.component.html',
  styleUrls: ['./best-seller-product-charts.component.scss']
})
export class BestSellerProductChartsComponent implements OnInit, OnChanges {

  @Input()
  products: ProductBestSeller[];
  Highcharts: any;
  chartOptions: any;

  constructor() { }

  ngOnChanges(changes: SimpleChanges): void {
    if(changes.products&&changes.products.currentValue) {
      this.buildChart();
    }
  }  

  ngOnInit() {
  }

  buildChart(): void {
    this.Highcharts=Highcharts;
    this.chartOptions={
      title: {
        text: 'Los 10 productos que han generado más ingresos'
      },
      xAxis: {
        type: 'category'
      },
      yAxis: {
        title: {
          text: '% del Total de Ingresos'
        }
      },
      tooltip: {
        pointFormat: '<span style="color:{point.color}">{point.name}</span>: <b>{point.y:.2f}%</b> del total<br/>'
      },
      series: [
        {
          name: 'Productos',
          data: this.products,
          type: 'column',
          colorByPoint: true
        }
      ],

      chart: {
        zoomType: 'xy'
      },      

      responsive: {
        rules: [{
          condition: {
            maxWidth: '100%'
          },
          chartOptions: {
            legend: {
              align: 'center',
              verticalAlign: 'bottom',
              layout: 'horizontal'
            },
            yAxis: {
              labels: {
                align: 'left',
                x: 0,
                y: -5
              },
              title: {
                text: 'Los 10 productos que han generado más ingresos'
              }
            },
            subtitle: {
              text: null
            },
            credits: {
              enabled: false
            }
          }
        }]
      }, 

    }
  }  

}
