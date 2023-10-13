export interface Project {
  id?: string
  name: string
  description: string
  active: boolean
}

export class ProjectImpl implements Project {
  active: boolean;
  description: string;
  name: string;

  constructor(name: string, description: string, active: boolean) {
    this.active = active
    this.name = name
    this.description = description
  }


}
