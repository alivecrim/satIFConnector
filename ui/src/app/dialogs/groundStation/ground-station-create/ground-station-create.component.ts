import {Component} from '@angular/core';
import {FormBuilder, FormGroup} from "@angular/forms";
import {DynamicDialogRef} from "primeng/dynamicdialog";
import {GroundStation} from "../../../entity/groundStation";
import {GroundStationService} from "../../../service/ground-station.service";

@Component({
  selector: 'app-ground-station-create',
  templateUrl: './ground-station-create.component.html',
  styleUrls: ['./ground-station-create.component.scss']
})
export class GroundStationCreateComponent {
  formGroup: FormGroup;


  constructor(
    public ref: DynamicDialogRef,
    private fb: FormBuilder,
    private groundStationService: GroundStationService
  ) {
    this.formGroup = fb.group({
      name: [''],
      description: [''],
    })
  }

  create() {
    let gs: GroundStation = {calibrationIds: [], groundConnectors: [], los: [], name: this.g("name"), description: this.g("description")}
    this.groundStationService.create(gs).subscribe(r => {
      this.ref.close({closeCode: 1})
    })

  }

  g(v: string) {
    return this.formGroup.get(v)?.value
  }

}

