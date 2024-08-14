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

const fetchCourseSuggestions = async (termCode: string, input: string): Promise<Course[]> => {
    const response = await fetch(`/coursesearch/${termCode}/${input}`);
    if (!response.ok) {
        throw new Error("Failed to fetch course suggestions");
    }
    return response.json();
};

const CourseCombobox: React.FC<CourseComboboxProps> = ({ term, query, onQueryChange, onChange }) => {
    const [suggestions, setSuggestions] = useState<Course[]>([]);
    const [selectedCourse, setSelectedCourse] = useState<Course | null>(null);

    useEffect(() => {
        if (query.length === 0) {
            setSuggestions([]);
            return;
        }

        const fetchData = async () => {
            try {
                const data = await fetchCourseSuggestions(term, query);
                setSuggestions(data);
            } catch (error) {
                console.error("Error fetching course suggestions:", error);
                setSuggestions([]);
            }
        };

        fetchData();
    }, [term, query]);

    const handleOnChange = (course: Course) => {
        const courseDisplayValue = getCourseDisplay(course);
        setSelectedCourse(course);
        onQueryChange(courseDisplayValue);

        fetch(`/course/${term}/${course.id}`)
            .then((res) => res.json())
            .then((data) => onChange(data))
            .catch((err) => console.error("Failed to fetch course details:", err));
    };

    const getCourseDisplay = (course: Course | null) => {
        if (!course) return "";
        const formattedCourseCode = course.courseCode.replace("-", " ");
        return `${formattedCourseCode} – ${course.courseTitle}`;
    };

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
                    {suggestions.map((course) => (
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