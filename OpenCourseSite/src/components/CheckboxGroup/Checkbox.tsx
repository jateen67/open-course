import { Checkbox as HeadlessCheckbox, Field, Label } from "@headlessui/react"
import { useState } from "react";
import CheckboxStyles from "./CheckboxGroup.module.css";

interface CheckboxProps {
    section: string;
    classNumber: number;
    mondays: boolean;
    tuesdays: boolean;
    wednesdays: boolean;
    thursdays: boolean;
    fridays: boolean;
    saturdays: boolean;
    sundays: boolean;
    times: string;
    onChange: (checked: boolean) => void;
}

const Checkbox: React.FC<CheckboxProps> = ({
        section,
        classNumber, 
        mondays, 
        tuesdays, 
        wednesdays, 
        thursdays, 
        fridays, 
        saturdays, 
        sundays, 
        times, 
        onChange 
    }) => {
    const [enabled, setEnabled] = useState(false);

    const toggleCheck = () => {
        setEnabled(!enabled);
        onChange(!enabled);
    };

    const formatDays = () => {
        const days = [
            { label: "M ", active: mondays },
            { label: "T ", active: tuesdays },
            { label: "W ", active: wednesdays },
            { label: "T ", active: thursdays },
            { label: "F ", active: fridays },
            { label: "S ", active: saturdays },
            { label: "S", active: sundays },
        ];

        return days.map((day, index) => (
            <span
                key={index}
                className={day.active ? CheckboxStyles.ActiveDay : CheckboxStyles.InactiveDay}
            >
                {day.label}
            </span>
        ));
    };
    
    return (
        <>
        <Field className={enabled ? `${CheckboxStyles.Row} ${CheckboxStyles.Checked}` : CheckboxStyles.Row} onClick={toggleCheck}>
            <div className={CheckboxStyles.BoxContainer}>
                <HeadlessCheckbox
                    checked={enabled}
                    onChange={toggleCheck}
                    className={`${CheckboxStyles.CheckboxContainer} ${enabled ? CheckboxStyles.checked : ""}`}
                >
                    <svg
                        className={`${CheckboxStyles.CheckIcon} ${enabled ? CheckboxStyles.checked : ""}`}
                        viewBox="0 0 14 14"
                        fill="none"
                    >
                        <path
                            d="M3 8L6 11L11 3.5"
                            strokeWidth={2}
                            strokeLinecap="round"
                            strokeLinejoin="round"
                        />
                    </svg>
                </HeadlessCheckbox>
            </div>
            <Label className={CheckboxStyles.LabelRow}>
                    <div className={CheckboxStyles.SectionContainer}>
                        <p>{section}</p>
                    </div>
                    <div className={CheckboxStyles.CRNContainer}>
                        <p>{classNumber}</p>
                    </div>
                    <div className={CheckboxStyles.DayContainer}>
                        <p>{formatDays()}</p>
                    </div>
                    <div className={CheckboxStyles.TimeContainer}>
                        <p>{times}</p>
                    </div>
            </Label>
        </Field>
        </>
    );
}

export default Checkbox;