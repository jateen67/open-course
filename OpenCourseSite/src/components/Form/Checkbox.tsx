import { Checkbox as HeadlessCheckbox } from '@headlessui/react'
import { useState } from 'react'
import CheckboxStyles from "./CheckboxGroup.module.css"

const Checkbox = () => {
    const [enabled, setEnabled] = useState(false);
    
    return (
        <>
        <div className={CheckboxStyles.Row}>
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
            <div className={CheckboxStyles.SectionContainer}>
                <p>Lec 001</p>
            </div>
            <div className={CheckboxStyles.CRNContainer}>
                <p>8102</p>
            </div>
            <div className={CheckboxStyles.DayContainer}>
                <p>tuesday</p>
                <p>thursday</p>
            </div>
            <div className={CheckboxStyles.TimeContainer}>
                <p>11:00 – 12:45</p>
            </div>
        </div>
        <hr/>
        </>
    );
}

export default Checkbox;