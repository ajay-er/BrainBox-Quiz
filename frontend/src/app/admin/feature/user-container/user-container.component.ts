import { Component } from '@angular/core';
import { initFlowbite } from 'flowbite';

@Component({
  selector: 'app-user-container',
  templateUrl: './user-container.component.html',
  styleUrls: ['./user-container.component.css']
})
export class UserContainerComponent {
  ngOnInit(){
    initFlowbite()
  }
}
