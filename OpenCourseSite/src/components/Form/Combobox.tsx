import { Combobox, ComboboxInput, ComboboxButton, ComboboxOption, ComboboxOptions } from "@headlessui/react"
import { ChevronDownIcon } from '@heroicons/react/20/solid'
import { useState, useEffect, useMemo } from "react"
import { Course } from "../../typing"
import ComboboxStyles from "./Combobox.module.css"
import coursesData from "../../data/courses.json"

interface CourseComboboxProps {
    term: string;
    query: string;
    onQueryChange: (query: string) => void;
    onChange: (course: Course[]) => void;
}

const CourseCombobox: React.FC<CourseComboboxProps> = ({ term, query, onQueryChange, onChange }) => {
    const [selectedCourse, setSelectedCourse] = useState<Course | null>(null);

    useEffect(() => {
        setSelectedCourse(null);
    }, [term]);

    const getCourseDisplay = (course: Course | null) => {
        if (!course) return "";
        const formattedCourseCode = course.courseCode.replace("-", " ");
        return `${formattedCourseCode} – ${course.courseTitle}`;
    };

    console.log("Current term in Combobox:", term);
    console.log("Current query in Combobox:", query);

    const highlightMatch = (text: string, query: string) => {
        if (!query) return text;
    
        const normalizedQuery = query.toLowerCase().split(/\s+/);
        const regex = new RegExp(`(${normalizedQuery.join('|')})`, 'gi');
        const parts = text.split(regex);
    
        return parts.map((part, index) =>
            normalizedQuery.includes(part.toLowerCase()) ? (
                <span key={index} className={ComboboxStyles.MatchingQuery}>{part}</span>
            ) : (
                <span key={index}>{part}</span>
            )
        );
    };

    const filterCourses = (courses: Course[], term: string, query: string) => {
        const normalizedTerm = term.toLowerCase();
        const normalizedQuery = query.toLowerCase().split(/\s+/);

        return courses.filter((course) => {
            const normalizedSemester = course.semester.toLowerCase();
            const matchesTerm = normalizedTerm === "fall|winter"
                ? normalizedSemester.includes("fall") && normalizedSemester.includes("winter")
                : normalizedSemester.includes(normalizedTerm);

            const courseString = getCourseDisplay(course).toLowerCase();
            const matchesQuery = normalizedQuery.every((word) => courseString.includes(word));

            return matchesTerm && matchesQuery;
        });
    };

    const getUniqueCourses = (courses: Course[]) => {
        const uniqueCourses = new Map<string, Course>();
        courses.forEach((course) => {
            const key = `${course.courseCode}-${course.courseTitle}`;
            if (!uniqueCourses.has(key)) {
                uniqueCourses.set(key, course);
            }
        });
        return Array.from(uniqueCourses.values());
    };

    const filteredCourses = useMemo(() => {
        const coursesMatchingTermAndQuery = filterCourses(coursesData, term, query);
        return getUniqueCourses(coursesMatchingTermAndQuery);
    }, [term, query]);

    const handleOnChange = (course: Course) => {
        const courseDisplayValue = getCourseDisplay(course);
        setSelectedCourse(course);
        onQueryChange(courseDisplayValue);

        const relatedCourses = filterCourses(
            coursesData.filter(
                (c) => c.courseCode === course.courseCode && c.courseTitle === course.courseTitle
            ),
            term,
            ""
        );

        onChange(relatedCourses);
    };

    return (
        <div className={ComboboxStyles.Container}>
            <Combobox
                value={selectedCourse}
                onChange={handleOnChange}
            >
                <div className={ComboboxStyles.InputContainer}>
                    <ComboboxInput
                        aria-label="Selected Course"
                        placeholder="Search for courses..."
                        displayValue={getCourseDisplay}
                        value={query}
                        onChange={(event) => onQueryChange(event.target.value)}
                        className={ComboboxStyles.Input}
                    />
                    <ComboboxButton className={ComboboxStyles.Button}>
                        <ChevronDownIcon className={ComboboxStyles.ChevronIcon} />
                    </ComboboxButton>
                </div>
                
                <ComboboxOptions
                    anchor="bottom"
                    className={ComboboxStyles.Options}
                >
                    {filteredCourses.map((course) => (
                        <ComboboxOption
                            key={course.id}
                            value={course}
                            className={ComboboxStyles.Option}
                        >
                            <div className={ComboboxStyles.Text}>
                                {highlightMatch(getCourseDisplay(course), query)}
                            </div>
                        </ComboboxOption>
                    ))}
                </ComboboxOptions>
            </Combobox>
        </div>
    )
}

export default CourseCombobox;