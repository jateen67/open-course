import { Combobox, ComboboxInput, ComboboxButton, ComboboxOption, ComboboxOptions } from "@headlessui/react"
import { ChevronDownIcon } from '@heroicons/react/20/solid'
import { useState, useEffect, useMemo } from "react"
import { Course } from "../../typing"
import ComboboxStyles from "./Combobox.module.css"
import coursesData from "../../data/courses.json"

interface CourseComboboxProps {
    onChange: (value: string) => void;
}

const CourseCombobox: React.FC<CourseComboboxProps> = ({ onChange }) => {
    const [selectedCourse, setSelectedCourse] = useState<Course | null>(null);
    const [query, setQuery] = useState("");

    const getCourseDisplay = (course: Course | null) => {
        if (!course) return "";
        const formattedCourseCode = course.courseCode.replace("-", " ");
        return `${formattedCourseCode} – ${course.courseTitle}`;
    };
    
    const filteredCourses = useMemo(() => {
        const normalizedQuery = query.toLowerCase().split(/\s+/);
        return coursesData.filter((course) => {
            const courseString = getCourseDisplay(course).toLowerCase();
            return normalizedQuery.every((word) => courseString.includes(word));
        });
    }, [query]);

    const handleOnChange = (course: Course) => {
        setSelectedCourse(course);
        onChange(getCourseDisplay(course));
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
                        onChange={(event) => setQuery(event.target.value)}
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
                            <div className={ComboboxStyles.Text}>{`${getCourseDisplay(course)}`}</div>
                        </ComboboxOption>
                    ))}
                </ComboboxOptions>
            </Combobox>
        </div>
    )
}

export default CourseCombobox;