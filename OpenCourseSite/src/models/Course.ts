export class Course {
  public id: number;
  public courseCode: string;
  public courseTitle: string;
  public semester: string;
  public section: string;
  public credits: number;
  public openSeats: number;
  public waitlistAvailable: number;
  public waitlistCapacity: number;
  public createdAt: Date;
  public updatedAt: Date;

  constructor(other?: Partial<Course>) {
    this.id = other?.id || 0;
    this.courseCode = other?.courseCode || "";
    this.courseTitle = other?.courseTitle || "";
    this.semester = other?.semester || "";
    this.section = other?.section || "";
    this.credits = other?.credits || 0;
    this.openSeats = other?.openSeats || 0;
    this.waitlistAvailable = other?.waitlistAvailable || 0;
    this.waitlistCapacity = other?.waitlistCapacity || 0;
    this.createdAt = other?.createdAt || new Date();
    this.updatedAt = other?.updatedAt || new Date();
  }
}
