import { catchError, Observable, switchMap, throwError } from "rxjs";
import { fromFetch } from "rxjs/fetch";

export const Fetch = <T>(
  route: string,
  params: RequestInit,
  responseType: "json" | "text" = "json"
): Observable<T> => {
  return fromFetch(route, params).pipe(
    switchMap((response) => {
      if (!response.ok) {
        // Parse the response body even if the status code is not OK (e.g., 400)
        return response.json().then((error) => {
          // Throw a new error with the extracted message
          throw { status: response.status, message: error.message || 'Unknown error' };
        });
      }

      return (
        responseType === "json" ? response.json() : response.text()
      ) as Promise<T>;
    }),
    catchError((error) => {
      // Ensure the error is passed down to the subscriber
      return throwError(() => error);
    })
  );
};