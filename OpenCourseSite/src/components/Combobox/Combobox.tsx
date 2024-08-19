import { useState, useEffect } from "react";
import { Course } from "../../models";
import {
  Combobox,
  ComboboxInput,
  ComboboxButton,
  ComboboxOption,
  ComboboxOptions,
} from "@headlessui/react";
import { ChevronDownIcon } from "@heroicons/react/20/solid";
import ComboboxStyles from "./Combobox.module.css";
import { useFormContext } from "../../contexts";

interface CourseComboboxProps {
  onChange: (course: Course | null) => void;
}

const fetchCourseOptions = async (
  selectedTerm: string,
  query: string
): Promise<Course[]> => {
  const response = await fetch(
    `http://localhost:8081/coursesearch/${selectedTerm}/${query}`
  );
  if (!response.ok) {
    throw new Error("Failed to fetch course options");
  }
  return response.json();
};

export const CourseCombobox: React.FC<CourseComboboxProps> = ({ onChange }) => {
  const [options, setOptions] = useState<Course[]>([]);
  const { selectedTerm, selectedCourses, query, setQuery, setSelectedCourses } =
    useFormContext();

  useEffect(() => {
    if (query.length === 0) {
      setOptions([]);
      return;
    }

    const fetchData = async () => {
      try {
        const data = await fetchCourseOptions(selectedTerm, query);
        console.log("Fetched course options:", data);
        setOptions(data);
      } catch (error) {
        console.error("Error fetching course options:", error);
        setOptions([]);
      }
    };

    fetchData();
  }, [selectedTerm, query]);

  const getCourseDisplay = (course: Course | null) => {
    if (!course) return "";
    return `${course.subject} ${course.catalog} – ${course.courseTitle}`;
  };

  const highlightQuery = (text: string, query: string) => {
    if (!query) return text;

    const escapeRegExp = (string: string) => {
      return string.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
    };

    const normalizedQuery = query.toLowerCase().split(/\s+/).map(escapeRegExp);

    const regex = new RegExp(`(${normalizedQuery.join("|")})`, "gi");
    const parts = text.split(regex);

    return parts.map((part, index) =>
      normalizedQuery.includes(part.toLowerCase()) ? (
        <span key={index} className={ComboboxStyles.MatchingQuery}>
          {part}
        </span>
      ) : (
        <span key={index}>{part}</span>
      )
    );
  };

  const handleOnChange = (course: Course) => {
    setSelectedCourses(course);
    setQuery(getCourseDisplay(course));

    fetch(`http://localhost:8081/course/${selectedTerm}/${course.courseId}`)
      .then((res) => res.json())
      .then(() => onChange(course))
      .catch((err) => console.error("Failed to fetch course details:", err));
  };

  return (
    <div className={ComboboxStyles.Container}>
      <Combobox value={selectedCourses} onChange={handleOnChange}>
        <div className={ComboboxStyles.InputContainer}>
          <ComboboxInput
            aria-label="Selected Course"
            placeholder="Search for courses..."
            displayValue={getCourseDisplay}
            value={query}
            onChange={(event) => setQuery(event.target.value)}
            className={ComboboxStyles.Input}
          />
          <ComboboxButton className={ComboboxStyles.Button}>
            <ChevronDownIcon className={ComboboxStyles.ChevronIcon} />
          </ComboboxButton>
        </div>

        <ComboboxOptions anchor="bottom" className={ComboboxStyles.Options}>
          {options &&
            options.length > 0 &&
            options.map((course) => (
              <ComboboxOption
                key={course.classNumber}
                value={course}
                className={ComboboxStyles.Option}
              >
                <div className={ComboboxStyles.Text}>
                  {highlightQuery(getCourseDisplay(course), query)}
                </div>
              </ComboboxOption>
            ))}
        </ComboboxOptions>
      </Combobox>
    </div>
  );
};
