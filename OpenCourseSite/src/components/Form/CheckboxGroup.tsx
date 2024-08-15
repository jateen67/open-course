import { useState, useEffect } from "react";
import Checkbox from "./Checkbox";
import CheckboxGroupStyles from "./CheckboxGroup.module.css"
import { Course } from "../../models"

interface CheckboxGroupProps {
    termCode: string;
    course: Course;
}

const headers = [
    { label: "Section", containerClass: CheckboxGroupStyles.SectionContainer },
    { label: "CRN", containerClass: CheckboxGroupStyles.CRNContainer },
    { label: "Days", containerClass: CheckboxGroupStyles.DayContainer },
    { label: "Time", containerClass: CheckboxGroupStyles.TimeContainer }
];

const fetchCourseInfo = async (termCode: string, courseId: number): Promise<Course[]> => {
    console.log("termCode:", termCode);
    console.log("courseId:", courseId);

    const response = await fetch(`http://localhost:8081/course/${termCode}/${courseId}`);
    if (!response.ok) {
        throw new Error("Failed to fetch course sections");
    }
    return response.json();
};

const CheckboxGroup: React.FC<CheckboxGroupProps> = ({ termCode, course }) => {
    const [sections, setSections] = useState<Course[]>([]);

    useEffect(() => {
        console.log(course);

        if (!course) {
            setSections([]);
            return;
        }

        const fetchData = async () => {
            try {
                const data = await fetchCourseInfo(termCode, course.courseId);
                setSections(data);
            } catch (error) {
                console.error("Error fetching course sections:", error);
                setSections([]);
            }
        };

        fetchData();
    }, [termCode, course]);

    function formatTime(timeString: string): string {
        const [hour, minute] = timeString.split('.');
    
        return `${hour}:${minute}`;
    }

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
                    section={section.componentCode + " " + section.section}
                    crn={section.classNumber}
                    mondays={section.mondays}
                    tuesdays={section.tuesdays}
                    wednesdays={section.wednesdays}
                    thursdays={section.thursdays}
                    fridays={section.fridays}
                    saturdays={section.saturdays}
                    sundays={section.sundays}
                    times={formatTime(section.classStartTime) + " – " + formatTime(section.classEndTime)}
                />
            ))}
        </div>
    );
}

export default CheckboxGroup;