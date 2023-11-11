import { Component, inject } from '@angular/core';
import { AuthService } from 'src/app/shared/data-access/auth.service';

@Component({
  selector: 'app-nav',
  templateUrl: './nav.component.html',
})
export class NavComponent {
  protected auth = inject(AuthService);
}
