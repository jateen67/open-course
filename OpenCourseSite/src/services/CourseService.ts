import { Fetch, getParams, GlobalEnvironment } from "helpers";
import { Course } from "models";
import { Observable } from "rxjs";

export class CourseService {
  static GetAll(): Observable<Course[]> {
    return Fetch<Course[]>(
      `${GlobalEnvironment.GetUrlApi()}/courses`,
      getParams
    );
  }

  static GetByTermCodeAndCourseId(termCode: string, courseId: number): Observable<Course[]> {
    return Fetch<Course[]>(
      `${GlobalEnvironment.GetUrlApi()}/course/${termCode}/${courseId}`,
      getParams
    );
  }
}
