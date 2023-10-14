import {Component, OnInit} from '@angular/core';
import {DynamicDialogConfig, DynamicDialogRef} from "primeng/dynamicdialog";
import {FormBuilder, FormGroup} from "@angular/forms";
import {ProjectService} from "../../../service/project.service";
import {InitState} from "../../../service/init-setup.service";

@Component({
  selector: 'app-project-select',
  templateUrl: './project-select.component.html',
  styleUrls: ['./project-select.component.scss']
})
export class ProjectSelectComponent implements OnInit {
  formGroup: FormGroup;
  state?: InitState
  constructor(
    public ref: DynamicDialogRef,
    protected fb: FormBuilder,
    public c: DynamicDialogConfig,
    private projectService: ProjectService
  ) {
    this.formGroup = fb.group({
      selectedProjectId: [''],
    })
  }

  ngOnInit(): void {
    this.state = this.c.data
  }

  apply() {
    let value = this.formGroup.get("selectedProjectId")?.value;
    if (value) {
      this.projectService.setActive(value).subscribe(() => {
        this.ref.close({closeCode: 1})
      })
    }
  }
}
