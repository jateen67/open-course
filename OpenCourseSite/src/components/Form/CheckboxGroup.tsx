import { useState, useEffect } from "react";
import Checkbox from "./Checkbox";
import CheckboxGroupStyles from "./CheckboxGroup.module.css"
import { Course } from "../../models"
import { useFormContext } from "contexts";

interface CheckboxGroupProps {
    course: Course;
    onChange: (courseIds: number[]) => void;
}

const headers = [
    { label: "Section", containerClass: CheckboxGroupStyles.SectionContainer },
    { label: "CRN", containerClass: CheckboxGroupStyles.CRNContainer },
    { label: "Days", containerClass: CheckboxGroupStyles.DayContainer },
    { label: "Time", containerClass: CheckboxGroupStyles.TimeContainer }
];

const fetchCourseInfo = async (termCode: string, courseId: number): Promise<Course[]> => {
    console.log("selectedTerm:", termCode);
    console.log("courseId:", courseId);

    const response = await fetch(`http://localhost:8081/course/${termCode}/${courseId}`);
    if (!response.ok) {
        throw new Error("Failed to fetch course sections");
    }
    return response.json();
};

const CheckboxGroup: React.FC<CheckboxGroupProps> = ({ onChange }) => {
    const { selectedTerm, selectedCourses, selectedCheckboxes, setSelectedCheckboxes } = useFormContext();
    const [sections, setSections] = useState<Course[]>([]);

    useEffect(() => {
        if (!selectedCourses) {
            setSections([]);
            return;
        }

        const fetchData = async () => {
            try {
                const data = await fetchCourseInfo(selectedTerm, selectedCourses.courseId);
                setSections(data);
            } catch (error) {
                console.error("Error fetching course sections:", error);
                setSections([]);
            }
        };

        fetchData();
    }, [selectedTerm, selectedCourses]);

    useEffect(() => {
        console.log("Selected Sections Updated:", selectedCheckboxes);
        onChange(selectedCheckboxes);
    }, [selectedCheckboxes, onChange]);

    function formatTime(time: string): string {
        const [hour, minute] = time.split('.');
    
        return `${hour}:${minute}`;
    }

    const handleOnChange = (checked: boolean, classNumber: number) => {
        setSelectedCheckboxes((prevSelectedCheckboxes: number[]) => {
            const updatedSections = checked
                ? [...prevSelectedCheckboxes, classNumber]
                : prevSelectedCheckboxes.filter((id : number) => id !== classNumber);
    
            return updatedSections;
        });
    };

    return (
        <div className={CheckboxGroupStyles.Container}>
            <div className={CheckboxGroupStyles.HeaderRow}>
                <div className={CheckboxGroupStyles.BoxContainer}>
                </div>
                {headers.map(header => (
                    <div className={header.containerClass} key={header.label}>
                        <p className={CheckboxGroupStyles.Heading}>{header.label}</p>
                    </div>
                ))}
            </div>
            {sections.map((section) => (
                <Checkbox
                    key={section.classNumber}
                    section={section.componentCode + " " + section.section}
                    classNumber={section.classNumber}
                    mondays={section.mondays}
                    tuesdays={section.tuesdays}
                    wednesdays={section.wednesdays}
                    thursdays={section.thursdays}
                    fridays={section.fridays}
                    saturdays={section.saturdays}
                    sundays={section.sundays}
                    times={formatTime(section.classStartTime) + " – " + formatTime(section.classEndTime)}
                    onChange={(checked: boolean) => handleOnChange(checked, section.classNumber)}
                />
            ))}
        </div>
    );
}

export default CheckboxGroup;