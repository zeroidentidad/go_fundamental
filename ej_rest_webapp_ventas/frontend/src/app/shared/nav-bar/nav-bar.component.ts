import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.scss']
})
export class NavBarComponent implements OnInit {

  @Input()
  isAuthenticated: boolean;
  @Output()
  login: EventEmitter<any>=new EventEmitter<any>();
  @Output()
  logout: EventEmitter<any>=new EventEmitter<any>();

  constructor() { }

  ngOnInit() {
  }

  onLogin() {
  }

  onLogout() {
  }

}
