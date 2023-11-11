import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AddUserRoutingModule } from './add-user-routing.module';
import { AddUserComponent } from './add-user.component';
import { UpsertInputModule } from '../../ui/upsert-input/upsert-input.module';

@NgModule({
  declarations: [AddUserComponent],
  imports: [CommonModule, AddUserRoutingModule, UpsertInputModule],
})
export class AddUserModule {}
