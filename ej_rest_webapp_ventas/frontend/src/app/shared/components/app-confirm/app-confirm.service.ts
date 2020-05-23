import { Injectable } from '@angular/core';
import {MatDialog} from '@angular/material';
import {ConfirmData} from '../../models/confirm-data';
import {Observable} from 'rxjs';
import {AppConfirmComponent} from './app-confirm.component';

@Injectable({
  providedIn: 'root'
})
export class AppConfirmService {

  constructor(private dialog: MatDialog) {}

  public confirm(data: ConfirmData={}): Observable<boolean> {
    const dialogRef=this.dialog.open(AppConfirmComponent, {
      width: '380px',
      data,
      disableClose: true
    });

    return dialogRef.afterClosed();
  }
}
