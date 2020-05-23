import { Component, OnInit, Inject } from '@angular/core';
import {MatDialogRef, MAT_DIALOG_DATA} from '@angular/material';
import {ConfirmData} from '../../models/confirm-data';

@Component({
  selector: 'app-app-confirm',
  templateUrl: './app-confirm.component.html',
  styleUrls: ['./app-confirm.component.scss']
})
export class AppConfirmComponent implements OnInit {

  constructor(public dialogRef: MatDialogRef<AppConfirmComponent>,
    @Inject(MAT_DIALOG_DATA) public data: ConfirmData) {}

  ngOnInit() {
  }

}
