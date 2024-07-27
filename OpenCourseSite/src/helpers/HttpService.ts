import { catchError, Observable, switchMap } from "rxjs";
import { fromFetch } from "rxjs/fetch";

export const Fetch = <T>(
  route: string,
  params: RequestInit,
  responseType: "json" | "text" = "json"
): Observable<T> => {
  const data = fromFetch(route, params).pipe(
    switchMap((response) => {
      if (!response.ok) throw response.status;

      return (
        responseType === "json" ? response.json() : response.text()
      ) as Promise<T>;
    }),
    catchError((error) => {
      throw error;
    })
  );
  return data;
};
