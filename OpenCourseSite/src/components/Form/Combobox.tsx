import { Course } from "../../models"
import { Combobox, ComboboxInput, ComboboxButton, ComboboxOption, ComboboxOptions } from "@headlessui/react"
import { ChevronDownIcon } from '@heroicons/react/20/solid'
import { useState, useEffect } from "react"
import ComboboxStyles from "./Combobox.module.css"

interface CourseComboboxProps {
    termCode: string;
    query: string;
    onQueryChange: (query: string) => void;
    onChange: (course: Course[]) => void;
}

const fetchCourseSuggestions = async (termCode: string, input: string): Promise<Course[]> => {
    const response = await fetch(`http://localhost:8081/coursesearch/${termCode}/${input}`);
    if (!response.ok) {
        throw new Error("Failed to fetch course suggestions");
    }
    return response.json();
};

const CourseCombobox: React.FC<CourseComboboxProps> = ({ termCode, query, onQueryChange, onChange }) => {
    const [suggestions, setSuggestions] = useState<Course[]>([]);
    const [selectedCourse, setSelectedCourse] = useState<Course | null>(null);

    useEffect(() => {
        if (query.length === 0) {
            setSuggestions([]);
            return;
        }

        const fetchData = async () => {
            try {
                const data = await fetchCourseSuggestions(termCode, query);
                setSuggestions(data);
            } catch (error) {
                console.error("Error fetching course suggestions:", error);
                setSuggestions([]);
            }
        };

        fetchData();
    }, [termCode, query]);

    const getCourseDisplay = (course: Course | null) => {
        if (!course) return "";
        return `${course.subject} ${course.catalog} – ${course.courseTitle}`;
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

    const handleOnChange = (course: Course) => {
        setSelectedCourse(course.courseId);
        onQueryChange(getCourseDisplay(course));

        fetch(`http://localhost:8081/coursesearch/${termCode}/${input}`)
            .then((res) => res.json())
            .then((data) => onChange(data))
            .catch((err) => console.error("Failed to fetch course details:", err));
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
                            key={course.classNumber}
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