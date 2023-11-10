import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { EditUserRoutingModule } from './edit-user-routing.module';
import { EditUserComponent } from './edit-user.component';
import { UpsertInputModule } from '../../ui/upsert-input/upsert-input.module';

@NgModule({
  declarations: [EditUserComponent],
  imports: [CommonModule, EditUserRoutingModule, UpsertInputModule],

})
export class EditUserModule {}
