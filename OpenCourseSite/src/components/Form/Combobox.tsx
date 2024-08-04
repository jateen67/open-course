import { Combobox, ComboboxInput, ComboboxButton, ComboboxOption, ComboboxOptions } from "@headlessui/react"
import { ChevronDownIcon } from '@heroicons/react/20/solid'
import { useState } from "react"
import ComboboxStyles from "./Combobox.module.css"

const courses = [
    { id: 1, name: "COMP 202 – Introduction to Programming idk" },
    { id: 2, name: "COMP 250 – Introduction to Computer Science" },
]

const CourseCombobox = () => {
    const [selectedCourse, setSelectedCourse] = useState(courses[0])
    const [query, setQuery] = useState("")

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
                        autofocus
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