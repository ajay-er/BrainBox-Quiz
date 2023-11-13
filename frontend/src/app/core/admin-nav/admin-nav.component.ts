import { Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import {  initFlowbite } from 'flowbite';
import { AuthService } from 'src/app/shared/data-access/auth.service';
import { PageLayoutService } from 'src/app/shared/data-access/page-layout.service';
import { PageLayout } from 'src/app/shared/interfaces';

@Component({
  selector: 'app-admin-nav',
  templateUrl: './admin-nav.component.html',
  styleUrls: ['./admin-nav.component.css'],
})
export class AdminNavComponent {
  protected authService = inject(AuthService);
  protected router = inject(Router);
  protected pageLayoutService = inject(PageLayoutService);

  isDropDownOpen: boolean = false;

  toogleDrop() {
    this.isDropDownOpen = !this.isDropDownOpen;
  }

  logOut() {
    this.pageLayoutService.setLayout(PageLayout.User);
    this.authService.adminLogout();
    this.router.navigateByUrl('/auth/admin-login');
  }

  ngOnInit(){
    initFlowbite();
  }


}
