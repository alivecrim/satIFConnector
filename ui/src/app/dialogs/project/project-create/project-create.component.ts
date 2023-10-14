import {Component} from '@angular/core';
import {DynamicDialogConfig, DynamicDialogRef} from "primeng/dynamicdialog";
import {FormBuilder, FormGroup} from "@angular/forms";
import {ProjectService} from "../../../service/project.service";
import {ProjectImpl} from "../../../entity/project";

@Component({
  selector: 'app-project-create',
  templateUrl: './project-create.component.html',
  styleUrls: ['./project-create.component.scss']
})
export class ProjectCreateComponent {
  formGroup: FormGroup;
  categories: any[] = [
    { name: 'True', key: true },
    { name: 'False', key: false },
  ];

  constructor(
    public ref: DynamicDialogRef,
    public config: DynamicDialogConfig,
    private fb: FormBuilder,
    private projectService: ProjectService
  ) {
    this.formGroup = fb.group({
      nameForm: [''],
      descriptionForm: [''],
      activeForm: ['']
    })
  }

  create() {
    this.projectService.create(new ProjectImpl(
      this.formGroup.get("nameForm")?.value,
      this.formGroup.get("descriptionForm")?.value,
      this.formGroup.get("activeForm")?.value["key"])).subscribe(r => {
      this.ref.close({closeCode: 1})
    })

  }
}
