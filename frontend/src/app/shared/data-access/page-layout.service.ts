import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { PageLayout } from '../interfaces';

@Injectable({
  providedIn: 'root',
})
export class PageLayoutService {
  private layoutSubject = new BehaviorSubject<PageLayout>(PageLayout.User);

  public layout$ = this.layoutSubject.asObservable();

  setLayout(value: PageLayout) {
    this.layoutSubject.next(value);
  }
}
