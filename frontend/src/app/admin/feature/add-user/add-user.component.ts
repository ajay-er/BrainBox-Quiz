import { Component, inject } from '@angular/core';
import { ISignup } from 'src/app/shared/interfaces';
import { AdminService } from '../../data-access/admin.service';

@Component({
  selector: 'app-add-user',
  templateUrl: './add-user.component.html',
  styleUrls: ['./add-user.component.css'],
})
export class AddUserComponent {
  private adminService = inject(AdminService);

  addNewUserSubmission(userdata: ISignup) {
    this.adminService.addNewUser(userdata).subscribe((res) => {
      console.log(res);
    });
  }
}
