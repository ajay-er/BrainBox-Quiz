import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/shared/data-access/auth.service';

@Component({
  selector: 'app-nav',
  templateUrl: './nav.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class NavComponent {
  protected auth = inject(AuthService);
  protected router = inject(Router);

  isDropDownOpen: boolean = false;

  toogleDrop() {
    this.isDropDownOpen = !this.isDropDownOpen;
  }

  logout(){
    this.auth.logout()
    this.router.navigateByUrl('/home')
  }
}
