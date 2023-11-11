import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { AuthService } from 'src/app/shared/data-access/auth.service';

@Component({
  selector: 'app-nav',
  templateUrl: './nav.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class NavComponent {
  protected auth = inject(AuthService);
 
}
