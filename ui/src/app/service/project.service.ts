import {Injectable} from '@angular/core';
import {Project} from "../entity/project";
import {Observable} from "rxjs";
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class ProjectService {

  private api = "/api/crud/project"

  constructor(private http: HttpClient) {
  }

  findAll(): Observable<Project[]> {
    return this.http.get<Project[]>(this.api)
  }

  create(p: Project): Observable<Project> {
    return this.http.post<Project>(this.api, p)
  }

  delete(id: string): Observable<any> {
    return this.http.delete<any>(this.api + `/${id}`)
  }

  setActive(id: string): Observable<any> {
    return this.http.put<any>(this.api + `/setActive/${id}`, null)
  }
}
