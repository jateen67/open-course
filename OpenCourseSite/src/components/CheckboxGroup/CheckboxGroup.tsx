import { useState, useEffect } from "react";
import Checkbox from "./Checkbox";
import CheckboxGroupStyles from "./CheckboxGroup.module.css";
import { Course } from "../../models";
import { useFormContext } from "contexts";
import { CourseService } from "services";

interface CheckboxGroupProps {
  onChange: (courseIds: number[]) => void;
}

const headers = [
  { label: "Section", containerClass: CheckboxGroupStyles.SectionContainer },
  { label: "CRN", containerClass: CheckboxGroupStyles.CRNContainer },
  { label: "Days", containerClass: CheckboxGroupStyles.DayContainer },
  { label: "Time", containerClass: CheckboxGroupStyles.TimeContainer },
];

function formatTime(time: string): string {
  const [hour, minute] = time.split(".");

  return `${hour}:${minute}`;
}

export const CheckboxGroup: React.FC<CheckboxGroupProps> = ({ onChange }) => {
  const {
    selectedTerm,
    selectedCourses,
    selectedCheckboxes,
    setSelectedCheckboxes,
  } = useFormContext();
  const [sections, setSections] = useState<Course[]>([]);

  useEffect(() => {
    if (!selectedCourses) {
      setSections([]);
      return;
    }

    const fetchData = () => {
      CourseService.GetByTermCodeAndCourseId(selectedTerm, selectedCourses.courseId).subscribe({
        next: (courses) => setSections(courses),
        error: () => setSections([])
      })
    };

    fetchData();
  }, [selectedTerm, selectedCourses]);

  useEffect(() => {
    onChange(selectedCheckboxes);
  }, [selectedCheckboxes, onChange]);

  const handleOnChange = (checked: boolean, classNumber: number) => {
    setSelectedCheckboxes((prevSelectedCheckboxes: number[]) => {
      const updatedSections = checked
        ? [...prevSelectedCheckboxes, classNumber]
        : prevSelectedCheckboxes.filter((id: number) => id !== classNumber);

      return updatedSections;
    });
  };

  return (
    <div className={CheckboxGroupStyles.Container}>
      <div className={CheckboxGroupStyles.HeaderRow}>
        <div className={CheckboxGroupStyles.BoxContainer}></div>
        {headers.map(({ label, containerClass }) => (
          <div className={containerClass} key={label}>
            <p className={CheckboxGroupStyles.Heading}>{label}</p>
          </div>
        ))}
      </div>
      {sections.map(
        ({
          classNumber,
          componentCode,
          section,
          classStartTime,
          classEndTime,
          ...days
        }) => (
          <Checkbox
            key={classNumber}
            section={`${componentCode} ${section}`}
            classNumber={classNumber}
            {...days}
            times={`${formatTime(classStartTime)} – ${formatTime(
              classEndTime
            )}`}
            onChange={(checked: boolean) =>
              handleOnChange(checked, classNumber)
            }
          />
        )
      )}
    </div>
  );
};
