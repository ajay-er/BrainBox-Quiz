import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UpsertInputComponent } from './upsert-input.component';

describe('UpsertInputComponent', () => {
  let component: UpsertInputComponent;
  let fixture: ComponentFixture<UpsertInputComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [UpsertInputComponent]
    });
    fixture = TestBed.createComponent(UpsertInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
