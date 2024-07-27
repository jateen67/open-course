export class Order {
  public id: number;
  public name: string;
  public email: string;
  public phone: string;
  public courseId: number;
  public isActive: boolean;
  public createdAt: Date;
  public updatedAt: Date;

  constructor(other?: Partial<Order>) {
    this.id = other?.id || 0;
    this.name = other?.name || "";
    this.email = other?.email || "";
    this.phone = other?.phone || "";
    this.courseId = other?.courseId || 0;
    this.isActive = other?.isActive || true;
    this.createdAt = other?.createdAt || new Date();
    this.updatedAt = other?.updatedAt || new Date();
  }
}
