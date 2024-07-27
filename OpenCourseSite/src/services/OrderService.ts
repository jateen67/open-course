import {
  Fetch,
  getParams,
  GlobalEnvironment,
  postParams,
  putParams,
} from "helpers";
import { Order } from "models";
import { Observable } from "rxjs";

export class OrderService {
  static GetAll(): Observable<Order[]> {
    return Fetch<Order[]>(`${GlobalEnvironment.GetUrlApi()}/orders`, getParams);
  }

  static GetOrderById(id: number): Observable<Order> {
    return Fetch<Order>(
      `${GlobalEnvironment.GetUrlApi()}/orderbyid/${id}`,
      getParams
    );
  }

  static GetOrderByEmail(email: string): Observable<Order[]> {
    return Fetch<Order[]>(
      `${GlobalEnvironment.GetUrlApi()}/ordersbyemail/${email}`,
      getParams
    );
  }

  static GetOrderByCourseId(courseId: number): Observable<Order[]> {
    return Fetch<Order[]>(
      `${GlobalEnvironment.GetUrlApi()}/ordersbycourseid/${courseId}`,
      getParams
    );
  }

  static UpodateOrder(order: Order): Observable<Order> {
    return Fetch<Order>(
      `${GlobalEnvironment.GetUrlApi()}/orders`,
      putParams(JSON.stringify(order))
    );
  }

  static CreateOrder(order: Order): Observable<Order> {
    return Fetch<Order>(
      `${GlobalEnvironment.GetUrlApi()}/orders`,
      postParams(JSON.stringify(order))
    );
  }
}
