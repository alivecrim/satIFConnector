import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {GroundStation} from "../entity/groundStation";

@Injectable({
  providedIn: 'root'
})
export class GroundStationService {

  private api = "/api/crud/groundStation"

  constructor(private http: HttpClient) {
  }

  findAll(): Observable<GroundStation[]> {
    return this.http.get<GroundStation[]>(this.api)
  }

  create(p: GroundStation): Observable<GroundStation> {
    return this.http.post<GroundStation>(this.api, p)
  }

  delete(id: string): Observable<any> {
    return this.http.delete<any>(this.api + `/${id}`)
  }
}
