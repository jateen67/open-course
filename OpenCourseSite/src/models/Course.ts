export class Course {
  public classNumber: number;
  public courseId: number;
  public termCode: number;
  public session: string;
  public subject: string;
  public catalog: string;
  public section: string;
  public componentCode: string;
  public courseTitle: string;
  public classStartTime: string;
  public classEndTime: string;
  public mondays: boolean;
  public tuesdays: boolean;
  public wednesdays: boolean;
  public thursdays: boolean;
  public fridays: boolean;
  public saturdays: boolean;
  public sundays: boolean;

  constructor(other?: Partial<Course>) {
    this.classNumber = other?.classNumber || 0;
    this.courseId = other?.courseId || 0;
    this.termCode = other?.termCode || 0;
    this.session = other?.session || "";
    this.subject = other?.subject || "";
    this.catalog = other?.catalog || "";
    this.section = other?.section || "";
    this.componentCode = other?.componentCode || "";
    this.courseTitle = other?.courseTitle || "";
    this.classStartTime = other?.classStartTime || "";
    this.classEndTime = other?.classEndTime || "";
    this.mondays = other?.mondays || false;
    this.tuesdays = other?.tuesdays || false;
    this.wednesdays = other?.wednesdays || false;
    this.thursdays = other?.thursdays || false;
    this.fridays = other?.fridays || false;
    this.saturdays = other?.saturdays || false;
    this.sundays = other?.sundays || false;
  }
}
