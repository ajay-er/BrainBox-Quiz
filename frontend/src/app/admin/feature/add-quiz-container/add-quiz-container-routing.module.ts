import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AddQuizContainerComponent } from './add-quiz-container.component';

const routes: Routes = [
  {
    path: '',
    component: AddQuizContainerComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class AddQuizContainerRoutingModule {}
