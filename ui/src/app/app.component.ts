import {Component, OnInit} from '@angular/core';
import {InitSetupService, InitState} from "./service/init-setup.service";
import {Satellite} from "./entity/satellite";
import {GroundStation} from "./entity/groundStation";
import {DialogService, DynamicDialogRef} from "primeng/dynamicdialog";
import {ProjectCreateComponent} from "./dialogs/project/project-create/project-create.component";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  providers: [ DialogService]
})
export class AppComponent implements OnInit {
  private initState?: InitState;
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
      if (this.initState.projects.length==0) {
        this.createProject()
        return
      }
      if (this.initState.activeProject==undefined) {
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

  private createProject() {
    this.ref = this.dialogService.open(ProjectCreateComponent, {header: 'Create new project'});
    this.ref.onClose.subscribe(r => {
      this.loadState()
    })
  }

  private selectActiveProject() {

  }

  private createGroundStation() {

  }

  private selectGroundStation() {

  }

  private createSatellite() {

  }

  private createOrConnectSatellite() {

  }
}
