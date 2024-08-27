import { Fetch, GlobalEnvironment, postParams } from "helpers";
import { Order } from "models";
import { Observable } from "rxjs";

export class OrderService {
  static CreateOrder(order: Order): Observable<Order> {
    return Fetch<Order>(
      `${GlobalEnvironment.GetUrlApi()}/orders`,
      postParams(JSON.stringify(order))
    );
  }
}
