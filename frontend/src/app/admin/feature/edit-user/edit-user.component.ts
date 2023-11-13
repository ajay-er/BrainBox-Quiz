import { Component, inject } from '@angular/core';
import { ISignup } from 'src/app/shared/interfaces';
import { AdminService } from '../../data-access/admin.service';

@Component({
  selector: 'app-edit-user',
  templateUrl: './edit-user.component.html',
  styleUrls: ['./edit-user.component.css'],
})
export class EditUserComponent {
  private adminService = inject(AdminService);
  editUserSubmission(userdata: ISignup) {
    this.adminService.updateUserData(userdata).subscribe((res) => {
      console.log(res);
    });
  }
}
