import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Project} from "../entity/project";
import {GroundStation} from "../entity/groundStation";
import {Satellite} from "../entity/satellite";
import {Observable} from "rxjs";

export interface InitState {
  projects: Project[]
  activeProject?: Project
  relatedGroundStations: GroundStation[]
  groundStations: GroundStation[]
  relatedSatellite?: Satellite
  satellites: Satellite[]
}

@Injectable({
  providedIn: 'root'
})
export class InitSetupService {
  private api = "/api/relationState"

  constructor(private http: HttpClient) {
  }

  initState(): Observable<InitState> {
    return this.http.get<InitState>(this.api)
  }
}
