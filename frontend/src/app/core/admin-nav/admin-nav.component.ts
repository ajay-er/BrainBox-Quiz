import { Component, inject } from '@angular/core';
import { AuthService } from 'src/app/shared/data-access/auth.service';

@Component({
  selector: 'app-admin-nav',
  templateUrl: './admin-nav.component.html',
  styleUrls: ['./admin-nav.component.css']
})
export class AdminNavComponent {
  protected authService = inject(AuthService)

}
