import { Course } from "../../models"
import { Combobox, ComboboxInput, ComboboxButton, ComboboxOption, ComboboxOptions } from "@headlessui/react"
import { ChevronDownIcon } from "@heroicons/react/20/solid"
import { useState, useEffect } from "react"
import ComboboxStyles from "./Combobox.module.css"

interface CourseComboboxProps {
    termCode: string;
    query: string;
    onQueryChange: (query: string) => void;
    onChange: (course: Course | null) => void;
}

const fetchCourseOptions = async (termCode: string, query: string): Promise<Course[]> => {
    const response = await fetch(`http://localhost:8081/coursesearch/${termCode}/${query}`);
    if (!response.ok) {
        throw new Error("Failed to fetch course options");
    }
    return response.json();
};

const CourseCombobox: React.FC<CourseComboboxProps> = ({ termCode, query, onQueryChange, onChange }) => {
    const [options, setOptions] = useState<Course[]>([]);
    const [selectedCourse, setSelectedCourse] = useState<Course | null>(null);

    useEffect(() => {
        if (query.length === 0) {
            setOptions([]);
            return;
        }

        const fetchData = async () => {
            try {
                const data = await fetchCourseOptions(termCode, query);
                setOptions(data);
            } catch (error) {
                console.error("Error fetching course options:", error);
                setOptions([]);
            }
        };

        fetchData();
    }, [termCode, query]);

    const getCourseDisplay = (course: Course | null) => {
        if (!course) return "";
        return `${course.subject} ${course.catalog} – ${course.courseTitle}`;
    };

    const highlightQuery = (text: string, query: string) => {
        if (!query) return text;
    
        const escapeRegExp = (string: string) => {
            return string.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
        };
    
        const normalizedQuery = query
            .toLowerCase()
            .split(/\s+/)
            .map(escapeRegExp);
    
        const regex = new RegExp(`(${normalizedQuery.join("|")})`, "gi");
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
        setSelectedCourse(course);
        onQueryChange(getCourseDisplay(course));

        fetch(`http://localhost:8081/coursesearch/${termCode}/${query}`)
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
                    {options && options.length > 0 && (
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
                        ))
                    )}
                </ComboboxOptions>
            </Combobox>
        </div>
    )
}

export default CourseCombobox;