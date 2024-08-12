import { Combobox, ComboboxInput, ComboboxButton, ComboboxOption, ComboboxOptions } from "@headlessui/react"
import { ChevronDownIcon } from '@heroicons/react/20/solid'
import { useState, useEffect } from "react"
import ComboboxStyles from "./Combobox.module.css"

interface Course {
    courseCode: string;
    courseTitle: string;
}

const CourseCombobox = () => {
    const [selectedCourse, setSelectedCourse] = useState("")
    const [query, setQuery] = useState("")
    const [courses, setCourses] = useState<Course[]>([])

    useEffect(() => {
        fetch("../../data/courses.json")
            .then(response => response.json())
            .then(data => setCourses(data))
            .catch(error => console.error('Error fetching courses:', error))
    }, [])

    const filteredCourses =
        query === ""
        ? courses
        : courses.filter((course) => {
            return course.name.toLowerCase().includes(query.toLowerCase())
            })

    return (
        <div className={ComboboxStyles.Container}>
            <Combobox
                value={selectedCourse}
                onChange={(value) => setSelectedCourse(value as { id: number; name: string })}
            >
                <div className={ComboboxStyles.InputContainer}>
                    <ComboboxInput
                        aria-label="Selected Course"
                        placeholder="Search for courses..."
                        displayValue={(selectedCourse: { id: number; name: string }) => selectedCourse?.name}
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
                            <div className={ComboboxStyles.Text}>{course.name}</div>
                        </ComboboxOption>
                    ))}
                </ComboboxOptions>
            </Combobox>
        </div>
    )
}

export default CourseCombobox;