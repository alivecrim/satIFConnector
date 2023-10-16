import {Component, OnInit} from '@angular/core';
import {InitSetupService, InitState} from "./service/init-setup.service";
import {Satellite} from "./entity/satellite";
import {GroundStation} from "./entity/groundStation";
import {DialogService, DynamicDialogRef} from "primeng/dynamicdialog";
import {ProjectCreateComponent} from "./dialogs/project/project-create/project-create.component";
import {ProjectSelectComponent} from "./dialogs/project/project-select/project-select.component";
import {
  GroundStationCreateComponent
} from "./dialogs/groundStation/ground-station-create/ground-station-create.component";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  providers: [DialogService]
})
export class AppComponent implements OnInit {
  initState?: InitState;
  private sat?: Satellite;
  private gs?: GroundStation;
  ref: DynamicDialogRef | undefined;

  constructor(private initService: InitSetupService, public dialogService: DialogService) {

  }

  ngOnInit(): void {
    this.loadState();
  }

  private loadState() {
    this.initService.initState().subscribe(s => {
        this.initState = s
        this.checkState()
      }
    )
  }

  checkState() {
    if (this.initState) {
      if (this.initState.projects.length == 0) {
        this.createProject()
        return
      }
      if (!this.initState.activeProject) {
        this.selectActiveProject()
        return
      }
      if (this.initState.groundStations.length == 0) {
        this.createGroundStation()
        return
      }
      if (this.initState.relatedGroundStations.length > 1) {
        this.selectGroundStation()
        return
      }
      if (this.initState.satellites.length == 0) {
        this.createSatellite()
        return;
      }
      if (!this.gs?.satelliteId) {
        this.createOrConnectSatellite()
      }
    }


  }

  createProject() {
    this.ref = this.dialogService.open(ProjectCreateComponent, {header: 'Create new projectStore'});
    this.ref.onClose.subscribe(r => {
      this.loadState()
    })
  }

  selectActiveProject() {
    this.ref = this.dialogService.open(ProjectSelectComponent, {header: 'Select active projectStore', data: this.initState});
    this.ref.onClose.subscribe(r => {
      this.loadState()
    })
  }

  createGroundStation() {
    this.ref = this.dialogService.open(GroundStationCreateComponent, {
      header: 'Create new GroundStation',
      data: this.initState
    });
    this.ref.onClose.subscribe(r => {
      this.loadState()
    })
  }

   selectGroundStation() {

  }

  createSatellite() {

  }

  private createOrConnectSatellite() {

  }
}
