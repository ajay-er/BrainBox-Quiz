import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UserContainerRoutingModule } from './user-container-routing.module';
import { UserContainerComponent } from './user-container.component';
import { TableModule } from 'src/app/shared/ui/table/table.module';

@NgModule({
  declarations: [UserContainerComponent],
  imports: [CommonModule, UserContainerRoutingModule, TableModule],
})
export class UserContainerModule {}
