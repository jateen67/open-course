export class Order {
  // public id: number;
  public email: string;
  public phone: string;
  public classNumber: number;
  // public isActive: boolean;
  // public createdAt: Date;
  // public updatedAt: Date;

  constructor(other?: Partial<Order>) {
    // this.id = other?.id || 0;
    this.email = other?.email || "";
    this.phone = other?.phone || "";
    this.classNumber = other?.classNumber || 0;
    // this.isActive = other?.isActive || true;
    // this.createdAt = other?.createdAt || new Date();
    // this.updatedAt = other?.updatedAt || new Date();
  }
}
