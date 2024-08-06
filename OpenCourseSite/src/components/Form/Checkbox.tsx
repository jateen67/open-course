import { Checkbox as HeadlessCheckbox, Field, Label } from '@headlessui/react'
import { useState } from 'react'
import CheckboxStyles from "./CheckboxGroup.module.css"

interface CheckboxProps {
    section: string;
    crn: number;
    days: string[];
    times: string[];
}

const Checkbox: React.FC<CheckboxProps> = ({ section, crn, days, times }) => {
    const [enabled, setEnabled] = useState(false);
    const toggleCheck = () => setEnabled(!enabled);

    
    return (
        <>
        <Field className={enabled ? `${CheckboxStyles.Row} ${CheckboxStyles.Checked}` : CheckboxStyles.Row} onClick={toggleCheck}>
            <div className={CheckboxStyles.BoxContainer}>
                <HeadlessCheckbox
                    checked={enabled}
                    onChange={setEnabled}
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
                        <p>{crn}</p>
                    </div>
                    <div className={CheckboxStyles.DayContainer}>
                        {days.map((day) => (
                            <p>{day}</p>
                        ))}
                    </div>
                    <div className={CheckboxStyles.TimeContainer}>
                        {times.map((time) => (
                            <p>{time}</p>
                        ))}
                    </div>
            </Label>
        </Field>
        </>
    );
}

export default Checkbox;